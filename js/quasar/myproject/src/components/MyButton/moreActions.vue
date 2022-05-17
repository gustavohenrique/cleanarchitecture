<template>
  <q-btn-dropdown
    data-qa="actions_button"
    color="primary"
    dense
    dark
    flat
    :size="size"
    :disable="disable"
    :dropdown-icon="icon"
    :menu-anchor="menuAnchor"
    :menu-self="menuSelf"
  >
    <q-list separator bordered dense>
      <slot></slot>
    </q-list>
  </q-btn-dropdown>
</template>

<script>
export default {
  props: {
    disable: {
      type: Boolean
    },
    size: {
      type: String,
      default: 'md'
    },
    dropdownIcon: {
      type: String,
      default: 'menu'
    },
    menuAnchor: {
      type: String,
      default: 'bottom left'
    },
    menuSelf: {
      type: String,
      default: 'top start'
    }
  },
  computed: {
    icon () {
      return this.dropdownIcon || 'menu'
    }
  }
}
</script>
