<template>
  <q-btn
    icon="content_copy"
    flat
    dense
    :size="size"
    :color="color"
    @click="copy"
  />
</template>

<script>
import { copyToClipboard, useQuasar } from 'quasar'

export default {
  props: {
    color: {
      type: String,
      default: 'primary'
    },
    size: {
      type: String,
      default: 'sm'
    },
    text: {
      type: String,
      required: true
    }
  },
  setup () {
    const $q = useQuasar()
    return {
      showNotify (opts) {
        $q.notify(opts)
      }
    }
  },
  methods: {
    copy () {
      const text = this.text || ''
      const { $t } = this
      copyToClipboard(text).then(() => {
        this.showNotify({
          message: $t('common.notify.copied') + '!',
          color: 'positive',
          timeout: 900
        })
      }).catch(() => {
        this.showNotify({
          message: $t('common.notify.failed') + '!',
          color: 'negative',
          timeout: 900
        })
      })
    }
  }
}
</script>
