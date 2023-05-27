<template>
    <div style="justify-items: center; justify-content: center;" class="patch-bar">
  <device-button device-name="TestDevice" />

  <button class="patch-button" :style="'background-color: ' + accentColour">Patch!</button>
    </div>
</template>

<script lang="ts" setup>

import DeviceButton from "./DeviceButton.vue";
defineProps({
    "accentColour": String
})
</script>

<style scoped>
  .patch-bar {
      background-color: #2a2a2e;
  }
  .patch-button {
      border-radius: 50px;
      font-size: 45px;
      color: #36363b;
  }
</style>