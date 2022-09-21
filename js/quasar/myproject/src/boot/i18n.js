import { boot } from 'quasar/wrappers'
import i18n, { $t } from '../core/infra/i18n'
import { useI18nStore } from 'stores/i18n'

/*
 * Usage:
 * this.$i18n.setMessages({ en: { hello: 'Hello {name}' }, pt: { hello: 'Olá {name} ' } })
 * this.$t('hello', { name: 'Gustavo' })       // Hello Gustavo
 * this.$t('hello', { name: 'Gustavo' }, 'pt') // Olá Gustavo
 */
export default boot(async ({ app }) => {
  // Using store, all messages are updated when locale changes
  const store = useI18nStore()
  store.setLocale(i18n.getLocale())
  app.config.globalProperties.$i18n = i18n
  app.config.globalProperties.$t = (key, params) => {
    return $t(key, params, store.lang)
  }
})
