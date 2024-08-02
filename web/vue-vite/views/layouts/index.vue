<template>
  <el-container class="layout">
    <el-header>
      <div class="header-lf mask-image">
        <div class="logo flx-center">
          <img class="logo-img" src="/public/favicon.ico" alt="logo"/>
          <span class="logo-text">后台管理</span>
        </div>
        <NavLeft :isCollapse="isCollapse" @toggle-collapse="handleToggleCollapse"/>
      </div>
      <div class="header-ri">
        <NavRight/>
      </div>
    </el-header>
    <el-container class="classic-content">
      <el-aside>
        <div class="aside-box" :style="{ width: isCollapse ? '65px' : '210px' }">
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
      <el-container class="classic-main">
        <slot name="main">
          <Main/>
        </slot>
        <el-footer>
          <Footer />
        </el-footer>
      </el-container>
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
import menuJson from '@utils/data/menu/index.json';

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
</script>
<style scoped lang="scss">
@import "./index.scss";
.layout {
  min-width: 600px;
}
.el-main {
  box-sizing: border-box;
  padding: 10px 12px;
  overflow-x: hidden;
  background-color: var(--el-bg-color-page);
}
</style>