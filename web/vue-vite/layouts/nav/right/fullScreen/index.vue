<template>
  <div class="fullscreen">
    <i title="全屏" :class="['iconfont', isFullscreen ? 'icon-suoxiao' : 'icon-fangda']" class="toolBar-icon" @click="handleFullScreen"></i>
  </div>
</template>
<script setup>
import { onMounted, ref } from 'vue';
import screenfull from 'screenfull';
import pnotify from '@/utils/pnotify/alert';
const isFullscreen = ref(screenfull.isFullscreen);
const isMobile = ref(false);

onMounted(() => {
  checkIfMobile();
  if (!isMobile.value) {
    screenfull.on("change", () => {
      if (screenfull.isFullscreen) {
        isFullscreen.value = true;
      } else {
        isFullscreen.value = false;
      }
    });
  }
});

const checkIfMobile = () => {
  isMobile.value = /Mobi|Android/i.test(navigator.userAgent);
};

const handleFullScreen = () => {
  if (isMobile.value) {
    pnotify('移动端不支持全屏 ❌', 'error');
    return;
  } else {
    if (!screenfull.isEnabled) {
      pnotify('当前您的浏览器不支持全屏 ❌', 'info');
    }
    screenfull?.toggle();
  }
};
</script>