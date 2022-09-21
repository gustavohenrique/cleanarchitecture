<template>
  <q-page class="flex flex-center my-page">
    <q-card
      flat
      square
      class="my-card"
    >
      <q-card-section>
        <div class="text-h5 text-center text-black-panther text-uppercase">
          {{ $t('auth.authentication') }}
        </div>
      </q-card-section>
      <q-separator />
      <q-card-section>
        <q-carousel
          v-model="form.step"
          transition-prev="slide-right"
          transition-next="slide-left"
          animated
          height="100%"
          :padding="false"
          class="q-pa-none"
        >
          <q-carousel-slide name="check_mail">
            <my-auth-form-mail
              @next="checkEmailIsRegistered"
              :loading="form.loading"
            />
          </q-carousel-slide>
          <q-carousel-slide name="sign_up">
            <my-auth-form-create-password
              @next="signUp"
              @cancel="goToCheckMailStep"
            />
          </q-carousel-slide>
          <q-carousel-slide name="sign_in_first_time">
            <my-auth-form-password
              congratulations
              @next="signIn"
              @cancel="goToCheckMailStep"
            />
          </q-carousel-slide>
          <q-carousel-slide name="sign_in">
            <my-auth-form-password
              @next="signIn"
              @cancel="goToCheckMailStep"
            />
          </q-carousel-slide>
        </q-carousel>
      </q-card-section>
      <q-separator class="q-my-md" />

      <div class="q-mt-md text-center">
        <my-flag-list />
      </div>
    </q-card>
  </q-page>
</template>

<script>
import { mapState, mapActions } from 'pinia'
import { useUserStore } from 'stores/user'

export default {
  data () {
    return {
      key: 1
    }
  },
  computed: {
    ...mapState(useUserStore, ['form'])
  },
  methods: {
    ...mapActions(useUserStore, [
      'goToCheckMailStep',
      'goToSignInStep',
      'goToSignUpStep',
      'checkEmailIsRegistered',
      'register',
      'login'
    ]),
    async signUp () {
      try {
        this.register()
        this.$router.push({ name: 'home' })
      } catch (err) {
        console.log('>>>', err)
      }
    },
    async signIn () {
      try {
        this.login()
        this.$router.push({ name: 'dashboard' })
      } catch (err) {
        console.log('>>>', err)
      }
    }
  }
}
</script>

<style scoped>
body.screen--xs .my-card {
  width: 90vw;
}
body.screen--sm .my-card {
  width: 60vw;
}
.my-card {
  width: 500px;
  height: 480px;
  overflow-y: hidden;
}
.my-page {
  background-position-x: right;
  background: url(~assets/background.jpg);
}
</style>
