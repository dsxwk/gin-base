<template>
  <div class="fullscreen">
    <i :title="funcs.lang('Full Screen')" :class="['iconfont', isFullscreen ? 'icon-suoxiao' : 'icon-fangda']" class="toolBar-icon" @click="handleFullScreen"></i>
  </div>
</template>
<script setup>
import { onMounted, ref } from 'vue';
import screenfull from 'screenfull';
import pnotify from '@/utils/pnotify/alert';
import Functions from '@/utils/functions';

const funcs = new Functions();
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
    pnotify(funcs.lang('Mobile devices do not support full screen') + ' ❌', 'error');
    return;
  } else {
    if (!screenfull.isEnabled) {
      pnotify(funcs.lang('Current browser does not support full screen') + ' ❌', 'info');
    }
    screenfull?.toggle();
  }
};
</script>