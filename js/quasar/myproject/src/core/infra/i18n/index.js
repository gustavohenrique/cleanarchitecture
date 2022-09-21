import frenchkiss from 'frenchkiss'
import en from './en'
import pt from './pt'

const getBrowserLanguage = () => {
  const userLanguage = navigator.language || navigator.userLanguage
  const lang = userLanguage.startsWith('pt') ? 'pt' : 'en'
  return lang
}

const localStorageKey = 'i18n'

export class Internationalization {
  constructor () {
    const lang = localStorage.getItem(localStorageKey) || getBrowserLanguage()
    frenchkiss.locale(lang)
    frenchkiss.fallback('en')
    this.locale = frenchkiss.locale
    this.language = getBrowserLanguage
    localStorage.setItem(localStorageKey, lang)
  }

  setLocale (lang) {
    localStorage.setItem(localStorageKey, lang)
    this.locale(lang)
  }

  getLocale () {
    return this.locale()
  }

  setMessages (lang, messages) {
    frenchkiss.set(lang, messages)
  }

  pushMessages (lang, messages) {
    frenchkiss.extend(lang, messages)
  }

  translate (key, params, lang) {
    const language = lang || localStorage.getItem(localStorageKey)
    return frenchkiss.t(key, params, language)
  }
}

const i18n = new Internationalization()
i18n.setMessages('en', en)
i18n.setMessages('pt', pt)

export function $t (key, params, lang) {
  return i18n.translate(key, params, lang)
}

export default i18n
