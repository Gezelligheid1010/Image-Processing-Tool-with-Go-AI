<template>
  <div class="content">
    <div class="form-container">
      <h2>新建分组</h2>
      <form @submit.prevent="saveCategory" class="category-form">
        <div class="form-group">
          <label for="name">分组名称：</label>
          <input type="text" v-model="name" id="name" required placeholder="请输入分组名称" />
        </div>

        <div class="form-group">
          <label for="description">简介：</label>
          <textarea v-model="description" id="description" required placeholder="请输入简介"></textarea>
        </div>

        <div class="form-group">
          <label for="cover">封面图：</label>
          <div class="cover-preview" :style="coverStyle">
            <div v-if="!coverBase64" class="placeholder-text">点击上传</div> <!-- 提示文字层 -->
            <input type="file" @change="onFileChange" id="cover" required class="file-input"/>
          </div>
        </div>


        <div class="buttons">
          <button type="submit" class="btn primary">保存</button>
          <button type="button" @click="cancel" class="btn secondary">取消</button>
        </div>
      </form>
    </div>
  </div>
</template>
<script>

export default {
  data() {
    return {
      name: "",
      description: "",
      cover: null,
      coverBase64: "",  // 用于存储封面图的Base64编码
    };
  },
  computed: {
    coverStyle() {
      return this.coverBase64
          ? { backgroundImage: `url(${this.coverBase64})` }
          : { backgroundColor: "#f5f5f5" };
    }
  },
  methods: {
    onFileChange(event) {
      const file = event.target.files[0];
      if (file) {
        const reader = new FileReader();
        reader.onload = (e) => {
          this.coverBase64 = e.target.result;
        };
        reader.readAsDataURL(file);
      }
    },
    async saveCategory() {
      try {
        const response = await this.$axios({
          method: "post",
          url: "/createCategory",
          data: {
            // userID: userData.user_id,
            category_name: this.name,
            description: this.description,
            cover: this.coverBase64,  // 将Base64的封面图传给后端
          }
        });

        console.log("response:",response)
        if (response.code === 1000) {
          const categoryId = response.data;
          await this.$router.push({name: "Draw", params: {category_id: categoryId}});
        } else {
          console.log("保存分组失败1:", response.data.msg);
        }
      } catch (error) {
        console.error("保存分租失败2:", error);
      }
    },
    cancel() {
      this.$router.push({ name: "Gallery" });
    }
  },
  beforeDestroy() {
    if (this.coverPreview) {
      URL.revokeObjectURL(this.coverPreview);
    }
  }
};
</script>

<style scoped lang="less">
.content {
  max-width: 100%;
  box-sizing: border-box;
  display: flex;
  flex-direction: row;
  justify-content: center;
  margin: 48px auto 0;
  padding: 20px 24px;
  background: #6190E8;
  /* fallback for old browsers */
  background: -webkit-linear-gradient(to right, #6190E8, #A7BFE8);
  /* Chrome 10-25, Safari 5.1-6 */
  background: linear-gradient(to right, #6190E8, #A7BFE8);
  /* W3C, IE 10+/ Edge, Firefox 16+, Chrome 26+, Opera 12+, Safari 7+ */

  .form-container {
    width: 100%;
    max-width: 600px;
    background-color: #fff;
    padding: 40px;
    border-radius: 8px;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);

    h2 {
      font-size: 24px;
      margin-bottom: 20px;
      color: #333;
      text-align: center;
    }

    .category-form {
      .form-group {
        margin-bottom: 20px;

        label {
          display: block;
          font-size: 16px;
          font-weight: 600;
          margin-bottom: 8px;
          color: #555;
        }

        input[type="text"],
        textarea {
          width: 100%;
          padding: 10px;
          font-size: 16px;
          border: 1px solid #ccc;
          border-radius: 4px;
          outline: none;
          transition: border-color 0.3s;

          &:focus {
            border-color: #007bff;
          }
        }

        textarea {
          resize: vertical;
          min-height: 100px;
        }

        .cover-preview {
          width: 150px;
          height: 150px;
          border-radius: 8px;
          background-size: cover;
          background-position: center;
          border: 1px solid #ddd;
          position: relative;
          overflow: hidden;
          display: flex;
          align-items: center;
          justify-content: center;
          cursor: pointer;

          .placeholder-text {
            font-size: 16px; /* 调整字体大小 */
            color: #888;
            text-align: center;
            font-weight: bold; /* 字体加粗 */
          }
        }

        .file-input {
          position: absolute;
          top: 0;
          left: 0;
          width: 100%;
          height: 100%;
          opacity: 0;
          cursor: pointer;
        }
      }

      .buttons {
        display: flex;
        justify-content: space-between;

        .btn {
          padding: 10px 20px;
          font-size: 16px;
          border-radius: 4px;
          cursor: pointer;
          transition: background-color 0.3s, transform 0.1s;

          &.primary {
            background-color: #007bff;
            color: #fff;
            border: none;

            &:hover {
              background-color: #0056b3;
            }

            &:active {
              transform: scale(0.95);
              background-color: #004085;
            }
          }

          &.secondary {
            background-color: #f1f1f1;
            color: #555;
            border: 1px solid #ccc;

            &:hover {
              background-color: #e1e1e1;
            }

            &:active {
              transform: scale(0.95);
              background-color: #d1d1d1;
            }
          }
        }
      }
    }
  }
}
</style>
