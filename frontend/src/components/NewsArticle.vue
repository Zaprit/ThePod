<template>
    <div class="news-article" @click="onLinkClick">
        <h4 style="display: inline; margin-top: 5px;">{{ article!.title }}</h4>
        <img class="article-thumbnail" :src="imgURL" :alt="imgAlt"/>
        <br/>
        <p style="justify-content: left; ">{{ article!.description }}</p>
    </div>
</template>

<script lang="ts" setup>
import {BrowserOpenURL} from "../../wailsjs/runtime";
import {news} from "../../wailsjs/go/models";
import {ref} from "vue";


const props = defineProps({
    article: news.Item
});

let imgURL = ref(props.article?.image?.url);
let imgAlt = ref(props.article?.image?.title);

function onLinkClick() {
    BrowserOpenURL(props.article?.link!);
}

if (imgURL.value == null) {
    const enclosures = props.article?.enclosures;
    if (enclosures == undefined){
        imgURL = ref("nope");
    } else {
        imgURL = ref(enclosures[0].url);
        imgAlt = ref("article thumbnail");
    }
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
    display: block;
    flex-direction: column;
    margin-right: 40px;
    background: #36363b;
    line-height: 1.4;
    font-family: sans-serif;
    border-radius: 5px;
    overflow: hidden;
    z-index: 0;
    margin-bottom: 5px;
    padding-top: 5px;
    user-select: none;
    cursor: pointer;
}
</style>