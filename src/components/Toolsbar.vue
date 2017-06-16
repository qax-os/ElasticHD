<template>
    <div id="toolsbar-view">
        <section id='tools'>
            <ul class='breadcrumb' id='breadcrumb'>
                <li class='title' v-html="toggleTitle">{{ toggleTitle }}</li>
            </ul>
            <div id='toolbar'>
                <div class='btn-group'>
                    <a class="btn" v-for="(item, index) in btnGroup" @click="getInfo (index, item.value)">
                        <el-tooltip v-bind:content="item.toolsTip" placement="top" effect="dark">
                        <i v-bind:class="item.fontAwesome"></i>
                        </el-tooltip>
                    </a>
                </div>
            </div>
        </section>
        <el-dialog title="提示" v-model="dialogVisible" size="tiny" :before-close="handleClose">
            <div><tree-view :data="info" :options="{maxDepth: 5}"></tree-view></div>
            <span slot="footer" class="dialog-footer">
                <el-button @click="dialogVisible = false">取 消</el-button>
                <el-button type="primary" @click="dialogVisible = false">确 定</el-button>
            </span>
        </el-dialog>
    </div>
</template>

<script>
import { mapGetters } from 'vuex'
export default {
  name: 'toolsbar-view',
  data: function () {
    return {
      dialogVisible: false,
      info: ''
    }
  },
  computed: {
    ...mapGetters([
      'toggleTitle',
      'activeIndex',
      'btnGroup',
      'health'
    ])
  },
  methods: {
    getInfo: function (index, value) {
      this.dialogVisible = true
      if (value === 'tools-info') {
        this.info = this.health
      }
    },
    handleClose (done) {
      this.$confirm('确认关闭？')
      .then(_ => {
        done()
      })
      .catch(_ => {})
    }
  },
  created: function () {
    this.$store.dispatch('ToggleTitle', this.activeIndex)
  },
  watch: {
    activeIndex (val, oldval) {
      this.$store.dispatch('ToggleTitle', val)
    }
  }
}
</script>

<style>
#tools {
    margin-left: 80px;
}

#tools #breadcrumb {
    -webkit-border-radius: 0;
    -moz-border-radius: 0;
    -ms-border-radius: 0;
    -o-border-radius: 0;
    border-radius: 0;
    padding: 15px 20px 15px 20px;
    border-bottom: 1px solid #DFE7EC
}

#tools #breadcrumb .title {
    font-weight: 800;
    color: #1abc9c;
}
#tools .breadcrumb{
    /*// 2c3e50*/
    background-color: #f1f4f7;
}

#tools #toolbar {
    margin-top: -70px;
    margin-right: 10px;
    float: right;
    height: 50px;
    padding: 8px;
    border-bottom: 0
}

#tools .btn-group,
#tools .label,
#tools form {
    float: right;
    display: inline-block
}

#tools .btn-group .btn span.badge {
    background-color: #2c3e50;
    padding: 3px 5px;
    font-size: 0.7em;
    font-weight: 200
}

#tools .label {
    margin-right: 5px;
    margin-top: 6px;
    padding: 5px;
    height: 23px;
    font-size: 0.9em;
    font-weight: 200
}
#tools .btn-group .btn {
    color: white;
    background-color: #1abc9c;
    border-color: #17a689
}
.btn-success:hover,
#tools .btn-group .btn:hover,
.btn-success:focus,
#tools .btn-group .btn:focus,
.btn-success:active,
#tools .btn-group .btn:active,
.btn-success.active,
#tools .btn-group .active.btn,
.open>.btn-success.dropdown-toggle,
#tools .btn-group .open>.dropdown-toggle.btn {
    color: white;
    background-color: #148f77;
    border-color: #0f705d
}

</style>
