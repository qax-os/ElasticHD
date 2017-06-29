<template>
  <section id="indices-view">
    <el-row :gutter="10">
      <el-col :xs="24" :sm="24" :md="24" :lg="24"> 
  <el-input placeholder="请输入要查询的索引" icon="search" v-model="inputIndexName" :on-icon-click="handleIconClick">
  </el-input>
        <el-table :data="list.slice(0, showLen)" border style="width: 100%">
            <el-table-column prop="index" label="Index"  sortable show-overflow-tooltip></el-table-column>
            <el-table-column prop="docs" label="Docs" sortable show-overflow-tooltip></el-table-column>
            <el-table-column prop="size" label="Size" sortable show-overflow-tooltip></el-table-column>
            <el-table-column prop="prisize" label="主分片大小" show-overflow-tooltip></el-table-column>
            <el-table-column prop="health" label="Status" show-overflow-tooltip></el-table-column>
            <el-table-column prop="uuid" label="UUID" show-overflow-tooltip></el-table-column>
            <el-table-column prop="delete" label="Del Docs" show-overflow-tooltip></el-table-column>
            <el-table-column prop="status" label="Status" show-overflow-tooltip></el-table-column>
            <el-table-column label="操作">
              <template scope="scope">
                <el-popover ref="popover{{$index}}" placement="top" width="160" v-model="scope.row.visible" trigger="click">
                  <p>这个索引确定删除吗？Are you sure?</p>
                  <div style="text-align: right; margin: 0">
                    <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
                    <el-button type="primary" size="mini" @click="handleDelete(scope.$index, scope.row)">确定</el-button>
                  </div>
                </el-popover>
                <el-button size="small" type="danger" v-popover:popover{{$index}} >删除</el-button>
              </template>
            </el-table-column>
        </el-table>
      </el-col>
    </el-row>
    <p style="text-align:center;" v-if="indexList.length > showLen" v-on:click="addShowLen"><a style="margin:0 auto;">加载更多...</a></p>
  </section>
</template>


<script>
  import {
    mapGetters
  } from 'vuex'
  export default {
    name: 'indices-view',
    data: function () {
      return {
        showLen: 20,
        inputIndexName: '',
        list: []
      }
    },
    computed: {
      ...mapGetters([
        'indexList',
        'serverhost',
        'DeleteIndexSuccess',
        'DeleteIndexError'
      ])
    },
    watch: {
      DeleteIndexSuccess () {
        this.$message.success(this.DeleteIndexSuccess)
      },
      DeleteIndexError () {
        this.$message.error(this.DeleteIndexError)
      },
      indexList () {
        this.list = this.indexList
      }
    },
    methods: {
      filterItems (query) {
        var arr = this.indexList.filter(function (el) {
          return el['index'].toLowerCase().indexOf(query.toLowerCase()) > -1
        })
        return arr
      },
      handleIconClick () {
        if (this.inputIndexName === '') {
          this.list = this.indexList
          return false
        }
        this.list = this.filterItems(this.inputIndexName)
      },
      addShowLen () {
        this.showLen = this.showLen + 10
      },
      handleDelete (index, row) {
        row.visible = false
        var tmp = {'url': '/_indices/delete', 'index': row.index, host: this.serverhost}
        this.$store.dispatch('DeleteIndexByName', tmp)
      }
    },
    created: function () {
      this.list = this.indexList
    }
  }

</script>

<style>
.el-table .cell {
    white-space: normal;
    /*word-break: break-all;*/
    line-height: 24px;
    overflow:hidden !important;
    white-space:nowrap !important;
    text-overflow:ellipsis !important;
}
</style>
