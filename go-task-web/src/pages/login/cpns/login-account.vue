<script setup>
import { reactive } from "vue";

// account data
const account = reactive({
  name: "",
  password: "",
});

// validate rules
const accountRules = {
  name: [
    { required: true, message: "必须输入账号信息", trigger: "blur" },
    {
      pattern: /^[a-z0-9]{6,20}$/,
      message: "必须是6-20数字或者字母组成",
      trigger: "blur",
    },
  ],
  password: [
    { required: true, message: "必须输入密码", trigger: "blur" },
    {
      pattern: /^[a-z0-9]{7,}$/,
      message: "必须是7位以上数字或者字母",
      trigger: "blur",
    },
  ],
};
</script>

<template>
  <div class="login-account">
    <el-form
      :model="account"
      :rules="accountRules"
      label-width="55px"
      status-icon
      size="large"
    >
      <el-form-item label="账号" prop="name">
        <el-input v-model="account.name" />
      </el-form-item>
      <el-form-item label="密码" prop="password">
        <el-input v-model="account.password" />
      </el-form-item>
    </el-form>
  </div>
</template>

<style lang="less" scoped></style>
