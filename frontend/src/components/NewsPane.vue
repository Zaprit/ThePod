<script setup lang="ts">

import NewsArticle from "./NewsArticle.vue";

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
    <ul style="list-style: none;">
      <news-article />
        <news-article />
        <news-article />
        <news-article />
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