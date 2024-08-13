<template>
  用户数据
</template>
<script setup>
import { ref, onMounted } from 'vue';
import userModule from '@/app/modules/admin/user/index.js';
import createService from '@/utils/service.js';

const data = ref([]);
const userService = createService(userModule);
onMounted(async () => {
  data.value = await userService.list({page: 1, pageSize: 10});
});
</script>