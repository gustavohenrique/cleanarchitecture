<template>
  <q-layout
    view="lHh lpr lFf"
    :key="key"
  >
    <q-drawer
      v-model="leftDrawerOpen"
      class="flex column full-height justify-between items-center bg-primary text-white"
      :breakpoint="320"
      :width="70"
    >
      <q-list>
        <q-item>
          <q-item-section>
            <my-avatar
              :src="user ? user.picture : ''"
              size="md"
              class="q-mt-xs"
            />
          </q-item-section>
        </q-item>
      </q-list>
      <q-list class="flex column">
        <q-item
          v-for="i in availableMenus"
          :key="JSON.stringify(i)"
          tag="a"
          clickable
          :to="{ name: i.routeName }"
          :class="isActive(i.routeName) ? 'drawer-menu-active' : ''"
        >
          <q-item-section>
            <q-icon :name="i.icon" size="xs" />
          </q-item-section>
        </q-item>
      </q-list>
      <q-list>
        <q-item>
          <q-item-section>
            <q-btn-dropdown
              dense
              size="md"
              flat
              rounded
              stack
              dropdown-icon="fas fa-globe"
              no-icon-animation
              auto-close
            >
              <q-list bordered separator>
                <q-item-label header>
                  {{ $t('common.language') }}
                </q-item-label>
                <q-item
                  clickable
                  @click="setLocale('en')"
                >
                  <q-item-section avatar>
                    <my-flag-en
                      size="32px"
                      :style="lang === 'en' ? '' : 'filter:grayscale(1)'"
                    />
                  </q-item-section>
                  <q-item-section>
                    <q-item-label>
                      {{ $t('common.locale.en') }}
                    </q-item-label>
                  </q-item-section>
                </q-item>
                <q-item
                  clickable
                  @click="setLocale('pt')"
                >
                  <q-item-section avatar>
                    <my-flag-pt
                      size="32px"
                      :style="lang === 'pt' ? '' : 'filter:grayscale(1)'"
                    />
                  </q-item-section>
                  <q-item-section>
                    <q-item-label>
                      {{ $t('common.locale.pt') }}
                    </q-item-label>
                  </q-item-section>
                </q-item>
              </q-list>
            </q-btn-dropdown>
          </q-item-section>
        </q-item>
      </q-list>
    </q-drawer>
    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script>
import { mapState } from 'pinia'
import { useUserStore } from 'stores/user'

export default {
  data () {
    return {
      key: 0,
      lang: this.$i18n.getLocale(),
      leftDrawerOpen: true
    }
  },
  async created () {
    const store = useUserStore()
    await store.loadLocalUser()
  },
  computed: {
    ...mapState(useUserStore, ['user', 'availableMenus'])
  },
  methods: {
    isActive (routeName) {
      const { $route } = this
      if (!$route || !$route.name) {
        return false
      }
      const a = routeName.split('_')[0]
      const b = $route.name.split('_')[0]
      return a === b
    },
    setLocale (lang) {
      if (lang === this.lang) {
        return
      }
      this.$i18n.setLocale(lang)
      this.lang = lang
      this.key++
    }
  }
}
</script>

<style scoped>
.drawer-menu-active {
  border-left: 2px solid #fff;
  background: #f8f8f8;
  color: var(--q-primary);
}
</style>
