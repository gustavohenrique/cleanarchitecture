<template>
  <div>
    <my-button-outline
      @click="search"
      :label="$t('common.component.button.loadMore')"
      :loading="loading"
    />
    <span v-if="noMoreItems" class="q-ml-md text-negative">
      <translate text="common.notify.noMoreResults" />
    </span>
  </div>
</template>

<script>
export default {
  props: {
    loading: {
      type: Boolean
    },
    noMoreItems: {
      type: Boolean
    }
  },
  emits: ['search'],
  data () {
    return {
      page: 1
    }
  },
  methods: {
    search () {
      this.page++
      this.$emit('search', { page: this.page })
    }
  }
}
</script>
