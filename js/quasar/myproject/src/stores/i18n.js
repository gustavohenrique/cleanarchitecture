import { defineStore } from 'pinia'

export const useI18nStore = defineStore('i18n', {
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
