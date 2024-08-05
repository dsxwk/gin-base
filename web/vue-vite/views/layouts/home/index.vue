<template>
  <Layout>
    <template v-slot:main>
      <el-main>
        <Home/>
      </el-main>
    </template>
  </Layout>
</template>
<script setup>
import Layout from '@layouts/index.vue';
import Home from '@components/home/index.vue';
</script>