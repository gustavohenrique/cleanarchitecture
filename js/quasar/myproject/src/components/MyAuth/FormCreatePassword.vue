<template>
  <my-auth-form-base
    @save="$emit('next')"
  >
    <div class="text-left text-dolphin q-pb-md row justify-start q-gutter-x-md items-center">
      <my-avatar
        :src="user.picture"
        icon="portrait"
        size="62"
      />
      <div>
        <div class="text-h6 q-pt-none q-mt-none">
          {{ user.fullName }}
        </div>
        <div class="text-body2">
          {{ user.email }}
        </div>
      </div>
    </div>
    <div class="q-pt-md text-positive">
      {{ $t('auth.firstTime') }}
    </div>
    <div class="q-pt-md">
      <q-input
        :model-value="user.password"
        @update:model-value="val => setUser({ password: val.trim() })"
        :label="$t('auth.password')"
        :type="isHiddenPassword ? 'password' : 'text'"
        data-qa="auth_password"
        :hint="$t('auth.hintPassword')"
        square
        dense
        :rules="[
        val => !!val || $t('common.validation.required'),
        val => val.trim().length >= 6 || $t('common.validation.min', { min: 6 })
        ]"
        maxlength="30"
        filled
        bg-color="blue-grey-1"
        autocomplete
      >
        <template v-slot:append>
          <q-icon
            :name="isHiddenPassword ? 'visibility_off' : 'visibility'"
            class="cursor-pointer"
            @click="isHiddenPassword = !isHiddenPassword"
          />
        </template>
      </q-input>
      <div class="q-mt-sm">
        <my-button-primary
          :label="$t('auth.signInWithNewPassword')"
          :loading="false"
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
  data () {
    return {
      isHiddenPassword: true
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
