import { createStore, createLogger } from 'vuex'
import posts from './postStore/index'

const debug = process.env.NODE_ENV !== 'production'

export default createStore({
    modules: {
        posts
    },
    strict: debug,
    plugins: debug ? [createLogger()] : []
})