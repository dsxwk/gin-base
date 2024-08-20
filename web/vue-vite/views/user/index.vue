<template>
  {{ funcs.lang('User') }}
</template>
<script setup>
import { ref, onMounted } from 'vue';
import userModule from '@/app/modules/admin/user';
import createService from '@/utils/service';
import Functions from '@/utils/functions';

const funcs = new Functions();
const data = ref([]);
const userService = createService(userModule);
onMounted(async () => {
  data.value = await userService.list({page: 1, pageSize: 10});
});
</script>