<template>
   <section id="indices-view">
     <el-col :xs="24" :sm="24" :md="24" :lg="24">
                    <el-table :data="indexList.slice(0, showLen)" border style="width: 100%">
                        <el-table-column label="Index list" >
                            <el-table-column prop="index" label="Index" width="220" sortable></el-table-column>
                            <el-table-column prop="docs" label="索引文档个数" sortable ></el-table-column>
                            <el-table-column prop="size" label="Store Size" sortable></el-table-column>
                            <el-table-column prop="prisize" label="主分片大小" ></el-table-column>
                            <el-table-column prop="health" label="索引健康状态" ></el-table-column>
                            <el-table-column prop="uuid" label="UUID" ></el-table-column>
                            <el-table-column prop="delete" label="索引删除文档个数" ></el-table-column>
                            <el-table-column prop="status" label="Status"></el-table-column>
                            <el-table-column label="操作">
                            <template scope="scope">
                                <el-button size="small" type="danger" @click="handleDelete(scope.$index, scope.row)">删除</el-button>
                            </template>
                            </el-table-column>
                        </el-table-column>
                    </el-table>
                 </el-col>
        <p style="text-align:center;" v-if="indexList.length > showLen" v-on:click="addShowLen"><a style="margin:0 auto;">加载更多...</a></p>
   </section>
</template>

<script>
import { mapGetters } from 'vuex'
export default {
  name: 'indices-view',
  data: function () {
    return {
      showLen: 20
    }
  },
  computed: {
    ...mapGetters([
      'indexList'
    ])
  },
  methods: {
    addShowLen () {
      this.showLen = this.showLen + 10
    }
  }
}
</script>

<style>

</style>

