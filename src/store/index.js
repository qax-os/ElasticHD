import Vue from 'vue'
import Vuex from 'vuex'
import mutations from './mutations'
import actions from './actions'
import getters from './getters'
import headerx from './modules/headerx'
import sidebarx from './modules/sidebarx'
import toolsbarx from './modules/toolsbarx'
import overviewx from './modules/overviewx'
import searchx from './modules/searchx'
import mappingx from './modules/mappingx'
import toolsx from './modules/toolsx'

Vue.use(Vuex)

const state = {
  showIndex: 0
}

export default new Vuex.Store({
  state,
  mutations,
  actions,
  getters,
  modules: {
    headerx,
    sidebarx,
    toolsbarx,
    overviewx,
    searchx,
    mappingx,
    toolsx
  }
})
