<template>
  <el-row type="flex" justify="center" align="middle">
    <el-col :span="8">
      <div class="infinite-list-wrapper" style="overflow:auto">
        <ul
            class="list"
            v-infinite-scroll="load"
            infinite-scroll-disabled="disabled">
            <el-card v-for="question in questions" class="box-card" shadow="hover">
              <div slot="header" class="clearfix">
          <span>{{question.title}}</span>
        </div>
              <div class="text item">
                {{question.context}}
              </div>
              <div class="bottom clearfix">
                  <time class="time">{{question.created_at}} ｜ {{question.author.user_name}}  | {{question.answer_num}} 回答</time>
                  <el-button type="text" class="button" @click="gotoDetail(question.id)">去看看</el-button>
              </div>
            </el-card>
        </ul>
        <p v-if="loading" class="loading_tips">加载中...</p>
        <p v-if="disabled" class="loading_tips">没有更多了</p>
      </div>
    </el-col>
  </el-row>
</template>

<script>
import request from "../../utils/request";

export default {
  data () {
    return {
      count: 10,
      start: 0,
      size: 10,
      questions: [],
      loading: false,
      noMore: false
    }
  },
  created() {
    this.getQuestions();
  },
  computed: {
    disabled () {
      return this.loading || this.noMore
    }
  },
  methods: {
    load () {
      if (this.noMore === true) {
        return
      }
      this.loading = true
      setTimeout(() => {
        this.loading = false
        this.getQuestions()
      }, 2000)
    },
    getQuestions() {
      const that = this;
      request({
        url: '/question/list',
        method: 'get',
        params: {
          start: this.start,
          size: this.size,
        }
      }).then(function (response) {
        const questions = response.data
        if (questions === null || questions.length === 0) {
          that.noMore = true
        }
        that.questions = that.questions.concat(questions)
        that.start = that.start + questions.length
      })
      this.loading = false;
    },
    gotoDetail(id) {
      // go to detail page
      this.$router.push({path: '/detail', query:{'id': id}})
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
