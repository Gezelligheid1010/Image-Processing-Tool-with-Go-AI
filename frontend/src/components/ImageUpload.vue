<template>
  <div class="image-upload">
    <input
        type="file"
        accept="image/*"
        @change="onImageChange"
        id="file-upload"
        class="file-input"
    />
    <label for="file-upload" class="upload-label">上传本地图片</label>
  </div>
</template>

<script>
export default {
  name: 'ImageUpload',
  methods: {
    onImageChange(event) {
      const file = event.target.files[0];
      if (file) {
        this.$emit('image-upload', file); // 将文件对象传递给父组件
      }
    },
  },
};
</script>

<style scoped>
.image-upload {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 1rem;
  border: 2px dashed #ccc;
  border-radius: 8px;
  cursor: pointer;
}
.file-input {
  display: none;
}
.upload-label {
  color: #007bff;
  cursor: pointer;
}
</style>
