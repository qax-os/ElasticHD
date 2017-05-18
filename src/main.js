// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import store from './store'

import 'bootstrap/dist/css/bootstrap.css'
import 'font-awesome/css/font-awesome.min.css'
import 'element-ui/lib/theme-default/index.css'

import 'jquery-jsonview/dist/jquery.jsonview.css'
import 'jquery-jsonview'

Vue.config.productionTip = false

// Tell Vue to use the plugin
import VueCookie from 'vue-cookie'
Vue.use(VueCookie)

// 使用VueResource插件
import VueResource from 'vue-resource'
Vue.use(VueResource)

// 使用tool-tip 插件
import VTooltip from 'v-tooltip'
Vue.use(VTooltip)

// 使用element ui
import Element from 'element-ui'
Vue.use(Element)

// use json tree view
import TreeView from 'vue-json-tree-view'
Vue.use(TreeView)

// import VueTreeView from 'vue-tree-view'
// Vue.use(VueTreeView)

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store: store,
  template: '<App/>',
  components: { App }
})
