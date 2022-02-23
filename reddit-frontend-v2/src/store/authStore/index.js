export default {
    namespaced: true,
    state: {
        user: {
            id: 1,
            email: 2,
            token:"hello",
            //here there will be the logic for auth and so on...
            loggedIn: false
        }
    },
    mutations: {
        LOGIN(state) {
            state.user.loggedIn = true;
            state.user.username = email;
            state.user.token = token;
        },
        LOGOUT(state) {
            state.user.loggedIn = false;
            state.user.username = "";
            state.user.token = "";
        }
    },
    actions: {
        async login(context, { email, password }) {
            return fetch("http://localhost:9100/api/auth/login", {
                method: "POST",
                body: JSON.stringify({
                    email: email,
                    password: password })
            })
                .then(response => {
                    if (!response.ok) {
                        throw new Error("Cannot login!");
                    }
                    return response.json();
                }).then(data => {
                    context.commit("LOGIN",
                        { email: email, token: data.token });
                }).catch(error => {
                    context.commit("LOGOUT");
                    throw error;
                });
        },
        async logout(context) {
            context.commit("LOGOUT");
        },
        async signup(context, { username, password }) {
            return fetch("http://localhost:9100/api/auth/signup", {
                method: "POST",
                body: JSON.stringify(
                    { username: username, password: password })
            }).then(response => {
                if (!response.ok) {
                    throw new Error("Cannot signup!");
                }
                return response.json();
            }).then(data => {
                context.commit("LOGIN",
                    { username: username, token: data.token });
            }).catch(error => {
                context.commit("LOGOUT");
                error.read().then((data) => {
                    throw Error(data);
                });
            });
        },
    },
    getters: {
        currentUser(state) {
            return state.user;
        },
        isLoggedIn(state) {
            if (!state.user) return false;
            return state.user.loggedIn;
        }
    }
};