<template>
  <div class="register">
    <el-card>
      <h2>注册</h2>
      <el-form
          class="register-form"
      >
        <el-form-item prop="username">
          <el-input :model="form.username" placeholder="用户名" prefix-icon="fas fa-user"></el-input>
        </el-form-item>
        <el-form-item prop="email">
          <el-input :model="form.email" placeholder="邮箱" prefix-icon="fas fa-user"></el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input
              placeholder="密码"
              type="password"
              :model="form.password"
              prefix-icon="fas fa-lock"
          ></el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input
              placeholder="确认密码"
              type="repassword"
              :model="form.repassword"
              prefix-icon="fas fa-lock"
          ></el-input>
        </el-form-item>
        <el-form-item>
          <el-button
              :loading="loading"
              class="login-button"
              type="primary"
              native-type="submit"
              @submit="submitForm"
              block
          >注册</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script>
import request from '../../utils/request'

export default {
  name: "register",
  data() {
    return {
      form: {
        username: '',
        password: '',
        email: '',
        repassword: ''
      },
      loading: false,
    };
  },
  methods: {
    submitForm: function(e) {
      if (this.form.repassword !== this.form.password) {
        this.$message.error("两次输入密码不一致");
        return;
      }
      request({
        url: '/user/register',
        method: 'post',
        params: this.form
      }).then(function (response) {
        const msg = response.data
        if (response.status === 200) {
          this.$message.success(msg);
        } else {
          this.$message.error(msg);
        }
      }).catch(e){
        this.$message.error(e)
      }
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.register {
  flex: 1;
  display: flex;
  justify-content: center;
  align-items: center;
  margin-top: 20px;
}

.login-button {
  width: 100%;
  margin-top: 40px;
}
.register-form {
  width: 390px;
}
.forgot-password {
  margin-top: 10px;
}
.send_verify_code{

}
</style>
