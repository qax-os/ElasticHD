<template>
  <section id="mapping-view">
    <el-row :gutter="0">
      <el-col :xs="24" :sm="24" :md="8" :lg="8">
        <el-card :body-style="{ padding: '0px', height: '300px' }">
          <img src="../assets/hero.svg" class="image">
          <div style="background-color:#D3DCE6;">
            <div style="text-align:center;">Indices Templates</div>
            <hr class="mapping-hr">
            <p class="text-content" style="text-align:center;">
              Index templates allow you to define templates that will automatically be applied when new indices are created. The templates include both settings and mappings, and a simple pattern template that controls whether the template should be applied to the new index.
            </p>
            <div style="text-align:center;height:50px;line-height:50px;">
              <el-button :plain="true" type="info" @click="mappingDirect">Learn More</el-button>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="12" :md="8" :lg="8">
        <div>
          <img src="../assets/rocket.svg" class="image-rocket">
        </div>
      </el-col>
      <el-col :xs="24" :sm="12" :md="8" :lg="8">
        <div>
          <el-upload class="upload-demo" drag :action="action" :on-change="handleChange" :file-list="fileList" :before-upload="beforeFileUpload" :on-success="fileUploadSuccess" :on-error="fileUploadError" style="width:400px;">
            <i class="el-icon-upload"></i>
            <div class="el-upload__text">将文件拖到此处，或
              <em>点击上传</em>
            </div>
            <div class="el-upload__tip" style="text-align:center; background-color:#D3DCE6;margin-top:0px; width:360px" slot="tip">更新ElasticSerach Index Template: 要求只能上传application/json文件、不超过2M、且文件名是对应的Index Template名称</div>
          </el-upload>
        </div>
      </el-col>
    </el-row>
    <el-row :gutter="10">
      <el-col :xs="24" :sm="24" :md="24" :lg="24" style="background-color:#ffffff; padding: 0;">
        <p class="mapping-p">
          <b>Index Templates中Type个数统计</b>
        </p>
        <div class="chart" id="temp-chartx" style="width: 100%; height:275px; "></div>
      </el-col>
    </el-row>
    <el-row :gutter="0">
      <el-col :xs="24" :sm="24" :md="24" :lg="24">
        <el-table :data="IndexTemplateList" border style="width: 100%">
          <el-table-column label="Index Template List【索引模板列表】">
            <el-table-column prop="name" label="Name" width="220"></el-table-column>
            <el-table-column prop="type" label="Types"></el-table-column>
            <el-table-column label="操作" width="200">
              <template scope="scope">
                <el-button size="small" type="success" @click="handleSee(scope.$index, scope.row)">查看</el-button>
                <el-popover ref="popover{{$index}}" placement="top" width="160" v-model="scope.row.visible" trigger="click">
                  <p>这个模板确定删除吗？Are you sure?</p>
                  <div style="text-align: right; margin: 0">
                    <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
                    <el-button type="primary" size="mini" @click="handleDelete(scope.$index, scope.row)">确定</el-button>
                  </div>
                </el-popover>
                <el-button size="small" type="danger" v-popover:popover{{$index}}>删除</el-button>
                <!--<el-button size="small" type="danger" @click="handleDelete(scope.$index, scope.row)">删除</el-button>-->
              </template>
            </el-table-column>
          </el-table-column>
        </el-table>
      </el-col>
    </el-row>
    <el-dialog title="Index Mapping Info" v-model="showMapping">
      <div id="mappinginfo">{{ mappingInfo}}</div>
    </el-dialog>
  </section>
</template>


<script>
import $ from 'jquery'
import echarts from 'echarts'
import { mapGetters } from 'vuex'
export default {
  name: 'mapping-view',
  data: function () {
    return {
      currentDate: new Date(),
      action: '/_upload/template',
      fileList: [],
      showMapping: false,
      jsonView: false
    }
  },
  computed: {
    ...mapGetters([
      'serverhost',
      'templateInfo',
      'IndexTemplateList',
      'toggleMapping',
      'mappingInfo',
      'DeleteMappingSuccess',
      'DeleteMappingError'
    ])
  },
  created: function () {
    this.action = '/_upload/template?serverhost=' + this.serverhost
  },
  mounted: function () {
    echarts.init(document.getElementById('temp-chartx')).setOption(this.tempChart(this.templateInfo))
  },
  watch: {
    DeleteMappingSuccess () {
      this.$message.success(this.DeleteMappingSuccess)
      this.$store.dispatch('GetTemplateInfo', {url: '/_template', host: this.serverhost})
    },
    DeleteMappingError () {
      this.$message.error(this.DeleteMappingError)
      // this.$store.dispatch('GetTemplateInfo', {url: '/_template', host: this.serverhost})
    },
    serverhost (val, oldval) {
      this.action = '/_upload/template?serverhost=' + this.serverhost
    },
    templateInfo (val, oldval) {
      // this.$store.dispatch('GetTemplateInfo', {url: '/_template', host: this.serverhost})
      echarts.init(document.getElementById('temp-chartx')).setOption(this.tempChart(this.templateInfo))
    },
    fileList () {
      this.$store.dispatch('GetTemplateInfo', {url: '/_template', host: this.serverhost})
    },
    mappingInfo (val, oldval) {
      this.showMapping = true
      this.handleShow()
    }
  },
  methods: {
    mappingDirect () {
      window.open('https://www.elastic.co/guide/en/elasticsearch/reference/5.4/indices-templates.html')
    },
    handleShow (data) {
      $('#mappinginfo').JSONView(this.mappingInfo)
    },
    handleSee (index, row) {
      var tmp = {'url': '/_template/handle', 'method': 'GET', 'indices': row.name, 'toggle': this.toggleMapping, host: this.serverhost}
      this.$store.dispatch('HandleIndexTemplate', tmp)
    },
    handleDelete (index, row) {
      row.visible = false
      var tmp = {'url': '/_template/handle', 'method': 'DELETE', 'indices': row.name, 'toggle': this.toggleMapping, host: this.serverhost}
      this.$store.dispatch('HandleIndexTemplate', tmp)
    },
    fileUploadError (file, error) {
      this.$message.error(error.name + ' 模板部署失败，请打开开发者模式查看错误信息提示')
    },
    fileUploadSuccess (res, file) {
      // console.log(res, file, 2222)
      this.$message.success(file.name + ' 模板部署成功')
    },
    handleChange (file, fileList) {
      this.fileList = fileList.slice(-3)
    },
    beforeFileUpload (file) {
      var isLt2M = file.size / 1024 / 1024 < 2
      var isJson = file.type === 'application/json'
      if (file.type === '') {
        isJson = true
      }
      console.log(file)
      if (!isJson) {
        this.$message.error('上传的文件类型必须是application/json!')
      }
      if (!isLt2M) {
        this.$message.error('上传文件大小不能超过2MB!')
      }
      return isLt2M && isJson
    },
    tempChart: function (templateInfo) {
      var tempOption = {
        color: ['#3398DB'],
        tooltip: {
          trigger: 'axis',
          axisPointer: {
            type: 'shadow'
          }
        },
        grid: {
          left: '3%',
          right: '4%',
          bottom: '3%',
          containLabel: true
        },
        xAxis: [{
          type: 'category',
          data: templateInfo.templateName,
          axisTick: {
            alignWithLabel: true
          }
        }],
        yAxis: [{
          axisTick: {
            alignWithLabel: true
          }
        }],
        series: [{
          name: '表数量',
          type: 'bar',
          barWidth: '40%',
          data: templateInfo.typeCount
        }],
        label: {
          normal: {
            show: true,
            position: 'top',
            formatter: '{c}'
          }
        },
        itemStyle: {
          normal: {
            color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [{
              offset: 0,
              color: 'rgba(17, 168,171, 1)'
            },
            {
              offset: 1,
              color: 'rgba(17, 168,171, 0.1)'
            }]),
            shadowColor: 'rgba(0, 0, 0, 0.1)',
            shadowBlur: 10
          }
        }
      }
      return tempOption
    }
  }
}
</script>

<style>
.el-upload__input {
    display: none !important;
}
.image-rocket{
  width: 100%;
  display: block;
  height:300px;
  background-color: #fff;
}
.image {
    width: 100%;
    display: block;
  }
.text-content {
    color: black;
    margin: 20px 0;
    font-size: 12px;
    text-align:justify;
}
/*p {
    margin: 0 0 0 0 !important;
}*/
.mapping-p{
    word-wrap: normal;
    text-overflow: ellipsis;
    line-height: 39px;
    width: 100%;
    box-sizing: border-box;
    background-color: #eef1f6;
    border: 1px solid #dfe6ec;
    margin-right: 20px;
    padding: 0 20px;
}
.mapping-hr{
    margin-top: 0px;
    margin-bottom: 0px;
    background-color: #99A9BF;
}
</style>

