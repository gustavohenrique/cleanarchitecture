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
          v-model="step"
          transition-prev="slide-right"
          transition-next="slide-left"
          animated
          height="100%"
          :padding="false"
          class="q-pa-none"
        >
          <q-carousel-slide name="check-email">
            <my-auth-form-mail
              @next="checkEmail"
              :loading="loading"
            />
          </q-carousel-slide>
          <q-carousel-slide name="sign-up">
            <my-auth-form-create-password
              @next="signUp"
              @cancel="goToStep('check-email')"
            />
          </q-carousel-slide>
          <q-carousel-slide name="sign-in-first-time">
            <my-auth-form-password
              congratulations
              @next="signIn"
              @cancel="goToStep('check-email')"
            />
          </q-carousel-slide>
          <q-carousel-slide name="sign-in">
            <my-auth-form-password
              @next="signIn"
              @cancel="goToStep('check-email')"
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
import { useAuthUserStore } from 'stores/user'

export default {
  setup () {
    return {
      authUserStore: useAuthUserStore()
    }
  },
  data () {
    return {
      loading: false,
      key: 1,
      step: 'check-email'
    }
  },
  methods: {
    goToStep (step) {
      this.step = step
    },
    async checkEmail (email) {
      this.loading = true
      const found = await this.$accountController.findUserByEmail(email)
      const nextStep = found.salt ? 'sign-in' : 'sign-up'
      this.authUserStore.setUser(found || { email })
      this.loading = false
      this.goToStep(nextStep)
    },
    async signUp (rawPassword) {
      try {
        const { user } = this.authUserStore
        const authenticated = await this.$accountController.signUp({ ...user, rawPassword })
        this.authUserStore.setUser(authenticated)
        this.$router.push({ name: 'home' })
      } catch (err) {
        console.log('>>>', err)
      }
    },
    async signIn (rawPassword) {
      try {
        const { user } = this.authUserStore
        const authenticated = await this.$accountController.signIn({ ...user, rawPassword })
        this.authUserStore.setUser(authenticated)
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
