<script lang="ts" setup>
  import {GetSelectedServer} from "../../wailsjs/go/main/App";
  import {server_repo} from "../../wailsjs/go/models";
  import {ref} from "vue";
  import {EventsOn} from "../../wailsjs/runtime";
  import DeviceButton from "./DeviceButton.vue";
  import DeviceMenu from "./DeviceMenu.vue";

  let menuVisible = false;

  const server = ref(await GetSelectedServer());

  function DeviceButtonClick() {
    menuVisible = !menuVisible;
  }

  EventsOn("server_change", async data => {
      server.value = data as server_repo.Server;
  })
</script>

<template>
    <div style="justify-items: center; justify-content: center;" class="patch-bar">
        <device-button @click="DeviceButtonClick"/>
        <device-menu v-show="menuVisible" />

        <button class="patch-button" :style="'background-color: ' + server.accent_colour">Patch!</button>
    </div>
</template>

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