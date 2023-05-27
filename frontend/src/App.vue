<script lang="ts" setup>

import {SidebarHeaderItem, SidebarItem, SidebarMenu} from "vue-sidebar-menu";
import {ref} from "vue";
import NewsPane from "./components/NewsPane.vue";
import PatchingBar from "./components/PatchingBar.vue";

const menu = ref([
    {
        header: 'Servers',
        hiddenOnCollapse: true
    } as SidebarHeaderItem,
    {
        title: 'Beacon',
        icon: {
            element: "img",
            attributes: {src: "https://static.wixstatic.com/media/5e70b0_50c2f82fe3bf4d4b893a8746aa7da5b5~mv2.png/v1/fill/w_636,h_592,al_c,q_90,usm_0.66_1.00_0.01,enc_auto/Seal%202021.png"}
        }
    } as SidebarItem,

] as SidebarMenu['menu'])

menu.value.push({
    title: 'Beacon Dev',

    icon: {
        element: "img",
        attributes: {src: "https://crashhelper.lbpunion.com/static/img/CrashHelper.png"}
    }
} as SidebarItem)



// const selectedServer = reactive(GetServer())

function Fard(event: Event, item: any) {
    console.log(item)
}

</script>

<template>

    <sidebar-menu :menu="menu" hide-toggle @click="fard" />

    <news-pane background="https://static.wixstatic.com/media/5e4ff2_593104cd3e264e93b92bbe7f2e745535~mv2.jpg/v1/fill/w_1549,h_967,al_c,q_85,enc_auto/5e4ff2_593104cd3e264e93b92bbe7f2e745535~mv2.jpg"
    icon="https://static.wixstatic.com/media/5e70b0_50c2f82fe3bf4d4b893a8746aa7da5b5~mv2.png/v1/fill/w_636,h_592,al_c,q_90,usm_0.66_1.00_0.01,enc_auto/Seal%202021.png"
    />

    <patching-bar style="position: absolute; bottom: 0; width: 100%; height: 8%;" accent-colour="blue"/>
</template>

<style>

</style>
