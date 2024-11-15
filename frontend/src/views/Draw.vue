<template>
  <div class="content">
    <div class="app-container">
      <div class="panel left-panel">
        <h2>手绘图像</h2>
        <canvas
            ref="canvas"
            width="600"
            height="400"
            @mousedown="startDrawing"
            @mousemove="draw"
            @mouseup="stopDrawing"
            @mouseleave="stopDrawing"
            class="drawing-canvas"
        />
        <div class="canvas-controls">

<!--        <div class="controls">-->
<!--          <label>-->
            <i class="fas fa-palette"></i>
            <input type="color" v-model="color" />
          <span class="separator"></span>
            <i class="fas fa-paint-brush"></i>
            <input type="range" min="1" max="20" v-model="brushSize" />
          <span class="separator"></span>
<!--          </label>-->
<!--        </div>-->
<!--        <div class="canvas-controls">-->
          <button @click="setTool('brush')" class="icon-button"><i class="fas fa-paint-brush"></i></button>
<!--          <span class="separator"></span>-->
          <button @click="setTool('eraser')" class="icon-button"><i class="fas fa-eraser"></i></button>
<!--          <span class="separator"></span>-->
          <button @click="clearCanvas" class="icon-button"><i class="fas fa-trash"></i></button>
          <span class="separator"></span>
          <button @click="handleZoomIn" class="icon-button"><i class="fas fa-search-plus"></i></button>
          <button @click="handleZoomOut" class="icon-button"><i class="fas fa-search-minus"></i></button>
<!--          <span class="separator"></span>-->
          <button @click="handleRotateLeft" class="icon-button"><i class="fas fa-undo"></i></button>
          <button @click="handleRotateRight" class="icon-button"><i class="fas fa-redo"></i></button>
<!--          <span class="separator"></span>-->
          <button @click="resetCanvas" class="icon-button"><i class="fas fa-sync"></i></button>
        </div>

        <!--        <canvas ref="canvas" class="drawing-canvas" width="600" height="400" @mousedown="startDrawing" @mousemove="draw" @mouseup="stopDrawing" @mouseleave="stopDrawing"/>-->

        <ImageUpload @image-upload="handleImageUpload" />
      </div>

      <div class="separator"></div>

      <div class="panel right-panel">
        <h2>AI 处理图像</h2>
        <div class="processed-image-area" :class="{ 'placeholder': !processedImage }">
          <img v-if="processedImage" :src="processedImage" alt="Processed" />
          <div v-else class="placeholder-text">图像处理结果将在此显示</div>
        </div>
        <input type="text" placeholder="请输入描述..." v-model="prompt" class="prompt-input" />
        <div class="button-group">
          <button class="styled-button" @click="handleProcessImage">AI 处理图像</button>
          <button class="styled-button" @click="saveImage">保存图像到分组</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import ImageUpload from '@/components/ImageUpload.vue';

export default {
  components: { ImageUpload },
  data() {
    return {
      uploadedImage: null,
      processedImage: null,
      prompt: '',
      color: '#000000',
      brushSize: 5,
      tool: 'brush',
      isDrawing: false,
      scale: 1,
      rotation: 0,
      // canvasBackgroundColor: '#f5f5f5', // 画板的背景色
    };
  },
  mounted() {
    this.fillCanvasBackground();
    console.log("this.$route.params.category_id:",this.$route.params.category_id)
  },
  methods: {
    fillCanvasBackground() {
      const canvas = this.$refs.canvas;
      const context = canvas.getContext('2d');

      // 填充白色背景
      context.fillStyle = '#ffffff';
      context.fillRect(0, 0, canvas.width, canvas.height);
    },
    handleImageUpload(file) {
      this.uploadedImage = URL.createObjectURL(file);
      this.drawImageOnCanvas();
    },
    drawImageOnCanvas() {
      const canvas = this.$refs.canvas;
      const context = canvas.getContext('2d');

      // 填充白色背景
      context.fillStyle = '#ffffff';
      context.fillRect(0, 0, canvas.width, canvas.height);

      const image = new Image();
      image.src = this.uploadedImage;
      image.onload = () => {
        const aspectRatio = image.width / image.height;
        let drawWidth = canvas.width;
        let drawHeight = drawWidth / aspectRatio;
        if (drawHeight > canvas.height) {
          drawHeight = canvas.height;
          drawWidth = drawHeight * aspectRatio;
        }
        context.save();
        context.translate(canvas.width / 2, canvas.height / 2);
        context.rotate((this.rotation * Math.PI) / 180);
        context.scale(this.scale, this.scale);
        context.drawImage(image, -drawWidth / 2, -drawHeight / 2, drawWidth, drawHeight);
        context.restore();
      };
    },
    handleZoomIn() {
      this.scale += 0.1;
      this.drawImageOnCanvas();
    },
    handleZoomOut() {
      this.scale -= 0.1;
      this.drawImageOnCanvas();
    },
    handleRotateLeft() {
      this.rotation -= 15;
      this.drawImageOnCanvas();
    },
    handleRotateRight() {
      this.rotation += 15;
      this.drawImageOnCanvas();
    },
    resetCanvas() {
      this.scale = 1;
      this.rotation = 0;
      this.drawImageOnCanvas();
    },
    clearCanvas() {
      const canvas = this.$refs.canvas;
      const context = canvas.getContext('2d');

      // 填充白色背景
      context.fillStyle = '#ffffff';
      context.fillRect(0, 0, canvas.width, canvas.height);
    },
    setTool(selectedTool) {
      this.tool = selectedTool;
    },
    startDrawing(event) {
      this.isDrawing = true;
      const ctx = this.$refs.canvas.getContext('2d');
      ctx.beginPath();
      ctx.moveTo(event.offsetX, event.offsetY);
    },
    draw(event) {
      if (!this.isDrawing) return;
      const ctx = this.$refs.canvas.getContext('2d');
      ctx.lineTo(event.offsetX, event.offsetY);
      ctx.lineWidth = this.brushSize;
      ctx.lineCap = 'round';
      ctx.strokeStyle = this.tool === 'brush' ? this.color : "#ffffff";
      ctx.stroke();
    },
    stopDrawing() {
      this.isDrawing = false;
    },
    async handleProcessImage() {
      try {
        const canvasDataURL = this.$refs.canvas.toDataURL('image/png');
        const response = await this.$axios({
          method: "post",
          url: "/processImage",
          data: {
            prompt: this.prompt,
            ori_image: canvasDataURL,
          }
        });

        console.log("response:",response)
        if (response.code === 1000) {
          this.processedImage = "data:image/png;base64," + response.data;
          // // 将返回的处理后图像转换为 base64
          // const imageBlob = await fetch(response.data).then(res => res.blob());
          // const reader = new FileReader();
          // reader.onloadend = () => {
          //   this.processedImage = reader.result; // base64 数据
          // };
          // reader.readAsDataURL(imageBlob);
          } else {
          console.log("处理图片失败1:", response.data.msg);
        }
      } catch (error) {
        console.error("处理图片失败2:", error);
      }
    },
    async saveImage() {
      try {
        if (!this.processedImage) {
          console.log("没有处理后的图像可保存");
          return;
        }

        // console.log("this.processedImage:",this.processedImage)
        // const canvasDataURL = this.$refs.canvas.toDataURL('image/png');
        const response = await this.$axios({
          method: "post",
          url: "/uploadWork",
          data: {
            work_image: this.processedImage,
            category_id: this.$route.params.category_id,
            prompt: this.prompt
          }
        });

        console.log("saveImage response:", response)
        if (response.code === 1000) {
          // await this.$router.push({name: "CategoryDetail", params: {category_id: this.$route.params.category_id}});
          console.log("上传成功：", response.message
          );

        } else {
          console.log("上传作品错误1:", response.data.msg);
        }
      } catch (error) {
        console.error("上传作品错误2:", error);
      }
    },
  },
};
</script>

<style scoped>
.content {
  display: flex;
  justify-content: center;
  padding: 20px 24px;
  margin-top: 48px;
}

.app-container {
  display: flex;
  justify-content: space-between;
  max-width: 100%;
}

.controls label {
  display: flex;
  align-items: center;
  font-size: 14px;
  margin-bottom: 8px;
}

.canvas-controls {
  display: flex;
  align-items: center;
  flex-wrap: wrap;
}

.icon-button {
  background: #007bff;
  color: #fff;
  border: none;
  padding: 8px;
  margin-right: 8px; /* 间距 */
  font-size: 16px;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s;
}

.icon-button:hover {
  background-color: #0056b3;
}

.separator {
  width: 1px;
  height: 24px;
  background-color: #ddd;
  margin: 0 8px;
}

.drawing-canvas {
  border: 1px solid #ddd;
  width: 100%;
  height: 100%;
}


.panel {
  flex: 1;
  margin: 0 10px;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.drawing-canvas,
.processed-image-area {
  width: 600px;
  height: 400px;
  border: 1px solid #ddd;
  margin-top: 20px;
}

.processed-image-area.placeholder {
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: #f5f5f5;
  color: #888;
}

.placeholder-text {
  font-size: 18px;
  color: #888;
}

.prompt-input {
  width: 100%;
  margin-top: 10px;
  padding: 8px;
}

.button-group {
  display: flex;
  gap: 8px;
  margin-top: 10px;
}

.styled-button {
  padding: 10px 20px;
  font-size: 16px;
  color: #fff;
  background-color: #007bff;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s ease;

  &:hover {
    background-color: #0056b3;
  }

  &:active {
    background-color: #004085;
  }
}

.styled-button.secondary {
  background-color: #6c757d;

  &:hover {
    background-color: #5a6268;
  }

  &:active {
    background-color: #495057;
  }
}
</style>
