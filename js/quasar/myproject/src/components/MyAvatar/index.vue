<template>
  <q-avatar
    :square="square"
    :size="size"
  >
    <img v-if="src" :src="src" />
    <q-icon v-else :name="icon" :size="size" />
  </q-avatar>
</template>

<script>
export default {
  props: {
    src: {
      type: String,
      default: '',
      required: true
    },
    icon: {
      type: String,
      default: 'person'
    },
    size: {
      type: String,
      default: 'lg'
    },
    square: {
      type: Boolean
    }
  }
}
</script>
