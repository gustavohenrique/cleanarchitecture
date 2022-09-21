/**
 * @jest-environment jsdom
 */

import chai from 'chai'
import { shallowMount } from '@vue/test-utils'
import { setActivePinia, createPinia } from 'pinia'
// import { createTestingPinia } from '@pinia/testing'

import MyFlagList from '../list.vue'
import MyFlagEn from '../en.vue'
import MyFlagPt from '../pt.vue'
import { useI18nStore } from '../../../stores/i18n'

const { expect } = chai

describe('MyFlagList', () => {
  let store = null
  let wrapper = null

  beforeEach(() => {
    const pinia = setActivePinia(createPinia())
    store = useI18nStore()
    wrapper = shallowMount(MyFlagList, {
      global: {
        plugins: [pinia],
        components: {
          'my-flag-en': MyFlagEn,
          'my-flag-pt': MyFlagPt
        },
        mocks: {
          $t: (msg) => msg
        }
      }
    })
  })

  describe('#setLocale', () => {
    it('Should change locale to english', async () => {
      const button = wrapper.find('.btn-en')
      await button.trigger('click')
      expect(store.lang).to.equal('en')
    })
  })
})
