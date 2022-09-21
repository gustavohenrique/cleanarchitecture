<template>
  <div class="row wrap q-pt-sm justify-center q-gutter-md">
    <div
      class="btn-en row no-wrap items-center cursor-pointer shadow-1 q-pa-sm"
      style="width: 100px"
      @click="setLocale('en')"
    >
      <my-flag-en
        size="22px"
        :style="lang === 'en' ? '' : 'filter:grayscale(1)'"
      />
      <div class="q-pl-sm" style="font-size:12px">
        {{ $t('common.locale.en') }}
      </div>
    </div>
    <div
      class="row no-wrap items-center cursor-pointer shadow-1 q-pa-sm"
      style="width: 100px"
      @click="setLocale('pt')"
    >
      <my-flag-pt
        size="22px"
        :style="lang === 'pt' ? '' : 'filter:grayscale(1)'"
      />
      <div class="q-pl-sm" style="font-size:12px">
        {{ $t('common.locale.pt') }}
      </div>
    </div>
  </div>
</template>

<script>
import { useI18nStore } from 'stores/i18n'

export default {
  setup () {
    return {
      store: useI18nStore()
    }
  },
  computed: {
    lang () {
      return this.store.lang
    }
  },
  methods: {
    setLocale (lang) {
      this.store.setLocale(lang)
    }
  }
}
</script>
