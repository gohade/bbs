<template>
  <div class="login">
    <el-card>
      <h2>登录</h2>
      <el-form
          class="login-form"
      >
        <el-form-item prop="username">
          <el-input placeholder="用户名" prefix-icon="fas fa-user"></el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input
              placeholder="密码"
              type="password"
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
          >登录</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script>
export default {
  name: "login",
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
.login {
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
.login-form {
  width: 290px;
}
.forgot-password {
  margin-top: 10px;
}
</style>
