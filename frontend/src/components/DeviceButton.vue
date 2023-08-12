<template>
    <div v-if="haveDevices">
        <button>{{dev.name}}</button>
    </div>
    <div v-else>
        <span>No Devices</span>
    </div>
</template>

<script lang="ts" setup>

import {GetSelectedDevice} from "../../wailsjs/go/main/App";
import {remote} from "../../wailsjs/go/models";

let dev = await GetSelectedDevice();
let haveDevices = true;
if (dev as remote.Device | null == null) {
    haveDevices = false;
}

</script>

<style scoped>

</style>