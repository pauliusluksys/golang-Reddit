export default {
    namespaced: true,
    state: {
        posts: [
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
            // console.log("works until here")
            // console.log((localStorage.getItem('JWT')))
            let token = "Bearer " + localStorage.getItem('JWT')
            console.log(token)
            fetch("http://localhost:9100/api/auth/posts", {
                method: "GET",
                headers: {
                    Authorization: token
                },
            }).then(response => {
                    if (response.ok) {
                        // console.log(response.json());
                        return response.json();

                    } else {
                        throw Error(response.body);
                    }
                })
                .then(data => {
                    console.log("data");
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
        getTokenHeader() {
            console.log(localStorage.getItem('JWT'))
            return "Bearer " + localStorage.getItem('JWT');
        }
    }
};