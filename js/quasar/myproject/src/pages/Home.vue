<template>
  <q-page class="q-pa-md">
    <div class="text-h2">
      Olá! {{ user.email }}
    </div>
  </q-page>
</template>

<script>
import { mapState } from 'pinia'
import { useUserStore } from 'stores/user'

export default {
  computed: {
    ...mapState(useUserStore, ['user'])
  }
}
</script>
