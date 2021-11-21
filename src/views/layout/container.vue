<template>
  <el-container>
    <el-header>
      <bbs-header></bbs-header>
    </el-header>
    <el-main>
        <router-view></router-view>
    </el-main>
  </el-container>
</template>

<script>
import BBSHeader from './header'

export default {
  components: {
    'bbs-header': BBSHeader
  },
  methods: {

  }
}
</script>

<style scoped>

.el-header {
  background-color: #fff;
  text-align: center;
  line-height: 1.6;
  z-index: 1000;
  height: 52px;
}


.el-main {
  background-color: #E9EEF3;
  /*color: #333;*/
  /*text-align: center;*/
}

body > .el-container {
  margin-bottom: 40px;
}

</style>
