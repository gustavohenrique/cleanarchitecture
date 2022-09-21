<template>
  <my-auth-form-base
    @save="$emit('next')"
  >
    <div class="text-center text-dolphin q-pb-md">
      <q-icon
        name="portrait"
        size="96px"
      />
    </div>
    <div>
      <q-input
        :model-value="user.email"
        @update:model-value="val => setUser({ email: val.trim() })"
        label="Email"
        data-qa="auth_email"
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
import { mapState, mapActions } from 'pinia'
import { useUserStore } from 'stores/user'

export default {
  props: {
    loading: {
      type: Boolean
    }
  },
  computed: {
    ...mapState(useUserStore, ['user'])
  },
  methods: {
    ...mapActions(useUserStore, ['setUser'])
  }
}
</script>
