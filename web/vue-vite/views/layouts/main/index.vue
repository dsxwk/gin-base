<template>
  <Breadcrumb/>
  <Tab/>
  <el-main>
    <router-view v-slot="{ component, route }">
      <transition appear name="fade-transform" mode="out-in">
        <keep-alive :include="keepAliveName">
          <component :is="createComponentWrapper(component, route)" v-if="isRouterShow" :key="route.fullPath"/>
        </keep-alive>
      </transition>
    </router-view>
  </el-main>
  <el-footer>
    <Footer/>
  </el-footer>
</template>
<script setup>
import {ref, provide, h} from 'vue';
import Breadcrumb from '@/views/layouts/main/breadcrumb/index.vue';
import Tab from '@/views/layouts/main/tab/index.vue';
import Footer from '@/views/layouts/footer/index.vue';

const keepAliveName = ref([
  '/',
  '/user'
]);
// 注入刷新页面方法
const isRouterShow = ref(true);
const refreshCurrentPage = (val) => (isRouterShow.value = val);
provide("refresh", refreshCurrentPage);

// 解决详情页 keep-alive 问题
const wrapperMap = new Map();
function createComponentWrapper(component, route) {
  console.log(component, route);
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