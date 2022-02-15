<template>
<form :action="sendMessage" @click.prevent="onSubmit">
  <input v-model="message" type="text">
  <input type="submit" value="Send" @click="sendMessage">
</form>
  <p>
    Two way data binding is fun!
  </p>
  <p>
    received message:
    {{rcvMessage}}
  </p>
  <p>
  </p>
</template>

<script>

export default {
  name: 'App',
  data(){
    return {
      message:"",
      socket: null,
      rcvMessage:""
    }
  },
  mounted() {
    this.socket = new WebSocket("ws://localhost:9100/socket")
        this.socket.onmessage = (msg) => {
          this.rcvMessage = msg.data
        }
  },
  methods: {
    sendMessage(){
      let msg = {
        "greeting":this.message
      }
          this.socket.send(JSON.stringify(msg))
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
