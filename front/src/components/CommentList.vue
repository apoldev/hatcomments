<template>
  <div class="comment-wrapper" >
    <Comment
        v-for="c in comments"
        :key="c.id"
        :comment="c"
        @onReplyTo="this.$emit('onReplyTo', $event)"
        @onClickReply="onClickReply($event)"
        @onCallParent="this.$emit('onCallParent', $event)"
    />
  </div>
</template>

<script>
import Comment from './Comment'
export default {
  name: 'TextInput',
  components: {
    Comment
  },
  props: {
    comments: Array
  },
  methods: {

    onClickReply: function(c){
      c['reply'] = !c.reply
    }

  }
}
</script>

<style scoped>

</style>