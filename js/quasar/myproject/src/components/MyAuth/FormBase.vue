<template>
  <q-form
    method="post"
    ref="form"
    class="full-width text-black-panther"
    @submit.prevent="submit"
    @reset="reset"
  >
    <slot></slot>
  </q-form>
</template>

<script>
export default {
  methods: {
    submit () {
      const form = this.$refs.form
      form.validate().then(success => {
        if (success) {
          form.resetValidation()
          this.$emit('save')
          return
        }
        const message = this.$t('common.form.validation')
        this.$logger.error(message)
      })
    },
    reset () {
      const form = this.$refs.form
      form.resetValidation()
      this.$emit('cancel')
    }
  }
}
</script>
