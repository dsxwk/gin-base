<template>
  <Breadcrumb/>
  <Tab/>
  <el-main>
    <router-view v-slot="{Component, route}">
      <keep-alive>
        <component :is="Component" :key="route.fullPath"/>
      </keep-alive>
    </router-view>
  </el-main>
  <el-footer>
    <Footer/>
  </el-footer>
</template>
<script setup>
import Breadcrumb from '@/layouts/main/breadcrumb/index.vue';
import Tab from '@/layouts/main/tab/index.vue';
import Footer from '@/layouts/footer/index.vue';
</script>
<style lang="scss" scoped>
@import './index.scss';
</style>