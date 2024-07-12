<script setup>
import Nav from '@views/layouts/nav.vue';
import Main from '@views/home/main.vue';
import Sidebar from '@views/layouts/sidebar.vue';
import Footer from '@views/layouts/footer.vue';
</script>
<template>
  <!-- Site wrapper -->
  <div class="wrapper">
    <Nav/>

    <Sidebar/>

    <Main/>

    <Footer/>
  </div>
  <!-- ./wrapper -->
</template>