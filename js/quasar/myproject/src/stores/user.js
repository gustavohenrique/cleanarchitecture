import { defineStore } from 'pinia'
import UserAdapter from 'src/application/adapters/userAdapter'

export const useAuthUserStore = defineStore('user', {
  state: () => {
    return {
      user: new UserAdapter().toJSON()
    }
  },

  actions: {
    setUser (payload) {
      this.user = payload
    }
  }
})
