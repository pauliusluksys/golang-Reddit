import { createStore, createLogger } from 'vuex'
import postsModule from "../pages/Posts/store/postStore";
import usersModule from "./userStore/index.js";
import AuthModule from "./authStore/index.js";
const debug = process.env.NODE_ENV !== 'production'

export default createStore({
    modules: {
        posts: postsModule,
        users: usersModule,
        auth: AuthModule,
    },
    strict: debug,
    plugins: debug ? [createLogger()] : []
})

