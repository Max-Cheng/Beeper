<template>
  <div class="register">
    <b-row class="mt-5">
      <b-col
        md="8"
        offset-md="2"
        lg="6"
        offset-lg="3"
      >
        <b-card title="Register">
          <b-form>
            <b-form-group label="Username">
              <b-form-input
                v-model="$v.user.username.$model"
                type="text"
                required
                placeholder="Enter Username"
                :state="validateState('username')"
              ></b-form-input>
              <b-form-invalid-feedback :state="validateState('username')">
                Username Can't be empty
              </b-form-invalid-feedback>
            </b-form-group>
            <b-form-group label="Password">
              <b-form-input
                v-model="$v.user.password.$model"
                type="password"
                required
                placeholder="Enter Password"
                :state="validateState('password')"
              ></b-form-input>
              <b-form-invalid-feedback :state="validateState('password')">
                Password between 6-16
              </b-form-invalid-feedback>
            </b-form-group>
            <b-form-group>
              <b-button
                @click="register"
                variant="outline-primary"
                block
              >Register</b-button>
            </b-form-group>
          </b-form>
        </b-card>
      </b-col>
    </b-row>
  </div>
</template>
<script>
import { required, minLength, maxLength } from 'vuelidate/lib/validators';

export default {
  data() {
    return {
      user: {
        username: '',
        password: '',
      },
    };
  },
  validations: {
    user: {
      username: {
        required,
        minLength: minLength(4),
      },
      password: {
        required,
        minLength: minLength(6),
        maxLength: maxLength(16),
      },
    },
  },
  methods: {
    validateState(name) {
      const { $dirty, $error } = this.$v.user[name];
      return $dirty ? !$error : null;
    },
    register() {
      this.$v.user.$touch();
      if (this.$v.user.$anyError) {
        return;
      }
      const api = 'http://localhost:1324/api/auth/register';
      this.axios.post(api, { ...this.user }).then((res) => {
        console.log(res.data);
      }).catch((err) => {
        console.log('err:', err.response.data.msg);
      });
      console.log('register');
    },
  },
};
</script>
<style lang="scss">
</style>
