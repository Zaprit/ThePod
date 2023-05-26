<script setup lang="ts">

import {GetArticles} from "../../wailsjs/go/news/ArticleFetcher";

defineProps({
    icon: String,
    background: String,
})

// GetArticles("").then(articles => {
//     articles.forEach(article => {
//         console.log(article.Title);
//     });
// });
</script>

<template>
  <div style="display: block; padding-left: 290px;">
    <img class="news-background" :src="background" alt=""/>
    <img class="news-icon" :src="icon" alt=""/>

      <h2>News</h2>
    <ul style="list-style: none">
        <li>
            <img style="width: 10%;" src="https://static.wixstatic.com/media/5e70b0_66f18f3e2efe4ea398da143211aadd32~mv2.png/v1/fill/w_454,h_341,fp_0.50_0.50,q_95,enc_auto/5e70b0_66f18f3e2efe4ea398da143211aadd32~mv2.png" alt="" />
            <span>LBP Union's Response to the Enhanced Individual Crisis</span>
            <p>Tile Denial reveals the LBP Union's plan to counter the spread of enhanced individuals across the planet.</p>
        </li>
    </ul>
  </div>
</template>

<style scoped>
.news-background {
    filter: blur(8px);
    -webkit-filter: blur(8px);
    position: absolute;
    z-index: -5;
}
.news-icon {
    width: 10%;
  z-index: 10;
}
</style>