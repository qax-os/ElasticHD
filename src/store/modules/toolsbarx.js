const toolsbarx = {
  state: {
    'toggleTitle': '',
    'btnGroup': [
      {fontAwesome: 'fa fa-info-circle', value: 'tools-info', toolsTip: 'cluster info【集群健康】'}
    ]
  },
  mutations: {
    SET_TOGGLE_TITLE: (state, title) => {
      state.toggleTitle = title
    }
  },
  actions: {
    ToggleTitle ({ commit }, index) {
      switch (index) {
        case 0:
          commit('SET_TOGGLE_TITLE', "<i class='fa fa-bar-chart-o'></i> Overview【集群概要】")
          break
        case 1:
          commit('SET_TOGGLE_TITLE', "<i class='fa fa-search'></i> Search【数据搜索】")
          break
        case 2:
          commit('SET_TOGGLE_TITLE', "<i class='fa fa-list-alt'></i> Indices【索引列表】")
          break
        case 3:
          commit('SET_TOGGLE_TITLE', "<i class='fa fa-table'></i> Mappings【数据模板】")
          break
        case 4:
          commit('SET_TOGGLE_TITLE', "<i class='fa fa-gear'></i> Tools Box【工具箱】")
          break
        case 5:
          commit('SET_TOGGLE_TITLE', "<i class='fa fa-comments'></i> Syntactic Sugar【语法糖】")
          break
        case 6:
          commit('SET_TOGGLE_TITLE', "<i class='fa fa-comments'></i> Monitor Eye【ES服务监控】")
          break
        case 7:
          commit('SET_TOGGLE_TITLE', "<i class='fa fa-comments'></i> Query UI【通用查询】")
          break
      }
    }
  }
}
export default toolsbarx
