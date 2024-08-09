<template>
  <Breadcrumb/>
  <Tab/>
  <el-main>
    <transition appear name="fade-transform" mode="out-in">
      <keep-alive :include="keepAliveName">
        <router-view :path="route.fullPath"/>
      </keep-alive>
    </transition>
  </el-main>
  <el-footer>
    <Footer/>
  </el-footer>
</template>
<script setup>
import {ref} from 'vue';
import {useRoute} from 'vue-router';
import Breadcrumb from '@/layouts/main/breadcrumb/index.vue';
import Tab from '@/layouts/main/tab/index.vue';
import Footer from '@/layouts/footer/index.vue';
import Home from '@/views/home/index.vue';
import UserIndex from '@/views/user/index.vue';
const route = useRoute();

const keepAliveName = ref([
  Home,
  UserIndex
]);
</script>
<style lang="scss" scoped>
@import './index.scss';
</style>