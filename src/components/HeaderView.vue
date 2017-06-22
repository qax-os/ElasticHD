<template>
<!-- nav bar -->
    <nav class="navbar navbar-inverse navbar-fixed-top" style="margin-bottom: 0;" role="navigation" id="elastic-hd-header">
        <div class="container-fluid">
            <!--放置logo...-->
            <div class="navbar-header">
                <a class="navbar-brand" id="elastic-logo-color" href="/"><i class="fa fa-dashboard fa-x"> </i> ElasticHD</a>
            </div>
            <!--end-->
            <!-- Collect the nav links, forms, and other content for toggling -->
            <div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
                <form class="navbar-form navbar-left" role="search">
                    <div class="form-group">
                        <div class="input-group">
                            <div class="input-group-addon"><i class="fa fa-sitemap fa-1x"></i></div>
                            <input class="form-control" v-model="serverhost" id="elastic-host">
                        </div>
                         <button @click.prevent="elasticHDConnect" class="btn btn-primary">Connect</button>
                    </div>
                </form>
                <p class="navbar-text">集群健康状态：<span id="elastic-status" :style="statusStyleObject">{{status}}</span></p>
                <ul class="nav navbar-nav pull-right">
                    <li>
                        <a v-bind:href="githubUrl" target="_blank"><i class="fa fa-github"></i> Star On GitHub</a>
                    </li>
                </ul>
            </div>
        </div>
    </nav>
</template>

<script>

import { mapGetters } from 'vuex'
export default {
  name: 'elastic-hd-header',
  data: function () {
    return {
      serverhost: ''
    }
  },
  computed: {
    ...mapGetters([
      'githubUrl',
      'status',
      'statusStyleObject'
    ])
  },
  methods: {
    elasticHDConnect: function () {
      this.$store.dispatch('CookieSetServerHost', this.serverhost)
    }
  },
  created: function () {
    this.$store.dispatch('CookieGetServerHost')
    this.serverhost = this.$store.getters.serverhost
    this.$store.dispatch('HttpPost', {url: '/_cat/health', host: this.serverhost})
    this.$store.dispatch('GetIndexList', {url: '/_cat/indices', host: this.serverhost})
  }
}
</script>

<style>
  .container-fluid{
    background-color: #2c3e50;
  }
  #elastic-logo-color {
    color:#1DC07F
  }
  #elastic-logo-color:hover{
    color:#19A26B
  }
  #elastic-host {
    width:14em;
    color:#181A1B;
  }
  .input-group .input-group-addon{
    background: #F6F7F7;
    color:#1DC07F;
  }
</style>

