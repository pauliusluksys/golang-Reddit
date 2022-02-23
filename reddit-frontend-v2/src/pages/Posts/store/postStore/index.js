export default {
    namespaced: true,
    state: {
        posts: [
            {
                id: 1,
                user: "idalmasso",
                date: "2021-01-19 15:30:30",
                post:
                    "Today I'm feeling sooooo well...",
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
                post: "Say something here"
            }
        ]
    },

    mutations: {
        SET_ALL_POSTS(state, posts) {
            state.posts = posts;
        },
        ADD_POST(state, post) {
            state.posts.push(post);
        },
        DELETE_POST(state, id) {
            state.posts = state.posts.filter(post => post.id != id);
        }
    },
    actions: {
        // async addPost(context, post) {
        //     context.commit("ADD_POST", post);
        // },
        // async deletePost(context, id) {
        //     context.commit("DELETE_POST", id);
        // }


        async getAllPosts(context) {
            fetch("http://localhost:9100/api/auth/posts", {
                method: "GET",
                headers: {
                    "Access-Control-Allow-Origin": "http://localhost:8080",
                    Authorization: context.rootGetters["auth/getTokenHeader"]
                },
            }).then(response => {
                    if (response.ok) {
                        return response.json();
                    } else {
                        throw Error(response.body);
                    }
                })
                .then(data => {
                    console.log(data);
                    context.commit("SET_ALL_POSTS", data);
                })
                .catch(error => {
                    console.log(error);
                });
        }
    },
    getters: {
        allPosts(state) {
            return state.posts;
        },
        userPosts: state => user => {
            return state.posts.filter(post => post.user === user);
        },
        getTokenHeader(state) {
            return "Bearer " + state.user.token;
        }
    }
};