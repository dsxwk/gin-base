<template>
  <el-container class="layout">
    <el-aside>
      <div class="aside-box" :style="{ width: isCollapse ? '65px' : '210px' }">
        <div class="logo flx-center">
          <img class="logo-img" src="/favicon.ico" alt="logo"/>
          <span v-show="!isCollapse" @show-span="showSpan" class="logo-text">后台管理</span>
        </div>
        <el-scrollbar>
          <el-menu
              :router="false"
              :default-active="activeMenu"
              :collapse="isCollapse"
              :unique-opened="false"
              :collapse-transition="false"
          >
            <Sidebar :menuList="menuList"/>
          </el-menu>
        </el-scrollbar>
      </div>
    </el-aside>
    <el-container>
      <el-header>
        <div class="header-lf mask-image">
          <NavLeft :isCollapse="isCollapse" @toggle-collapse="handleToggleCollapse"/>
        </div>
        <div class="header-ri">
          <NavRight/>
        </div>
      </el-header>
      <el-breadcrumb :separator-icon="ArrowRight">
        <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
        <el-breadcrumb-item>用户管理</el-breadcrumb-item>
        <el-breadcrumb-item :to="{ path: '/user' }">用户列表</el-breadcrumb-item>
      </el-breadcrumb>
      <slot name="main">
        <Main/>
      </slot>
      <el-footer>
        <Footer />
      </el-footer>
    </el-container>
  </el-container>
</template>
<script setup>
import { useRoute } from 'vue-router';
import NavLeft from '@layouts/nav/left/index.vue';
import NavRight from '@layouts/nav/right/index.vue';
import Main from '@layouts/main/index.vue';
import Sidebar from '@layouts/sidebar/menus.vue';
import Footer from '@views/layouts/footer/index.vue';
import { menuJson } from '@utils/data/menu';
import { ArrowRight } from '@element-plus/icons-vue';

import {computed, ref} from 'vue';

const isCollapse = ref(false);
const route = useRoute();
const menuList = computed(() => (menuJson));
const activeMenu = computed(() => {
  return route.meta.activeMenu ? route.meta.activeMenu : route.path;
});
function handleToggleCollapse() {
  isCollapse.value = !isCollapse.value;
}
function showSpan() {
  isCollapse.value = !isCollapse.value;
}
</script>
<style scoped lang="scss">
@import "./index.scss";
@import '@layouts/main/index.scss';
.el-breadcrumb {
  height: 55px;
  line-height: 55px;
  padding-left: 12px;
  font-size: 15px;
  border-bottom: 1px solid #e6e6e6;
}
</style>