<template>
  <router-view></router-view>
</template>
<script setup>
import toast from "@utils/toast";
toast.loading();
setTimeout(() => {
  toast.clear();
}, 2000);
</script>