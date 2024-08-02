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
@import './index.scss';
</style>