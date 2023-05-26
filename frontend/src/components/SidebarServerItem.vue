<script lang="ts" setup>
import {Server} from "../types/server";

defineProps({
    server: {
        type: Server,
        required: true
    }
})
</script>

<template>
  <div class="sidebar-list-item">
    <img style="vertical-align: middle;" id="server-icon" :src="$props.server.iconURL" :alt="$props.server.iconAlt" />
    {{ $props.server.serverName }}
    <span v-if="$props.server?.upStatus">🟢</span>
    <span v-else>🔴</span>
  </div>
</template>

<style scoped>
  #server-icon {
      width: 32px;
  }
  .sidebar-list-item {
  }
</style>