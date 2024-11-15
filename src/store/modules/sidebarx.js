const sidebarx = {
  state: {
    msg: 'ebcloud',
    activeIndex: 0,
    toggleShow: true,
    items: [
      { fontAwesome: 'fa fa-dashboard', value: '仪表盘' },
      { fontAwesome: 'fa fa-search', value: '查询' },
      { fontAwesome: 'fa fa-list', value: '索引' },
      { fontAwesome: 'fa fa-table', value: '映射' },
      {fontAwesome: 'fa fa-gears', value: '工具'},
      {fontAwesome: 'fa fa-comments-o', value: '帮助'}
    ]
  },
  mutations: {
    SET_MSG: (state, msg) => {
      state.msg = msg
    },
    SET_ACTIVE_INDEX: (state, index) => {
      state.activeIndex = index
    },
    SET_TOGGLE_SHOW: (state, show) => {
      state.toggleShow = show
    }
  },
  actions: {
    ToggleShow ({ commit }, show) {
      commit('SET_TOGGLE_SHOW', !show)
    },
    SetActiveIndex ({ commit }, index) {
      commit('SET_ACTIVE_INDEX', index)
    }
  }
}
export default sidebarx
