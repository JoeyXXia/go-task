<script setup>
import loginAccount from "./login-account.vue";
</script>

<template>
  <div class="login-panel">
    <!-- title -->
    <h1 class="title">Go Tasks 登录系统</h1>

    <!-- tabs -->
    <div class="tabs">
      <el-tabs type="border-card">
        <el-tab-pane>
          <template #label>
            <div class="label">
              <el-icon><UserFilled /></el-icon>
              <span class="text">账号登录</span>
            </div>
          </template>
          <login-account />
        </el-tab-pane>

        <el-tab-pane>
          <template #label>
            <div class="label">
              <el-icon><Cellphone /></el-icon>
              <span class="text">手机登录</span>
            </div>
          </template>
        </el-tab-pane>
      </el-tabs>
    </div>

    <!-- bottom -->
    <div class="controls">
      <el-checkbox v-model="checked1" label="记住密码" size="large" />
      <el-link type="info">忘记密码</el-link>
    </div>
    <el-button class="login-btn" type="primary" size="large">登录</el-button>
  </div>
</template>

<style lang="less" scoped>
.login-panel {
  width: 330px;
  margin-bottom: 150px;

  .title {
    text-align: center;
    margin-bottom: 15px;
  }

  .label {
    display: flex;
    align-items: center;
    justify-content: center;

    .text {
      margin-left: 6px;
    }
  }

  .controls {
    margin-top: 12px;
    display: flex;
    justify-content: space-between;
  }

  .login-btn {
    margin-top: 10px;
    width: 100%;
  }
}
</style>
