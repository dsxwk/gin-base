<template>
  <el-main>main</el-main>
  <el-footer>
    <Footer />
  </el-footer>
</template>
<script setup>
import Footer from '@views/layouts/footer/index.vue';
</script>
<style lang="scss" scoped>
.el-main {
  box-sizing: border-box;
  padding: 10px 12px;
  overflow-x: hidden;
  background-color: var(--el-bg-color-page);
}
</style>