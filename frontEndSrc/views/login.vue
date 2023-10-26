<template>
  <div class="login">
    <div class="login-form">
      <h1>Login</h1>
      <el-form :rules="formRules" :model="form" ref="formName">
        <el-form-item prop="account">
          <el-input
            class="input-account"
            placeholder="Account"
            prefix-icon="el-icon-user-solid"
            v-model="form.account"
          >
          </el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input
            class="input-password"
            type="password"
            placeholder="Password"
            prefix-icon="el-icon-view"
            @keyup.enter.native="submitForm"
            v-model="form.password"
          >
          </el-input>
        </el-form-item>
        <el-form-item>
          <el-button class="button" type="primary" round @click="submitForm"
            >Login
          </el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
import api from "../api/base.js";
export default {
  name: "Login",
  data() {
    return {
      formRules: {
        account: [{ required: true, message: "Account name", trigger: "blur" }],
        password: [{ required: true, message: "Password", trigger: "blur" }]
      },
      form: {
        account: "",
        password: "",
      },
      code: "",
    };
  },
  created() {},
  methods: {
    submitForm() {
      // refs allows us to obtain a direct reference to a specific DOM element or child component instance after it's mounted
      this.$refs.formName.validate((valid) => {
        if (valid) {
          let sha256 = require("js-sha256").sha256;
          let data = {
            username: this.form.account,
            password: sha256(this.form.password),
          };

          try {
            api.login(data).then((res) => {
              if (res.status) {
                sessionStorage.setItem("token", res.data);
                sessionStorage.setItem("user", data.username);
                setTimeout(() => {
                  this.$router.push({ path: "/home" });
                });
              }
            });
          } catch (e) {
            throw new Error(e);
          }
        } else {
          return false;
        }
      });
    }
  }
};
</script>

<style scoped lang="scss">
.login {
  width: 100%;
  height: 100%;
  background: rgb(26, 45, 63);
  display: flex;
  justify-content: center;
  align-items: center;

  h1 {
    text-align: center;
    font-size: 32px;
    line-height: 90px;
    letter-spacing: 5px;
    font-weight: normal;
    color: #555;
    margin-bottom: 10px;
  }

  .button {
    width: 310px;
    height: 40px;
    border-radius: 8px;
    margin-top: 10px;
  }

  .el-form-item {
    margin-right: 0 !important;
  }

  .input-account,
  .input-password {
    width: 310px;
  }

  span {
    display: block;
    width: 100%;
    color: #999;
    font-size: 13px;
    text-align: right;
    cursor: pointer;
  }
}
</style>
