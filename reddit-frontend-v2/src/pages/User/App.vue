<template>
  <div class="posts">
    <post-bar></post-bar>
    <h1>Posts</h1>
    <post-list :posts="posts"></post-list>
  </div>
</template>
<script>
import PostList from "../../components/PostList.vue";
import PostBar from "../../components/PostBar.vue";
export default {

  components: {    PostList,PostBar  },
  data: function() {
    return {
      posts: [
        {
          id: 1,
          user: "mycurrentuser",
          date: "2021-01-19 15:30:30",
          post:"Today I'm feeling sooooo well.",
          comments: [
            {
              id: 3,
              user: "Nostradamus",
              date: "2021-01-20 20:30:34",
              post: "LOL"
            },
            {
              id: 4,
              user: "FinnishMan",
              date: "2021-01-20 20:30:34",
              post: "Please..."
            }
          ]
        },
        {
          id: 2,
          user: "cshannon",
          date: "2021-01-19 15:25:20",
          post: "I think that I love giraffes"
        }
      ]

    };
  }};
</script>