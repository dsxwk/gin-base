<template>
  <Nav/>

  <Sidebar/>

  <Main/>

  <Footer/>
</template>
<script setup>
import Nav from '@views/layouts/nav/index.vue';
import Main from '@views/user/index.vue';
import Sidebar from '@views/layouts/sidebar/SubMenu.vue';
import Footer from '@views/layouts/footer/index.vue';
</script>