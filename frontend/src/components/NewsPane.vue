<script lang="ts" setup>

import NewsArticle from "./NewsArticle.vue";
import {GetArticles, GetSelectedServer} from "../../wailsjs/go/main/App";
import {EventsOn, LogInfo} from '../../wailsjs/runtime';
import {ref} from "vue";

let server = ref(await GetSelectedServer());

let articles = ref(await GetArticles());

EventsOn("server_change", async newServer => {
    server.value = newServer;
    articles.value = await GetArticles();
});

</script>

<template>
  <div style="display: block; padding-left: 290px;">
    <img class="news-background" :src="server.background_image" alt=""/>
    <img class="news-icon" :src="server.icon_url" alt=""/>

    <h2>News</h2>
    <ul style="list-style: none;">
      <li v-for="item in articles" :key="item.guid">
        <news-article :article=item />
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