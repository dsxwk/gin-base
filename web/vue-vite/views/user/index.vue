<template>
  <Layout>
    <template v-slot:main>
      <el-main>
        <UserList/>
      </el-main>
      <el-footer>
        <Footer />
      </el-footer>
    </template>
  </Layout>
</template>
<script setup>
import Layout from '@/views/layouts/index.vue';
import UserList from '@/components/user/index.vue';
import Footer from '@/views/layouts/footer/index.vue';
</script>
<style lang="scss" scoped>
@import '@/views/layouts/main/index.scss';
</style>