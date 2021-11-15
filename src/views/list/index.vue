<template>
  <el-row type="flex" justify="center" align="middle">
    <el-col :span="8">
      <el-carousel :interval="5000" arrow="always" class="carousel">
        <el-carousel-item v-for="item in 4" :key="item">
          <h3>{{ item }}</h3>
        </el-carousel-item>
      </el-carousel>

      <el-card class="box-card" shadow="hover">
        <div slot="header" class="clearfix">
          <span>卡片名称</span>
        </div>
        <div v-for="o in 4" :key="o" class="text item">
          {{'列表内容 ' + o }}
        </div>
      </el-card>

      <el-card class="box-card" shadow="hover">
        <div slot="header" class="clearfix">
          <span>卡片名称</span>
        </div>
        <div v-for="o in 4" :key="o" class="text item">
          {{'列表内容 ' + o }}
        </div>
      </el-card>

      <el-card class="box-card" shadow="hover">
        <div slot="header" class="clearfix">
          <span>卡片名称</span>
        </div>
        <div v-for="o in 4" :key="o" class="text item">
          {{'列表内容 ' + o }}
        </div>
      </el-card>
    </el-col>
  </el-row>
</template>

<style scoped>

.carousel {
  text-align: center;
}

.box-card {
  margin-top: 10px;
  /*height: 240px;*/
}

.text {
  font-size: 14px;
}

.item {
  margin-bottom: 18px;
}

.el-carousel__item h3 {
  color: #475669;
  font-size: 18px;
  opacity: 0.75;
  line-height: 300px;
  margin: 0;
}

.el-carousel__item:nth-child(2n) {
  background-color: #99a9bf;
}

.el-carousel__item:nth-child(2n+1) {
  background-color: #d3dce6;
}

.time {
  font-size: 13px;
  color: #999;
}

.bottom {
  margin-top: 13px;
  line-height: 12px;
}

.button {
  padding: 0;
  float: right;
}

.image {
  width: 100%;
  display: block;
}

.clearfix:before,
.clearfix:after {
  display: table;
  content: "";
}

.clearfix:after {
  clear: both
}


</style>
