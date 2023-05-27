<template>
    <li class="news-article" style="display: block; padding-top: 5px; user-select: none; cursor: pointer" @click="onLinkClick">

        <h4 style="display: inline; margin-top: 5px;">LBP Union's Response to the Enhanced Individual Crisis</h4>
        <img class="article-thumbnail" src="https://static.wixstatic.com/media/5e70b0_66f18f3e2efe4ea398da143211aadd32~mv2.png/v1/fill/w_454,h_341,fp_0.50_0.50,q_95,enc_auto/5e70b0_66f18f3e2efe4ea398da143211aadd32~mv2.png" alt="" />

        <br/>

        <p style="justify-content: left; ">Tile Denial reveals the LBP Union's plan to counter the spread of enhanced individuals across the planet.</p>

    </li>
</template>

<script lang="ts" setup>
import {BrowserOpenURL} from "../../wailsjs/runtime";

function onLinkClick() {
    BrowserOpenURL("https://www.lbpunion.com/post/lbp-union-s-response-to-the-enhanced-individual-crisis");
}
</script>

<style scoped>
.article-thumbnail {
    display: inline;
    float: left;
    width: 20%;
    border-radius: 5px;
    margin: 5px;
    box-shadow: 0 14px 80px rgba(34, 35, 58, 0.5);
}
.news-article {
    display: flex;
    flex-direction: column;
    margin-right: 40px;
    background: #36363b;
    line-height: 1.4;
    font-family: sans-serif;
    border-radius: 5px;
    overflow: hidden;
    z-index: 0;
    margin-bottom: 5px;
}
</style>