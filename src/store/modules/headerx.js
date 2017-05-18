import VueCookie from 'vue-cookie'
import Vue from 'vue'
import VueResource from 'vue-resource'
Vue.use(VueResource)
const headerx = {
  state: {
    serverHost: 'http://127.0.0.1:9200',
    status: '未连接',
    githubUrl: 'https://github.com/farmerx',
    statusStyleObject: {
      color: '#001422',
      background: '#ffffff',
      textAlign: 'center',
      borderRadius: '4px',
      fontWeight: '800'
    },
    health: {
      'active_primary_shards': 0,
      'active_shards': 0,
      'active_shards_percent_as_number': 0,
      'cluster_name': 'docker-es-master',
      'delayed_unassigned_shards': 0,
      'initializing_shards': 0,
      'number_of_data_nodes': 0,
      'number_of_in_flight_fetch': 0,
      'number_of_nodes': 0,
      'number_of_pending_tasks': 0,
      'relocating_shards': 0,
      'status': '未连接',
      'task_max_waiting_in_queue_millis': 0,
      'timed_out': false,
      'unassigned_shards': 0
    },
    indexList: [
    ]
  },
  mutations: {
    SET_SERVER_HOST: (state, serverhost) => {
      state.serverHost = serverhost
    },
    SET_STATUS: (state, status) => {
      state.status = status
    },
    SET_GITHUB_URL: (state, url) => {
      state.githubUrl = url
    },
    SET_STATUS_STYLE: (state, style) => {
      state.statusStyleObject = style
    },
    SET_HEALTH: (state, health) => {
      state.health = health
    },
    SET_INDEX_LIST: (state, list) => {
      state.indexList = list
    }
  },
  actions: {
    CookieSetServerHost ({ commit }, serverhost) {
      serverhost = serverhost.trim()
      if (serverhost === '') {
        serverhost = 'http://127.0.0.1:9200'
      }
      if (serverhost.indexOf('http://') === -1) {
        serverhost = 'http://' + serverhost
      }
      commit('SET_SERVER_HOST', serverhost)
      VueCookie.set('elasticHDServerHost', serverhost, { expires: '1D' })
      window.location.reload()
    },
    CookieGetServerHost ({ commit }) {
      if (VueCookie.get('elasticHDServerHost') !== null) {
        commit('SET_SERVER_HOST', VueCookie.get('elasticHDServerHost'))
      }
    },
    HttpPost ({ commit }, body) {
      Vue.http.post(body.url, {'host': body.host})
      .then(
        response => {
          if (response.body.result === 0) {
            switch (response.body.data.status) {
              case 'green':
                commit('SET_STATUS', 'Green')
                commit('SET_STATUS_STYLE', { color: '#001422', background: '#28C768', textAlign: 'center', borderRadius: '4px', fontWeight: '800' })
                break
              case 'yellow':
                commit('SET_STATUS', 'Yellow')
                commit('SET_STATUS_STYLE', { color: '#001422', background: '#28C768', textAlign: 'center', borderRadius: '4px', fontWeight: '800' })
                break
              case 'red':
                commit('SET_STATUS', 'Red')
                commit('SET_STATUS_STYLE', { color: '#001422', background: '#28C768', textAlign: 'center', borderRadius: '4px', fontWeight: '800' })
                break
            }
            commit('SET_HEALTH', response.body.data)
          }
        },
        error => {
          console.log(error)
        }
      )
    },
    GetIndexList ({ commit }, body) {
      Vue.http.post(body.url, {'host': body.host})
      .then(
        response => {
          if (response.body.result === 0) {
            commit('SET_INDEX_LIST', response.body.data)
          }
        },
        error => {
          console.log(error)
        }
      )
    }
  }
}
export default headerx

