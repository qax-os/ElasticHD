<template>
<div id="sidebar-view">
 <section id='sidebar'>
    <i class='fa fa-reorder fa-2x' @click='toggle' id="toggle"></i>
        <ul id='dock'>
            <li v-for="(item, index) in items" @click="changeStyle (index)" :class="{'active': index === activeIndex}" class='launcher'>
                <i v-bind:class='item.fontAwesome'></i>
                    <a v-bind:id="item.value" v-if="toggleShow">{{item.value}}</a>
                    <a v-bind:id="item.value" v-else></a>
                </li>
            </li>
        </ul>
        <el-tooltip v-bind:content="msg" placement="top" effect="dark">
           <div id='beaker'></div>
        </el-tooltip>
    </section>
</div>
</template>

<script>
// import Vue from 'vue'
import { mapGetters } from 'vuex'
export default {
  name: 'sidebar-view',
  data: function () {
    return {
    }
  },
  computed: {
    ...mapGetters([
      'msg',
      'items',
      'toggleShow',
      'activeIndex'
    ])
  },
  methods: {
    toggle: function () {
      this.$store.dispatch('ToggleShow', this.$store.getters.toggleShow)
    },
    changeStyle: function (index) {
      this.$store.dispatch('SetActiveIndex', index)
    }
  }
}
</script>
<style>
#sidebar {
    position: absolute;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0
}

#sidebar {
    background-color: #2c3e50;
    width: 80px;
    z-index: 1;
    text-align: center
}

#sidebar i#toggle {
    margin-top: 20px;
    display: inline-block;
    color: #3e5771
}

#sidebar ul#dock {
    margin: -15px 0 0 0;
    padding: 0;
    width: 80px;
    text-align: center;
    color: #ecf0f1;
    border-right: 1px solid #1a242f
}

#sidebar ul#dock li.launcher {
    list-style: none;
    margin-top: 25px;
    cursor: pointer
}

#sidebar ul#dock li.launcher>i {
    display: block;
    font-size: 2.5em;
    margin-bottom: -2px
}

#sidebar ul#dock li.launcher>a {
    height: 36px;
    display: block;
    position: relative;
    padding-top: 36px;
    margin-top: -36px;
    text-decoration: none;
    color: #ecf0f1;
    font-size: 0.8em
}

#sidebar ul#dock li.launcher ul.dropdown-menu {
    width: 180px;
    border-left-width: 0;
    text-align: left;
    position: absolute;
    margin-left: 85px;
    top: 8px
}

#sidebar ul#dock li.launcher ul.dropdown-menu>li>a {
    padding: 5px 20px
}

#sidebar ul#dock li.launcher ul.dropdown-menu a,
#sidebar ul#dock li.launcher ul.dropdown-menu li.dropdown-header {
    display: inline-block;
    width: 100%;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis
}

#sidebar ul#dock li.launcher ul.dropdown-menu li.dropdown-header {
    -webkit-border-radius: 4px;
    -moz-border-radius: 4px;
    -ms-border-radius: 4px;
    -o-border-radius: 4px;
    border-radius: 4px;
    color: #1abc9c;
    margin-top: -5px;
    padding: 5px 20px
}

#sidebar ul#dock li.launcher ul.dropdown-menu li.dropdown-header:before {
    -webkit-transform: rotate(45deg);
    -ms-transform: rotate(45deg);
    -o-transform: rotate(45deg);
    transform: rotate(45deg);
    position: absolute;
    margin-top: 4px;
    margin-left: -25px;
    width: 10px;
    height: 10px;
    background-color: #2c3e50;
    content: " ";
    display: inline-block
}

#sidebar ul#dock li.launcher:hover>i,
#sidebar ul#dock li.launcher:hover>a,
#sidebar ul#dock li.launcher.active>i,
#sidebar ul#dock li.launcher.active>a {
    -webkit-transition-property: color;
    transition-property: color;
    -webkit-transition-duration: 0.4s;
    transition-duration: 0.4s;
    -webkit-transition-timing-function: ease-in-out;
    transition-timing-function: ease-in-out;
    color: #1abc9c
}

#sidebar ul#dock li.launcher:hover:before,
#sidebar ul#dock li.launcher.active:before {
    -webkit-transform: rotate(45deg);
    -ms-transform: rotate(45deg);
    -o-transform: rotate(45deg);
    transform: rotate(45deg);
    content: "";
    display: inline-block;
    width: 10px;
    height: 10px;
    background-color: white;
    position: absolute;
    right: -6px;
    margin-top: 22px
}
#beaker {
    position: absolute;
    bottom: 5px;
    background-position: 0 0;
    width: 12px;
    height: 24px;
    background-size: 100px;
    background-repeat: no-repeat;
    margin-left: 33px
}
#beaker {
    background-image: url(../assets/sprite-8460f675.png)
}
.tooltip {
  display: none;
  opacity: 0;
  transition: opacity .15s;
  pointer-events: none;
  padding: 4px;
  z-index: 10000;
}

.tooltip .tooltip-content {
  background: black;
  color: white;
  border-radius: 16px;
  padding: 5px 10px 4px;
}

.tooltip.tooltip-open-transitionend {
  display: block;
}

.tooltip.tooltip-after-open {
  opacity: 1;
}
</style>
