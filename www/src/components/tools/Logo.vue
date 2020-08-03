<template>
  <div class="logo" :class="className">
    <router-link to="/index">
      <img src="~@/assets/fwlogo.png" class="mylogo" alt="logo">
      <h1 style="font-weight: normal;font-size: 18px;" v-if="showTitle">{{ title }}</h1>
    </router-link>
  </div>
</template>

<script>
export default {
  name: 'Logo',
  props: {
    title: {
      type: String,
      default: '系统部署',
      required: false
    },
    showTitle: {
      type: Boolean,
      default: true,
      required: false
    },
    className: {
      type: String,
      default: '',
      required: false
    }
  }
}
</script>
<style type="text/css" lang="less">
  .mylogo {
    width: 22px !important;
    height: inherit !important;
  }
  .top-nav-header-blank{
    display: inline-block;
    padding: 0  24px;
  }
  .top-nav-header-blank h1{
    display: inline-block;
    padding: 0  12px;
    color: white;
    font-size: 24px;
  }
</style>
