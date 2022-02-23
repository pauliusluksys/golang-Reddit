<template>
  <base-card :expandable="post.comments && post.comments.length > 0">
    <template v-slot:header>
      <h3>{{ post.Title }}</h3>
      <h5>Written By: {{post.AuthorId}}</h5>
      <p>Created At: {{post.CreatedAt}}</p>
    </template>
    {{ post.post }}
    <template v-slot:footer>
      <base-card v-for="comment in post.comments" :key="comment.id" :expandable="false">
        <template v-slot:header>
          <h3>{{ postTitle(comment) }}</h3>
        </template>
        {{ comment.post }}
      </base-card>
    </template>
  </base-card>
</template>
<script>
import BaseCard from "./UI/BaseCard.vue";
export default {
  components: {BaseCard},
  props: ["post"],

  methods: {
    postTitle(post) {
      return post.user + "@" + post.date;
    }
  }
};
</script>