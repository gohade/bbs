<template>
    <el-row type="flex" justify="center" align="middle">
      <el-col :span="8">
        <el-row>
          <el-col :span="8">
            <el-image
                style="height: 60px"
                src="http://tuchuang.funaio.cn/hade/hadecast_logo_v2.png"
                :fit="fit"></el-image>
          </el-col>
          <el-col :span="8" :offset="8">
            <el-avatar style="margin-right: 10px" icon="el-icon-user-solid"></el-avatar>
          </el-col>
        </el-row>
      </el-col>
    </el-row>
</template>

<script>
export default {
  methods: {

  }
}
</script>

<style scoped>

.el-header {
  background-color: #409EFF;
  text-align: center;
  line-height: 80px;
}

</style>
