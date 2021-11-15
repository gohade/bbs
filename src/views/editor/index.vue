<template>
  <el-row type="flex" justify="center" align="middle">
    <el-col :span="8">
      <el-card class="box-card" shadow="never">
        <div v-for="o in 4" :key="o" class="text item">
          {{'列表内容 ' + o }}
        </div>
      </el-card>
      <el-divider content-position="left">回复</el-divider>
      <el-card class="box-card" shadow="never">
        <div v-for="o in 4" :key="o" class="text item">
          {{'列表内容 ' + o }}
        </div>
      </el-card>

    </el-col>
  </el-row>
</template>

<style scoped>
.text {
  font-size: 14px;
}

.item {
  padding: 18px 0;
}

.box-card {
  /*width: 480px;*/
}
</style>
