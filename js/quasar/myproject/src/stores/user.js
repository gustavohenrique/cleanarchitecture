import { defineStore } from 'pinia'

const STEPS = {
  CHECK_MAIL: 'check_mail',
  SIGN_IN: 'sign_in',
  SIGN_IN_FIRST_TIME: 'sign_in_first_time',
  SIGN_UP: 'sign_up'
}

export const useUserStore = defineStore('user', {
  state: () => {
    return {
      _user: {},
      form: {
        loading: false,
        step: STEPS.CHECK_MAIL
      }
    }
  },

  getters: {
    user () {
      return this._user
    },

    availableMenus () {
      return [
        {
          routeName: 'account',
          icon: 'manage_accounts'
        },
        {
          routeName: 'dns',
          icon: 'dns'
        },
        {
          routeName: 'insights',
          icon: 'analytics'
        },
        {
          routeName: 'help',
          icon: 'live_help'
        }
      ]
    }
  },

  actions: {
    async setUser (payload) {
      this._user = { ...this._user, ...payload }
    },
    goToCheckMailStep () {
      this.form.step = STEPS.CHECK_MAIL
    },
    goToSignInStep () {
      this.form.step = STEPS.SIGN_IN
    },
    goToSignInFirstTimeStep () {
      this.form.step = STEPS.SIGN_IN_FIRST_TIME
    },
    goToSignUpStep () {
      this.form.step = STEPS.SIGN_UP
    },
    async loadLocalUser () {
      const found = await this.$controllers.Account.getLocalUser()
      this.setUser(found)
    },
    async checkEmailIsRegistered () {
      this.form.loading = true
      const { email } = this.user
      const found = await this.$controllers.Account.findUserByEmail(email)
      this.setUser(found || { email })
      this.form.loading = false
      if (found.salt) {
        this.goToSignInFirsTimeStep()
        return
      }
      this.goToSignUpStep()
    },
    async register () {
      const { user } = this
      const authenticated = await this.$controllers.Account.signUp(user)
      this.setUser(authenticated)
    },
    async login () {
      const { user } = this
      const authenticated = await this.$controllers.Account.signIn(user)
      this.setUser(authenticated)
    }
  }
})
