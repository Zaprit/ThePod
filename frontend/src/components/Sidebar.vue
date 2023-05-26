<script lang="ts" setup>
import {Server} from "../types/server";
import SidebarServerItem from "./SidebarServerItem.vue";

let servers = [
    new Server("https://static.wixstatic.com/media/5e70b0_50c2f82fe3bf4d4b893a8746aa7da5b5~mv2.png/v1/fill/w_636,h_592,al_c,q_90,usm_0.66_1.00_0.01,enc_auto/Seal%202021.png", "LBPUnion Seal", "Beacon", true),
    new Server("https://crashhelper.lbpunion.com/static/img/CrashHelper.png", "LBPUnion Seal", "Beacon Dev", true)
]

</script>

<template>
  <div class="sidebar">
    <SidebarServerItem v-for="server in servers" :server="server"/>
  </div>
</template>

<style scoped>
  .sidebar {
      height: 100%;
      width: 15em;
      position: absolute;
      left: 0;
      top: 0;
      overflow: auto;
      padding-top: 40px;
      background-color: lightblue;
      resize: vertical;
  }
  .sidebar div {
      padding: 8px;
      font-size: 24px;
      display: block;
  }
</style>