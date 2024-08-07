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
import Layout from '@/views/layouts/index.vue';
import UserList from '@/components/user/index.vue';
</script>