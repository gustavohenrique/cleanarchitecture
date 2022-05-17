<template>
  <my-auth-form-base
    @save="next"
  >
    <div class="text-center text-dolphin q-pb-md">
      <q-icon
        name="portrait"
        size="96px"
      />
    </div>
    <div>
      <q-input
        label="Email"
        data-qa="auth_email"
        v-model="email"
        square
        dense
        :rules="[
        val => !!val || $t('common.validation.required'),
        val => val.trim().length <= 30 || $t('common.validation.max', { max: 100 })
        ]"
        maxlength="100"
        filled
        bg-color="blue-grey-1"
      />
      <div class="q-mt-sm">
        <my-button-primary
          :label="$t('auth.next')"
          :loading="loading"
          icon-right="arrow_right"
          align="between"
          class="full-width"
        />
      </div>
    </div>
  </my-auth-form-base>
</template>

<script>
export default {
  props: {
    loading: {
      type: Boolean
    }
  },
  data () {
    return {
      email: ''
    }
  },
  methods: {
    next () {
      this.$emit('next', this.email)
    }
  }
}
</script>
