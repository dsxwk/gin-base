<template>
  <el-breadcrumb :separator-icon="ArrowRight">
<!-- <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
    <el-breadcrumb-item>用户管理</el-breadcrumb-item>
    <el-breadcrumb-item :to="{ path: '/user' }">用户列表</el-breadcrumb-item> -->
    <el-breadcrumb-item v-for="(item, index) in breadcrumbList" :key="index" :to="item.path">
      {{ item.name }}
    </el-breadcrumb-item>
  </el-breadcrumb>
</template>
<script setup>
import { computed } from 'vue';
import { useRoute } from 'vue-router';
import {menuJson} from '@/utils/data/menu/index.js';
import { ArrowRight } from '@element-plus/icons-vue';
import {HOME_URL} from '@/config';

// 获取当前路由信息
const route = useRoute();

// 递归查找匹配的路由
const findBreadcrumb = (menu, path, breadcrumb = []) => {
  for (const item of menu) {
    // 如果路径匹配，加入面包屑
    if (item.path === path) {
      breadcrumb.push({
        name: item.meta.title || item.name,
        path: item.path
      });
      return breadcrumb;
    }
    // 如果有子菜单，递归查找
    if (item.children && item.children.length > 0) {
      const result = findBreadcrumb(item.children, path, [...breadcrumb, {
        name: item.meta.title || item.name,
        path: item.path
      }]);
      if (result) return result;
    }
  }
  return null;
};

const breadcrumbList = computed(() => {
  const matchedRoutes = route.matched;
  const breadcrumbs = [];

  // 如果当前路径是首页
  if (route.path === HOME_URL) {
    return [{
      name: '首页',
      path: HOME_URL
    }];
  }

  // 始终有首页
  breadcrumbs.push({
    name: '首页',
    path: HOME_URL
  });

  // 从匹配的路由中生成面包屑
  for (const route of matchedRoutes) {
    const breadcrumb = findBreadcrumb(menuJson, route.path);
    if (breadcrumb) {
      breadcrumbs.push(...breadcrumb);
    }
  }

  return breadcrumbs;
});
</script>
<style lang="scss" scoped>
@import './index.scss';
</style>