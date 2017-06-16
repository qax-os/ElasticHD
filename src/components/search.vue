<template>
    <div id="search">
        <section>
           <el-row :gutter="10">
              <el-col :xs="24" :sm="24" :md="10" :lg="7">
                  <div>
                      <el-card class="box-card">
                          <div slot="header" class="clearfix">
                            <span><i class="fa fa-search"></i> 搜索</span>
                            <!--<a  type="primary">History</a>-->
                            <span style="float: right;" v-on:click="history">
                               <a> <i class="fa fa-history"></i> 历史查询</a>
                            </span>
                            <span style="float: right; padding-right:10px" v-on:click="toolsTip">
                                <a><i class="fa fa-lightbulb-o"></i> 提示</a>
                            </span>
                          </div>
                          <div>
                              <el-row :gutter="0">
                                    <el-col :xs="12" :sm="12" :md="12" :lg="12">
                                      <el-input v-model="inputIndex" placeholder="请输入要查询的索引"></el-input>
                                    </el-col>
                                    <el-col :xs="12" :sm="12" :md="12" :lg="12">
                                         <el-input placeholder="请输入要查询的type" v-model="inputType">
                                            <template slot="prepend">/</template>
                                        </el-input>
                                    </el-col>
                              </el-row>
                               <el-row :gutter="0">
                                    <el-col :xs="15" :sm="15" :md="15" :lg="16">
                                         <el-input v-model="inputOptions" placeholder="请输入搜索参数"></el-input>
                                    </el-col>
                                    <el-col :xs="9" :sm="9" :md="9" :lg="8">
                                         <el-select v-model="inputMethod" clearable >
                                            <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value"> </el-option>
                                         </el-select>
                                    </el-col>
                              </el-row>
                              <el-row :gutter="0">
                                    <el-col :xs="24" :sm="24" :md="24" :lg="24">
                                         <el-input type="textarea" :rows="rows" placeholder="请输入内容" v-model="inputTextarea"></el-input>
                                    </el-col>
                              </el-row>
                              <el-row :gutter="0">
                                    <el-col :xs="24" :sm="24" :md="24" :lg="24">
                                         <el-button-group>
                                            <el-button type="primary" icon="search" size="small" v-on:click="search">查询</el-button>
                                            <el-button type="primary" icon="check" size="small" v-on:click="pretty">验证JSON</el-button>
                                        </el-button-group>
                                    </el-col>
                              </el-row>
                               <el-row :gutter="0">
                                    <el-col :xs="24" :sm="24" :md="24" :lg="24">
                                        <span style="color:red;">{{errorMessage}}</span>
                                    </el-col>
                              </el-row>
                          </div>
                      </el-card>
                  </div>
              </el-col>
              <el-col :xs="24" :sm="24" :md="14" :lg="17">
                  <div>
                      <el-card class="box-card">
                          <div slot="header" class="clearfix">
                            <span><i class="fa fa-search"></i>查询结果</span>
                          </div>
                          <div>
                            <div id="spinner" v-if="showSpinner"><i class="fa fa-spinner fa-spin fa-4x" ></i></div>
                            <div id="jsonView" >{{resultMessage}}</div>
                          </div>
                      </el-card>
                </div>
              </el-col>
            </el-row>
        </section>
        <el-dialog title="历史记录" v-model="showhistory">
          <el-table :data="historyList">
            <el-table-column type="expand">
              <template scope="props">
                <el-form label-position="left" inline class="demo-table-expand">
                  <el-form-item label="日期">
                    <span>{{ props.row.date }}</span>
                  </el-form-item>
                  <el-form-item label="Index">
                    <span>{{ props.row.indices }}</span>
                  </el-form-item>
                  <el-form-item label="Type">
                    <span>{{ props.row.types }}</span>
                  </el-form-item>
                  <el-form-item label="Method">
                    <span>{{ props.row.method }}</span>
                  </el-form-item>
                  <el-form-item label="Options">
                    <span>{{ props.row.options }}</span>
                  </el-form-item>
                  <el-form-item label="Body">
                    <span>{{ props.row.body }}</span>
                  </el-form-item>
                </el-form>
              </template>
            </el-table-column>
            <el-table-column property="date" label="日期" sortable></el-table-column>
            <el-table-column property="indices" label="Index" ></el-table-column>
            <el-table-column property="types" label="Type" ></el-table-column>
            <el-table-column label="操作">
                <template scope="props">
                  <el-button  size="small" type="danger" @click="handleSearch(props.$index, props.row)">搜索</el-button>
                </template>
            </el-table-column>
          </el-table>
        </el-dialog>
    </div>
</template>

<script>
import { mapGetters } from 'vuex'
import $ from 'jquery'
export default {
  name: 'search',
  data: function () {
    return {
      errorMessage: '',
      resultMessage: '暂无数据',
      inputMethod: 'GET',
      inputIndex: '',
      inputType: '',
      inputOptions: '_search',
      inputTextarea: '{"query":{"match_all":{}}}',
      rows: 20,
      showhistory: false,
      historyList: []
    }
  },
  computed: {
    ...mapGetters([
      'options',
      'result',
      'serverhost',
      'showSpinner'
    ])
  },
  watch: {
    result (val, oldval) {
      try {
        $('#jsonView').JSONView(this.result)
        this.resultMessage = ''
      } catch (e) {
        this.resultMessage = this.result
      }
    }
  },
  methods: {
    handleSearch: function (index, row) {
      this.showhistory = false
      this.inputMethod = row.method
      this.inputIndex = row.indices
      this.inputType = row.types
      this.inputOptions = row.options
      this.inputTextarea = row.body
      this.search()
    },
    history: function () {
      this.showhistory = true
      this.historyList = this.getLocalStorage().history
    },
    toolsTip: function () {
      alert(`toolstip`)
    },
    search: function () {
      try {
        JSON.stringify(JSON.parse(this.inputTextarea), null, 5)
        this.errorMessage = ''
      } catch (e) {
        if (this.inputMethod !== 'GET') {
          this.errorMessage = e.name + ': ' + e.message
          return
        }
      }
      // $('#jsonView').JSONView(null)
      // var list = window.localStorage ? localStorage.getItem('elasticHDHistor') : new Array()
      var tmpVal = {url: '/_search', date: this.getTime(), host: this.serverhost, method: this.inputMethod, indices: this.inputIndex, types: this.inputType, options: this.inputOptions, body: this.inputTextarea}
      this.setLocalStorage(tmpVal)
      this.$store.dispatch('SetShowSpinner', true)
      this.$store.dispatch('CURLSerachResult', tmpVal)
    },
    pretty: function () {
      try {
        this.inputTextarea = JSON.stringify(JSON.parse(this.inputTextarea), null, 5)
        this.errorMessage = ''
      } catch (e) {
        this.errorMessage = e.name + ': ' + e.message
      }
    },
    getLocalStorage: function () {
      if (!window.localStorage) {
        return {'history': []}
      }
      if (!localStorage.getItem('elasticHDHistor')) {
        localStorage.setItem('elasticHDHistor', `{"history":[]}`)
      }
      return JSON.parse(localStorage.getItem('elasticHDHistor'))
    },
    setLocalStorage: function (data) {
      if (!window.localStorage) {
        return
      }
      if (!localStorage.getItem('elasticHDHistor')) {
        localStorage.setItem('elasticHDHistor', `{"history":[]}`)
      }
      var history = JSON.parse(localStorage.getItem('elasticHDHistor'))
      if (history.history.length >= 15) {
        history.history.shift()
      }
      history.history.push(data)
      localStorage.setItem('elasticHDHistor', JSON.stringify(history))
    },
    getTime: function () {
      var date = new Date()
      var seperator1 = '-'
      var seperator2 = ':'
      var month = date.getMonth() + 1
      var strDate = date.getDate()
      if (month >= 1 && month <= 9) {
        month = '0' + month
      }
      if (strDate >= 0 && strDate <= 9) {
        strDate = '0' + strDate
      }
      var currentdate = date.getFullYear() + seperator1 + month + seperator1 + strDate + ' ' + date.getHours() + seperator2 + date.getMinutes() + seperator2 + date.getSeconds()
      return currentdate
    }
  }
}
</script>

<style>
.el-card__header {
    padding: 10px 20px !important;
    border-bottom: 1px solid #ccc !important;
    box-sizing: border-box;
    background-color: #f1f4f7;
}
#spinner {
  text-align:center;
  margin: 0 auto;
  color: #ccc;
}
</style>
