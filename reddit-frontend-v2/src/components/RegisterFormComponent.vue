<template>
  <div class="form pt-6">
    <div class="summary text-red" v-if="$v.form.$error">
      Form has errors
    </div>
    <form @submit.prevent="submit">
      <div class="flex justify-center my-6">
        <div
            class="px-4"
            :class="{ 'hasError': $v.form.firstName.$error }">
          <label class="mr-2 font-bold text-grey">First Name</label>
          <input type="text" class="input" v-model="form.firstName">
        </div>
        <div
            class="px-4"
            :class="{ 'hasError': $v.form.lastName.$error }">
          <label class="mr-2 font-bold text-grey">Last Name</label>
          <input type="text" class="input" v-model="form.lastName">
        </div>
        <div
            class="px-4"
            :class="{ 'hasError': $v.form.username.$error }">
          <label class="mr-2 font-bold text-grey">Name</label>
          <input type="text" class="input" v-model="form.username">
        </div>
        <div
            class="px-4"
            :class="{ 'hasError': $v.form.email.$error }">
          <label class="mr-2 font-bold text-grey">Email</label>
          <input type="email" class="input" v-model="form.email">
        </div>
      </div>
      <div class="text-center">
        <button type="submit" class="button">
          Submit
        </button>
      </div>
    </form>
  </div>
</template>

<script>
import { required, email, minLength,alpha } from '@vuelidate/validators'

export default {
  name: "LoginFormComponent",

  data() {
    return {
      form: {
        firstName:"",
        lastName:"",
        username: "",
        email: ""
      }
    };
  },

  validations: {
    form: {
      firstName: { required, min: minLength(3), alpha},
      lastName: {required, min: minLength(3),alpha},
      username: { required, min: minLength(10) },
      email: { required, email }
    }
  },

  methods: {
    submit() {
      this.$v.form.$touch();
      if(this.$v.form.$error) return
      // to form submit after this
      alert('Form submitted')
    }
  }
};
</script>
