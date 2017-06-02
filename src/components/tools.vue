<template>
  <section id="sqltools">
      <el-row>
        <el-col :span="24">
            <el-collapse v-model="activeName" accordion>
                <el-collapse-item title="SQL Convert DSL" name="1">
                    <el-row>
                        <el-col :span="24">
                            <el-input type="textarea" :rows="5" placeholder="请输入内容" v-model="textarea"></el-input>
                        </el-col>
                        <el-col :span="24">
                            <div style = "text-align:right;">
                                <el-button type="info" v-on:click="sqlExplain">Explain</el-button>
                            </div>
                        </el-col>
                        <el-col :span="24">
                            <div style = "color:red;">
                                {{currentError}}
                            </div>
                        </el-col>
                        <el-col :span="24">
                            <div style="display:none;" id="dslBody">
                                <el-input type="textarea" :rows="5" placeholder="请输入内容" v-model="dslBody"></el-input>
                            </div>
                        </el-col>
                    </el-row>
                </el-collapse-item>
                <el-collapse-item title="反馈 Feedback" name="2">
                    <div>控制反馈：通过界面样式和交互动效让用户可以清晰的感知自己的操作；</div>
                    <div>页面反馈：操作后，通过页面元素的变化清晰地展现当前状态。</div>
                </el-collapse-item>
                <el-collapse-item title="效率 Efficiency" name="3">
                    <div>简化流程：设计简洁直观的操作流程；</div>
                    <div>清晰明确：语言表达清晰且表意明确，让用户快速理解进而作出决策；</div>
                    <div>帮助用户识别：界面简单直白，让用户快速识别而非回忆，减少用户记忆负担。</div>
                </el-collapse-item>
                <el-collapse-item title="可控 Controllability" name="4">
                    <div>用户决策：根据场景可给予用户操作建议或安全提示，但不能代替用户进行决策；</div>
                    <div>结果可控：用户可以自由的进行操作，包括撤销、回退和终止当前操作等。</div>
                </el-collapse-item>
            </el-collapse>
        </el-col>
    </el-row>
  </section>
</template>

<script>
import { mapGetters } from 'vuex'
export default {
  name: 'sqltools',
  data: function () {
    return {
      textarea: 'SELECT * FROM myindex',
      activeName: 1
    }
  },
  computed: {
    ...mapGetters([
      'dslBody',
      'sqlBody',
      'currentError',
      'serverhost'
    ])
  },
  created: function () {
  },
  mounted () {
  },
  methods: {
    sqlExplain () {
      var tmpVal = {url: '/_sql2dsl', sql: this.textarea, host: this.serverhost}
      this.$store.dispatch('ConvertSQLToDSL', tmpVal)
      document.getElementById('dslBody').style.display = 'block'
    }
  }
}
</script>

<style>
</style>

