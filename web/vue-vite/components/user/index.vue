<template>
  用户数据
</template>
<script setup>
import {onMounted} from 'vue';
import userModule from '@/app/modules/admin/user';
import createService from '@/utils/service';
import request from '@/utils/request';

const userService = createService(userModule, request);

onMounted(async () => {
  await userService.list({page:1, pageSize: 10});
})
</script>