<template>
  <div class="fullscreen">
    <i title="全屏" :class="['iconfont', isFullscreen ? 'icon-suoxiao' : 'icon-fangda']" class="toolBar-icon" @click="handleFullScreen"></i>
  </div>
</template>
<script setup>
import { onMounted, ref } from 'vue';
import screenfull from 'screenfull';
const isFullscreen = ref(screenfull.isFullscreen);

onMounted(() => {
  screenfull.on("change", () => {
    if (screenfull.isFullscreen) {
      isFullscreen.value = true;
    } else {
      isFullscreen.value = false;
    }
  });
});

const handleFullScreen = () => {
  if (!screenfull.isEnabled) {
    pnotify('当前您的浏览器不支持全屏 ❌', 'info');
  }
  screenfull.toggle();
};
</script>