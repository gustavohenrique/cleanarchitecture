<template>
  <q-btn
    :color="color"
    :data-qa="qa"
    :type="to ? 'a' : type"
    :to="to"
    :label="label"
    :disable="disable"
    :icon="icon"
    :icon-right="iconRight"
    :unelevated="true"
    class="no-border-radius"
  />
</template>

<script>
export default {
  props: {
    qa: {
      type: String
    },
    icon: {
      type: String
    },
    iconRight: {
      type: String
    },
    color: {
      type: String,
      default: 'primary'
    },
    toRoute: {
      type: String
    },
    toParams: {
      type: Object
    },
    type: {
      type: String,
      default: 'submit'
    },
    label: {
      type: String,
      required: true
    },
    disable: {
      type: Boolean
    }
  },
  computed: {
    to () {
      const { toRoute, toParams } = this
      if (toRoute) {
        return {
          name: toRoute,
          params: toParams
        }
      }
      return null
    }
  }
}
</script>
