<script setup>
import loginPanel from "./cpns/login-panel.vue";
</script>

<template>
  <div class="login">
    <login-panel />
  </div>
</template>

<style lang="less" scoped>
.login {
  display: flex;
  align-items: center;
  justify-content: center;

  width: 100%;
  height: 100%;
}
</style>
