<template>
  <section id="help">
     <div v-html="compiledMarkdown"></div>
  </section>
</template>

<script>
import { mapGetters } from 'vuex'
import Marked from 'marked'
export default {
  name: 'help',
  data: function () {
    return {
      content: `#### 搜索
      `
    }
  },
  computed: {
    ...mapGetters([
    ]),
    compiledMarkdown: function () {
      return Marked(this.content, { sanitize: true })
    }
  },
  created: function () {
  },
  mounted () {
  }
}
</script>

<style>
</style>

