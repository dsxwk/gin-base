<template>
  <el-container class="layout">
    <el-aside>
      <div class="aside-box" :style="{ width: isCollapse ? '65px' : '210px' }">
        <div class="logo flx-center">
          <img class="logo-img" src="/favicon.ico" alt="logo"/>
          <span v-show="!isCollapse" @show-span="showSpan" class="logo-text">{{ funcs.lang('Manage Backend') }}</span>
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
        <NavLeft :isCollapse="isCollapse" @toggle-collapse="handleToggleCollapse"/>
        <NavRight/>
      </el-header>
      <Main/>
    </el-container>
  </el-container>
</template>
<script setup>
import { useRoute } from 'vue-router';
import NavLeft from '@/layouts/nav/left/index.vue';
import NavRight from '@/layouts/nav/right/index.vue';
import Main from '@/layouts/main/index.vue';
import Sidebar from '@/layouts/sidebar/menus.vue';
import { menuJson } from '@/utils/data/menu';
import {computed, ref} from 'vue';
import Function from '@/utils/functions';

const funcs = new Function();
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
@import './index.scss';
</style>