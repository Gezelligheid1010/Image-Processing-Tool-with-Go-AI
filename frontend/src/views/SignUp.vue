<template>
  <div class="main">
    <div class="container">
      <h2 class="form-title">注册</h2>
      <div class="form-group">
        <label for="avatar"><span style="color:red;">* </span>头像</label>
        <div class="avatar-upload">
          <div class="avatar-preview">
            <!-- 圆形区域显示上传的图片 -->
            <div :style="avatarStyle" class="avatar-image"></div>
          </div>
          <input type="file" id="avatarUpload" ref="avatar" accept="image/*" @change="onImageChange" class="file-input"/>
        </div>
      </div>
      <div class="form-group">
        <label for="name"><span style="color:red;">* </span>用户名</label>
        <el-input type="text" required name="name" id="name" placeholder="用户名" v-model="username" />
      </div>
      <div class="form-group">
        <label for="email"><span style="color:red;">* </span>邮箱</label>
        <el-input type="email" required name="email" id="email" placeholder="请输入邮箱" v-model="email" />
      </div>
      <div class="form-group">
        <label for="pass"><span style="color:red;">* </span>密码</label>
        <el-input type="password" required name="pass" id="pass" placeholder="密码" v-model="password" />
      </div>
      <div class="form-group">
        <label for="re_pass"><span style="color:red;">* </span>确认密码</label>
        <el-input type="password" required name="re_pass" id="re_pass" placeholder="确认密码" v-model="confirm_password" />
      </div>
      <div class="form-group">
        <label for="gender"><span style="color:red;">* </span>性别</label>
        <div id="gender">
          <el-radio v-model="gender" :label="1">男</el-radio>
          <el-radio v-model="gender" :label="2">女</el-radio>
        </div>
      </div>
      <div class="form-btn">
        <button type="button" class="btn btn-info" @click="submit">提交</button>
      </div>
    </div>
  </div>
</template>

<script>

export default {
  name: "SignUp",
  data() {
    return {
      username: "",
      password: "",
      email: '',
      gender: 1,
      confirm_password: "",
      submitted: false,
      avatarImageUrl: "",  // 用于存储预览的头像图片URL
    };
  },
  computed: {
    avatarStyle() {
      return this.avatarImageUrl
          ? { backgroundImage: `url(${this.avatarImageUrl})` }
          : {};
    },
  },
  created() {

  },

  methods: {
    
    // 处理文件变化并生成预览
    onImageChange(event) {
      const file = event.target.files[0];
      if (file) {
        console.log('File selected:', file);  // 确保文件正确选择
        const reader = new FileReader();
        reader.onload = (e) => {
          console.log('File read successfully');  // 确保 FileReader 正常工作
          this.avatarImageUrl = e.target.result;
          // console.log('onImageChange:',this.avatarImageUrl)
        };
        reader.readAsDataURL(file);
      } else {
        console.log('No file selected');
      }
    },
    submit() {
      if (!this.avatarImageUrl) {
        console.log('No avatar selected');
        return;
      }
      this.submitted = true; // 禁用提交

      // 2. 上传注册信息
      this.$axios({
        method: 'post',
        url: '/signup',
        data: {
          avatar: this.avatarImageUrl, // 仅传递 base64 数据部分
          username: this.username,
          email: this.email,
          gender: this.gender,
          password: this.password,
          confirm_password: this.confirm_password
        },
        headers: { 'Content-Type': 'application/json' }
      })
          .then((res) => {
            // console.log(res)
            this.submitted = false; // 请求完成，重新启用
            if (res.code === 1000) {
              console.log('Signup success');
              this.$message.success('注册成功！');
              this.$router.push({ name: "Login" });
            } else {
              console.log(res.data.msg);
            }
          })
          .catch((error) => {
            console.log('Signup error:', error);
            this.submitted = false;
          });

    }
  }
};
</script>
<style lang="less" scoped>
.main {
  background: #6190E8;
  /* fallback for old browsers */
  background: -webkit-linear-gradient(to right, #A7BFE8, #6190E8);
  /* Chrome 10-25, Safari 5.1-6 */
  background: linear-gradient(to right, #A7BFE8, #6190E8);
  /* W3C, IE 10+/ Edge, Firefox 16+, Chrome 26+, Opera 12+, Safari 7+ */

  padding: 150px 0;
  min-height: 60vh;

  .container {
    width: 600px;
    background: #fff;
    margin: 0 auto;
    max-width: 1200px;
    padding: 20px;

    .form-title {
      margin-bottom: 33px;
      text-align: center;
    }

    .form-group {
      margin: 15px;

      label {
        display: inline-block;
        max-width: 100%;
        margin-bottom: 5px;
        font-weight: 700;
      }
    }

    .form-btn {
      display: flex;
      justify-content: center;

      .btn {
        padding: 6px 20px;
        font-size: 18px;
        line-height: 1.3333333;
        border-radius: 6px;
        display: inline-block;
        margin-bottom: 0;
        font-weight: 400;
        text-align: center;
        white-space: nowrap;
        vertical-align: middle;
        -ms-touch-action: manipulation;
        touch-action: manipulation;
        cursor: pointer;
        border: 1px solid transparent;
      }

      .btn-info {
        color: #fff;
        background-color: #5bc0de;
        border-color: #46b8da;
      }
    }


    /* 头像上传样式 */
    .avatar-upload {
      position: relative;
      width: 100px;
      height: 100px;
      margin: 0 auto;
    }
    .avatar-preview {
      width: 100%;
      height: 100%;
      border-radius: 50%;
      background-color: #f0f0f0;
      background-position: center;
      background-size: cover;
      overflow: hidden;
    }
    .avatar-image {
      width: 100%;
      height: 100%;
      border-radius: 50%;
      background-size: cover;
      background-position: center;
    }
    /* 文件输入框样式 */
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
}
</style>