<template>
  <el-container>
    <el-header>
      <BBSHeader></BBSHeader>
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
    BBSHeader
  },
  methods: {

  }
}
</script>

<style scoped>

.el-header {
  background-color: #d3dce6;
  text-align: center;
  line-height: 80px;
}


.el-main {
  /*background-color: #E9EEF3;*/
  /*color: #333;*/
  /*text-align: center;*/
}

body > .el-container {
  margin-bottom: 40px;
}

</style>
