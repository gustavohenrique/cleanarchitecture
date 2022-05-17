import { defineStore } from 'pinia'

export const useLocaleStore = defineStore('locale', {
  state: () => {
    return {
      lang: ''
    }
  },

  actions: {
    setLocale (lang) {
      if (lang === this.lang) {
        return
      }
      this.lang = lang
    }
  }
})
