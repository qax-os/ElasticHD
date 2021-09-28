<template>
  <section id="help">
     <div v-html="compiledMarkdown" style="background-color:white;"></div>
  </section>
</template>

<script>
import { mapGetters } from 'vuex'
import Marked from 'marked'
export default {
  name: 'help',
  data: function () {
    return {
      content: `敬请期待`
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

