from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
import torch
from diffusers import StableDiffusionInstructPix2PixPipeline, EulerAncestralDiscreteScheduler
import requests
from io import BytesIO
from PIL import Image, ImageOps
import base64
from PIL import ImageFile

ImageFile.LOAD_TRUNCATED_IMAGES = True

app = FastAPI()

model_id = "timbrooks/instruct-pix2pix"
pipe = StableDiffusionInstructPix2PixPipeline.from_pretrained(model_id, torch_dtype=torch.float16, safety_checker=None)
pipe.to("cuda")
pipe.scheduler = EulerAncestralDiscreteScheduler.from_config(pipe.scheduler.config)


class ImageRequest(BaseModel):
    url: str
    prompt: str


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


@app.post("/generate")
def generate_image(request: ImageRequest):
    """生成图像接口"""
    image = download_image(request.url)
    try:
        # 推理过程
        images = pipe(prompt=request.prompt, image=image, num_inference_steps=10, image_guidance_scale=1).images
        buffered = BytesIO()
        images[0].save(buffered, format="PNG")
        img_str = base64.b64encode(buffered.getvalue()).decode("utf-8")
        return {"image": img_str}
    except Exception as e:
        raise HTTPException(status_code=500, detail=f"Model inference error: {str(e)}")
