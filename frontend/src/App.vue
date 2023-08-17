<script lang="ts" setup>

import {SidebarHeaderItem, SidebarItem, SidebarMenu} from "vue-sidebar-menu";
import {ref} from "vue";
import {GetServers, SelectServer, FetchImage} from "../wailsjs/go/main/App";
import NewsPane from "./components/NewsPane.vue";
import PatchingBar from "./components/PatchingBar.vue";

const menu = ref([
    {
        header: 'Servers',
        hiddenOnCollapse: true
    } as SidebarHeaderItem,

] as SidebarMenu['menu']);

GetServers().then(servers => {
    for (const serverKey in servers) {
        menu.value.push({
            title: servers[serverKey].display_name,
            key: serverKey,
            href: "#",
            icon: {
                element: "img",
                attributes: {src: servers[serverKey].icon_url}
            }
        } as SidebarItem);
    }
});



async function SidebarClick(event: Event, item: { key: string; }) {
   await SelectServer(item.key);
}

</script>

<template>
<!--    <sidebar-menu :menu="menu" hide-toggle @item-click="SidebarClick" />-->
  <server-bar></server-bar>
    <suspense>
        <news-pane/>
    </suspense>
    <suspense>
        <patching-bar style="position: absolute; bottom: 0; width: 100%; height: 8%; border-top: #bbc5d6 2px;"/>
    </suspense>
</template>

<style>

</style>
