<template>
  <Breadcrumb/>
  <Tab/>
  <el-main>
    <router-view v-slot="{Component, route}">
      <transition appear name="fade-transform" mode="out-in">
        <keep-alive :include="keepAliveName">
          <component v-if="isRouterShow" :is="createComponentWrapper(Component, route)" :key="route.fullPath"/>
        </keep-alive>
      </transition>
    </router-view>
    <!-- <transition appear name="fade-transform" mode="out-in">
      <keep-alive :include="keepAliveName">
        <router-view :path="route.fullPath"/>
      </keep-alive>
    </transition> -->
  </el-main>
  <el-footer>
    <Footer/>
  </el-footer>
</template>
<script setup>
import {ref, provide, h} from 'vue';
// import {useRoute} from 'vue-router';
import Breadcrumb from '@/layouts/main/breadcrumb/index.vue';
import Tab from '@/layouts/main/tab/index.vue';
import Footer from '@/layouts/footer/index.vue';
import Home from '@/views/home/index.vue';
import UserIndex from '@/views/user/index.vue';
// const route = useRoute();

const keepAliveName = ref([
  Home,
  UserIndex
]);
// 注入刷新页面方法
const isRouterShow = ref(true);
const refreshCurrentPage = (val) => (isRouterShow.value = val);
provide("refresh", refreshCurrentPage);

// 解决详情页 keep-alive 问题
const wrapperMap = new Map();
function createComponentWrapper(component, route) {
  if (!component) {
    return;
  }
  const wrapperName = route.fullPath;
  let wrapper = wrapperMap.get(wrapperName);
  if (!wrapper) {
    wrapper = {name: wrapperName, render: () => h(component)};
    wrapperMap.set(wrapperName, wrapper);
  }
  return h(wrapper);
}
</script>
<style lang="scss" scoped>
@import './index.scss';
</style>