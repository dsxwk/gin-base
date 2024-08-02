<template>
  <Layout>
    <template v-slot:main>
      <el-main>
        <UserList/>
      </el-main>
    </template>
  </Layout>
</template>
<script setup>
import Layout from '@layouts/index.vue';
import UserList from '@components/user/index.vue';
</script>
<style lang="scss" scoped>
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