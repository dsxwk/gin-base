<template>
  <ElConfigProvider :locale="funcs.getLang() !== 'zh-cn' ? en : zhCn">
    <router-view></router-view>
  </ElConfigProvider>
</template>
<script setup>
import {ref} from 'vue';
import zhCn from 'element-plus/es/locale/lang/zh-cn';
import en from 'element-plus/es/locale/lang/en';
import Functions from '@/utils/functions';

const funcs = ref(new Functions());
</script>