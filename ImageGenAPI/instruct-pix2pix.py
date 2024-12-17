from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
import torch
from diffusers import StableDiffusionInstructPix2PixPipeline, EulerAncestralDiscreteScheduler
import requests
from io import BytesIO
from PIL import Image, ImageOps
import base64
from PIL import ImageFile
import asyncio
from collections import deque

ImageFile.LOAD_TRUNCATED_IMAGES = True

app = FastAPI()

# 模型初始化
model_id = "timbrooks/instruct-pix2pix"
if not torch.cuda.is_available():
    raise RuntimeError("CUDA is not available. Please ensure your GPU is enabled.")

pipe = StableDiffusionInstructPix2PixPipeline.from_pretrained(model_id, torch_dtype=torch.float16, safety_checker=None)
pipe.to("cuda")
pipe.scheduler = EulerAncestralDiscreteScheduler.from_config(pipe.scheduler.config)

# 全局任务队列和并发控制
task_queue = deque()
batch_semaphore = asyncio.Semaphore(1)  # 限制同时处理的批次数
BATCH_SIZE = 4  # 每次批量处理的任务数量


class ImageRequest(BaseModel):
    url: str
    prompt: str


class Task:
    """表示一个任务的类"""
    def __init__(self, url, prompt, future):
        self.url = url
        self.prompt = prompt
        self.future = future


def download_image(url):
    """从 URL 下载图像并进行预处理"""
    response = requests.get(url, timeout=10)
    if response.status_code != 200:
        raise HTTPException(status_code=400, detail="Could not download image")
    try:
        image = Image.open(BytesIO(response.content))
        image = ImageOps.exif_transpose(image).convert("RGB")
        # 限制分辨率
        image = image.resize((512, 512), Image.Resampling.LANCZOS)
        return image
    except Exception as e:
        raise HTTPException(status_code=400, detail=f"Image processing error: {str(e)}")


async def process_batch():
    """批量处理队列中的任务"""
    while True:
        await asyncio.sleep(0.1)  # 控制任务轮询速度
        if not task_queue:
            continue

        async with batch_semaphore:  # 确保单次只执行一个批次
            batch = []
            while task_queue and len(batch) < BATCH_SIZE:
                batch.append(task_queue.popleft())

            if batch:
                try:
                    images = [download_image(task.url) for task in batch]
                    prompts = [task.prompt for task in batch]
                    results = pipe(prompt=prompts, image=images, num_inference_steps=10, image_guidance_scale=1)

                    for task, result in zip(batch, results.images):
                        buffered = BytesIO()
                        result.save(buffered, format="PNG")
                        task.future.set_result({"image": base64.b64encode(buffered.getvalue()).decode("utf-8")})
                except Exception as e:
                    for task in batch:
                        task.future.set_exception(HTTPException(status_code=500, detail=f"Model inference error: {str(e)}"))


@app.post("/generate")
async def generate_image(request: ImageRequest):
    """生成图像接口"""
    loop = asyncio.get_running_loop()
    future = loop.create_future()
    task_queue.append(Task(request.url, request.prompt, future))
    return await future


@app.get("/health")
async def health_check():
    """健康检查接口"""
    try:
        torch.cuda.memory_allocated()
        return {"status": "healthy", "gpu": "available"}
    except Exception as e:
        return {"status": "unhealthy", "error": str(e)}


# 启动后台任务
@app.on_event("startup")
async def start_background_tasks():
    asyncio.create_task(process_batch())
