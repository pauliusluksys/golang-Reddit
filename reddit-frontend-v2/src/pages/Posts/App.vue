<template>
  <div class="posts">
    <post-bar></post-bar>
    <h1>Posts</h1>
    <p></p>
    <post-list :posts="posts"></post-list>
        <ul v-for="post in posts" :key="post.ID">
      <li>{{ post }}
      </li>
        </ul>

<!--    <post-list :posts="posts"></post-list>-->
  </div>
</template>
<script>
import PostList from "../../components/PostList.vue";
import PostBar from "../../components/PostBar.vue";
import {mapState} from "vuex";

export default {
  mounted(){
    console.log((localStorage.getItem('JWT')));
    this.$store.dispatch("posts/getAllPosts");
    console.log('mounted!')
  },
  components: {    PostBar,PostList  },

  computed: mapState({
    posts: state => state.posts.posts.post
  }),
  data: function() {
    return {
    };
  }};
</script>