<template>
  <el-main>
    <Home/>
  </el-main>
  <el-footer>
    <Footer />
  </el-footer>
</template>
<script setup>
import Home from '@/components/home/index.vue';
import Footer from '@/views/layouts/footer/index.vue';
</script>
<style lang="scss" scoped>
@import './index.scss';
</style>