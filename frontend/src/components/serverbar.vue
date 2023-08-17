<script setup lang="ts">
import {GetServers} from "../../wailsjs/go/main/App";
import {server_repo} from "../../wailsjs/go/models";

let servers = await GetServers()
let serverList = Array<server_repo.Server>()
for (let serversKey in servers) {
  serverList.push(servers[serversKey])
}
</script>

<template>
<ul :v-for="serverList">

</ul>
</template>

<style scoped>

</style>