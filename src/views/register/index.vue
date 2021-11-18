<template>
  <div class="register">
    <el-card>
      <h2>注册</h2>
      <el-form
          class="register-form"
      >
        <el-form-item prop="username">
          <el-input placeholder="用户名" prefix-icon="fas fa-user"></el-input>
        </el-form-item>
        <el-form-item prop="email">
          <el-input placeholder="邮箱" prefix-icon="fas fa-user"></el-input>
        </el-form-item>
        <el-form-item prop="verify_code">
          <el-row  justify="space-around">
            <el-col :span="16">
              <el-input placeholder="验证码" prefix-icon="fas fa-user"></el-input>
            </el-col>
            <el-col :span="6" :offset="1" >
              <el-button class="send_verify_code" type="primary">发送验证码</el-button>
            </el-col>
          </el-row>

        </el-form-item>
        <el-form-item prop="password">
          <el-input
              placeholder="密码"
              type="password"
              prefix-icon="fas fa-lock"
          ></el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input
              placeholder="确认密码"
              type="repassword"
              prefix-icon="fas fa-lock"
          ></el-input>
        </el-form-item>
        <el-form-item>
          <el-button
              :loading="loading"
              class="login-button"
              type="primary"
              native-type="submit"
              block
          >注册</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script>
export default {
  name: "register",
  data() {
    return {
      model: {
        username: "",
        password: ""
      },
      loading: false,
    };
  },
  methods: {
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
