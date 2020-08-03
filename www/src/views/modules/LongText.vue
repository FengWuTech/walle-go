<template>
  <div style="cursor: pointer;">
    <span v-if="!content||content.length<=maxWord">{{ content || '-' }}</span>
    <a-popover v-else :title="title">
      <template slot="content">
        <p :style="`max-width: ${maxWord}em;`" class="content-p">{{ content }}</p>
      </template>
      <span :style="`max-width: ${maxWord}em`" class="ellipsis">{{ content }}</span>
    </a-popover>
  </div>
</template>

<script>
/**
 * 长文本展示组件，默认长度为15个中文字符
 * title：弹框展示标题
 * content：文本内容
 * maxWord：显示最长文本长度
 */
export default {
  name: 'LongText',
  props: {
    title: {
      type: String,
      default: ''
    },
    content: {
      required: true,
      type: String,
      default: ''
    },
    maxWord: {
      type: Number,
      default: 15
    }
  }
}
</script>

<style scoped lang="less">
  .content-p {
    word-break: break-all;
    overflow: hidden;
  }
  .ellipsis {
    display: block;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }
</style>
