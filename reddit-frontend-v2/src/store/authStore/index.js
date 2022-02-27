export default {
    namespaced: true,
    state: {
        user: {
            id: 1,
            email: "",
            token:"",
            //here there will be the logic for auth and so on...
            loggedIn: false
        }
    },
    mutations: {
        LOGIN(state, { email, token }) {
            state.user.loggedIn = true;
            state.user.email = email;
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
            // console.log(email,password)
            return fetch("http://localhost:9100/api/auth/login", {
                method: "POST",
                body: JSON.stringify({
                    email: email,
                    password: password })
            })
                .then(response => {
                    // console.log(response.json())
                    if (!response.ok) {

                        throw new Error("Cannot login!");
                    }
                    response.json()
                        .then(
                            data => localStorage.setItem('JWT', data.Token),
                            data => context.commit("LOGIN", { email: data.User.email, token: data.Token })

                        )

                }).catch(error => {
                    context.commit("LOGOUT");
                    if (localStorage.getItem('JWT')) {
                        localStorage.removeItem('JWT')
                    }
                    console.log(error);
                    throw error;

                });
        },
        async logout(context) {
            context.commit("LOGOUT");
        },
        async signup(context, { email, password }) {
            return fetch("http://localhost:9100/api/auth/signup", {
                method: "POST",
                body: JSON.stringify(
                    { email: email, password: password })
            }).then(response => {
                if (!response.ok) {
                    throw new Error("Cannot signup!");
                }
                return response.json();
            }).then(data => {
                context.commit("LOGIN",
                    { email: email, token: data.token });
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
        },
        getTokenHeader() {
            console.log("works until here")
            return "Bearer " + localStorage.getItem('JWT');
        },

    }
};