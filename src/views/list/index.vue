<template>
  <el-row type="flex" justify="center" align="middle">
    <el-col :span="8">
      <div class="infinite-list-wrapper" style="overflow:auto">
        <ul
            class="list"
            v-infinite-scroll="load"
            infinite-scroll-disabled="disabled">
            <el-card v-for="i in count" class="box-card" shadow="hover">
              <div slot="header" class="clearfix">
          <span>问题标题{{i}}</span>
        </div>
              <div class="text item">
                这个是问题的具体内容，显示前200个字...
              </div>
              <div class="bottom clearfix">
                  <time class="time">2021-10-10 10:10:10 ｜ jianfengye  | 10 回答</time>
                  <el-button type="text" class="button">去看看</el-button>
              </div>
            </el-card>
        </ul>
        <p v-if="loading" class="loading_tips">加载中...</p>
        <p v-if="noMore" class="loading_tips">没有更多了</p>
      </div>
    </el-col>
  </el-row>
</template>

<script>
export default {
  data () {
    return {
      count:10,
      loading: false
    }
  },
  computed: {
    noMore () {
      return this.count >= 20
    },
    disabled () {
      return this.loading || this.noMore
    }
  },
  methods: {
    load () {
      this.loading = true
      setTimeout(() => {
        this.count += 2
        this.loading = false
      }, 2000)
    }
  }
}
</script>

<style scoped>

.loading_tips {
  text-align: center;
  font-size: 13px;
  color: #999;
}

.time {
  font-size: 13px;
  color: #999;
}

.bottom {
  margin-top: 13px;
  line-height: 12px;
}

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
