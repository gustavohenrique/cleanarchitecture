<template>
  <q-btn-dropdown
    flat
    color="primary"
    :label="$t('common.actions')"
    icon="settings"
  >
    <q-list>
      <q-item
        v-for="action in actions"
        :key="JSON.stringify(action)"
        clickable
        :to="{ name: action.routeName, params: { id } }"
      >
        <q-item-section>
          <q-item-label>
            {{ action.name }}
          </q-item-label>
        </q-item-section>
      </q-item>
    </q-list>
  </q-btn-dropdown>
</template>

<script>
export default {
  props: {
    id: {
      type: String,
      required: true
    },
    actions: {
      type: Array,
      default: () => []
    }
  }
}
</script>
