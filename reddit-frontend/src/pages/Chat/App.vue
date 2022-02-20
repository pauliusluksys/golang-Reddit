<template>
  <div id="app">
    <h1>Basic form validation with Vuelidate</h1>
    <form-component />
  </div>
</template>

<script>
import FormComponent from "../../components/LoginFormComponent.vue";

export default {
  name: "App",
  components: {
    FormComponent
  }
};
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
