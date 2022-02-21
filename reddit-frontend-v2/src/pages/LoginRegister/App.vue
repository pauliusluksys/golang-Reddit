<template>
<div>
<section class="vh-100">
<div class="container-fluid h-custom">
  <div class="row d-flex justify-content-center align-items-center h-100">
    <div class="col-md-9 col-lg-6 col-xl-5">
      <img src="https://mdbcdn.b-cdn.net/img/Photos/new-templates/bootstrap-login-form/draw2.webp" class="img-fluid"
           alt="Sample image">
    </div>
    <div class="col-md-8 col-lg-6 col-xl-4 offset-xl-1">
      <form>
        <div class="d-flex flex-row align-items-center justify-content-center justify-content-lg-start">
          <p class="lead fw-normal mb-0 me-3">Sign in with</p>
          <button type="button" class="btn btn-primary btn-floating mx-1">
            <i class="fab fa-facebook-f"></i>
          </button>

          <button type="button" class="btn btn-primary btn-floating mx-1">
            <i class="fab fa-twitter"></i>
          </button>

          <button type="button" class="btn btn-primary btn-floating mx-1">
            <i class="fab fa-linkedin-in"></i>
          </button>
        </div>

        <div class="divider d-flex align-items-center justify-content-center my-4">
          <p class="text-center fw-bold mx-3 mb-0">Or</p>
        </div>
        <div style="width: 70%; margin: 0 auto;" class="" >
          <div class="">
        <!-- Email input -->
        <div class="form-outline mb-4">
          <input type="email" id="form3Example3" class="form-control form-control-lg"
                 placeholder="Enter a valid email address" />
          <label class="form-label d-flex justify-content-center" for="form3Example3">Email address</label>
        </div>

        <!-- Password input -->
        <div class="form-outline mb-3">
          <input type="password" id="form3Example4" class="form-control form-control-lg"
                 placeholder="Enter password" />
          <label class="form-label d-flex justify-content-center" for="form3Example4">Password</label>
        </div>
        </div>

        <div class="d-flex justify-content-between align-items-center">
          <!-- Checkbox -->
          <div class="form-check mb-0">
            <input class="form-check-input me-2" type="checkbox" value="" id="form2Example3" />
            <label class="form-check-label" for="form2Example3">
              Remember me
            </label>
          </div>
          <a href="#!" class="text-body">Forgot password?</a>
        </div>
        </div>
        <div class="text-center text-lg-start mt-4 pt-2 mb-2 d-flex justify-content-center">
          <button type="button" class="btn btn-primary btn-lg"
                  style="padding-left: 2.5rem; padding-right: 2.5rem;">Login</button>
        </div>


      </form>
      <div class="d-flex justify-content-center">
      <button data-bs-toggle="modal" data-bs-target="#registerDrop" class="btn btn-link p-0">
        <p class="small fw-bold d-flex aligns-items-center text-end">Don't have an account?</p>
      </button>
      </div>
    </div>
  </div>
</div>
<div class="d-flex flex-column flex-md-row text-center text-md-start justify-content-between py-4 px-4 px-xl-5 bg-primary">
  <!-- Copyright -->
  <div class="text-white mb-3 mb-md-0">
    Copyright © 2020. All rights reserved.
  </div>
  <!-- Copyright -->

  <!-- Right -->
  <div>
    <a href="#!" class="text-white me-4">
      <i class="fab fa-facebook-f"></i>
    </a>
    <a href="#!" class="text-white me-4">
      <i class="fab fa-twitter"></i>
    </a>
    <a href="#!" class="text-white me-4">
      <i class="fab fa-google"></i>
    </a>
    <a href="#!" class="text-white">
      <i class="fab fa-linkedin-in"></i>
    </a>
  </div>
  <!-- Right -->
</div>
</section>



  <div class="modal fade" id="registerDrop" data-bs-backdrop="static" data-bs-keyboard="false" tabindex="-1" aria-labelledby="staticBackdropLabel" aria-hidden="true">
    <div class="modal-dialog">
      <div class="modal-content">
        <div class="p-2 d-flex justify-content-between">
          <h5 class="modal-title" id="staticBackdropLabel">Registration Form</h5>
          <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
        </div>
        <div class="position-relative flex-1">
          <section class="h-100 h-custom" style="background-color: #0d6efd;">
            <div class="container py-4 h-100">
              <div class="row d-flex justify-content-center align-items-center h-100 px-5">
                <div class="">
                  <div class="card rounded-3">
                    <div class="card-body p-2 p-md-2">

                      <form class="px-md-2">

                        <div class="form-outline mb-2">
                          <label class="form-label" for="email">Email</label>
                          <input type="text" id="email" class="form-control" v-model="state.email" />
                          <span v-if="v$.email.$error">
                              {{ v$.email.$errors[0].$message }}
                          </span>

                        </div>
                        <div class="form-outline mb-2">
                          <label class="form-label" for="password" >Password</label>
                          <input type="text" id="password" class="form-control" v-model="state.password.password" />
                          <span v-if="v$.password.password.$error">
                              {{ v$.password.password.$errors[0].$message }}
                          </span>
                        </div>
                        <div class="form-outline mb-2">
                          <label class="form-label" for="confirm-password">Confirm Password</label>
                          <input type="text" id="confirm-password" class="form-control" v-model="state.password.confirm"/>
                          <span v-if="v$.password.confirm.$error">
                             {{ v$.password.confirm.$errors[0].$message }}
                          </span>
                        </div>


                        <div class="row mb-2 pb-2 pb-md-0 mb-md-5">
                          <div class="col-md-6">


                          </div>
                        </div>
                        <div class="d-flex justify-content-center">
                        <button @click.prevent="submitForm" type="submit" class="btn btn-primary btn-lg mb-1">Submit</button>
                      </div>
                      </form>

                    </div>
                  </div>
                </div>
              </div>
            </div>
          </section>
        </div>
      </div>
    </div>
  </div>


</div>

</template>
<script>
import "bootstrap/dist/css/bootstrap.min.css"
import "bootstrap"
import useValidate from "@vuelidate/core";
import { required,email,minLength,sameAs } from "@vuelidate/validators";
import { reactive, computed } from "vue";
export default {
  setup() {
    // const mustBeLearnVue = (value) => value.includes("learnvue");
    const state = reactive({
      email: '',
      password: {
        password: '',
        confirm: '',
      },
    })
    const rules = computed(() => {
      return {
        email: {
          required,
          email,
        },
        password: {
          password: {required, minLength: minLength(6)},
          confirm: {required, sameAs: sameAs(state.password.password)},
        },
      };
    })
    const v$ = useValidate(rules, state)

    return {state, v$}
  },
  methods: {
    submitForm() {
      // this.v$.$validate()
      // if (!this.v$.$error) { // if ANY fail validation
      //   alert('Form successfully submitted.')
      // } else {
      //   alert('Form failed validation')
      // }
      this.v$.$validate()
      if(!this.v$.$error) {
        alert('Form succesfully submitted.')

      } else {
        alert('Form failed validation')

      }

    }
  }
};


</script>