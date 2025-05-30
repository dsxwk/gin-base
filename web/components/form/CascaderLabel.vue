<script setup>
defineProps({
  node: Object,
  data: Object
});
</script>
<template>
  <span>
    {{ data.title }}
    <span v-if="!node.isLeaf"> ({{ data.children?.length || 0 }})</span>
  </span>
</template>