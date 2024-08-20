<template>
  <el-dropdown trigger="click">
    <i :title="funcs.lang('Switch Language')" :class="'iconfont icon-zhongyingwen'" class="toolBar-icon"></i>
    <template #dropdown>
      <el-dropdown-menu>
        <el-dropdown-item @click="swichLang('zh-cn')">
          {{ funcs.lang('Chinese') }}
        </el-dropdown-item>
        <el-dropdown-item @click="swichLang('en-us')">
          {{ funcs.lang('English') }}
        </el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
</template>
<script setup>
import { useRoute } from 'vue-router';
import Functions from '@/utils/functions';

const funcs = new Functions();
const route = useRoute();
const swichLang = (lang) => {
  funcs.setCookie('lang', lang);
  // 获取当前路由地址和查询参数
  let currentPath = route.path;
  let currentQuery = route.query;

  // 构造查询参数对象
  let query = new URLSearchParams();
  for (let key in currentQuery) {
    query.set(key, currentQuery[key]);
  }
  query.set('lang', lang);

  // 构造新的 URI 地址
  location.href = `${currentPath}?${query.toString()}`;
}
</script>