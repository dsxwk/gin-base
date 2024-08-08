<template>
  <el-breadcrumb :separator-icon="ArrowRight">
    <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
    <el-breadcrumb-item>用户管理</el-breadcrumb-item>
    <el-breadcrumb-item :to="{ path: '/user' }">用户列表</el-breadcrumb-item>
  </el-breadcrumb>
</template>
<script setup>
import { ArrowRight } from '@element-plus/icons-vue';
</script>
<style lang="scss" scoped>
@import './index.scss';
</style>