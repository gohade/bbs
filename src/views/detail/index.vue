<template>
  <el-row type="flex" justify="center" align="middle">
    <el-col :span="8">
      <el-card v-if="question" class="box-card" shadow="never">
        <div slot="header" class="clearfix">
          <span>{{question.title}} <span class="header_name" style="margin-right: 5px; float: right;"> <span @click="gotoQuestionEdit">修改</span> ｜ <span>删除</span> </span> </span>
        </div>
        <div>
          <viewer ref="questionViewer" :options="questionViewerOptions" :initialValue="question.context" />
        </div>
      </el-card>
      <el-divider content-position="left">所有回答</el-divider>
      <el-card v-for="answer in question.answers" style="margin-top: 5px; " class="box-card" shadow="hover">
        <div slot="header" class="clearfix">
          <span>{{answer.author.user_name}} | {{answer.created_at | formatDate}} <span class="header_name" style="margin-right: 5px; float: right;">删除</span></span>
        </div>
        <div>
          <viewer ref="answerViewer" :initialValue="answer.content" />
        </div>
      </el-card>
      <el-divider content-position="left">我来回答</el-divider>
      <el-card class="box-card" shadow="never">
        <el-form ref="form" :model="question">
          <el-form-item >
            <editor :options="editorOptions"
                    :initialValue = "answerContext"
                    height="200px"
                    initialEditType="wysiwyg"
                    ref="toastuiEditor"
                    previewStyle="vertical" />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="postAnswer">提交</el-button>
          </el-form-item>
        </el-form>
      </el-card>
    </el-col>
  </el-row>
</template>

<script>
import '@toast-ui/editor/dist/toastui-editor-viewer.css';
import '@toast-ui/editor/dist/toastui-editor.css';

import { Viewer} from '@toast-ui/vue-editor';
import { Editor } from '@toast-ui/vue-editor';
import request from "../../utils/request";

export default {
  components: {
    viewer: Viewer,
    editor: Editor,
  },
  data() {
    return {
      id: 0,
      question: null,
      questionViewerOptions: {
        usageStatistics: true,
      },
      answerContext: '',
      editorOptions: {
        minHeight: '200px',
        language: 'en-US',
        useCommandShortcut: true,
        usageStatistics: true,
        hideModeSwitch: true,
        toolbarItems: [
          ['heading', 'bold', 'italic', 'strike'],
          ['hr', 'quote'],
          ['ul', 'ol', 'task', 'indent', 'outdent'],
          ['table', 'image', 'link'],
          ['code', 'codeblock'],
          ['scrollSync'],
        ]
      }
    };
  },
  created() {
    if (this.$route.query.id) {
      let id = parseInt(this.$route.query.id)
      this.getDetail(id);
    }
  },
  methods: {
    getDetail: function (id) {
      const that = this
      this.id = id
      request({
        url: "/question/detail",
        method: 'GET',
        params: {
          "id": id
        }
      }).then(function (response) {
        that.question = response.data;
      })
    },
    postAnswer: function () {
      let html = this.$refs.toastuiEditor.$data.editor.getHTML()
      this.answerContext = html
      const that = this
      request({
        method: 'POST',
        url: "/answer/create",
        data: {
          "question_id": that.id,
          "context": that.answerContext,
        },
      }).then(function () {
        that.$router.go(0)
      })
    },
    gotoQuestionEdit: function () {
      this.$router.push({path: '/edit', query:{'id': this.id}})
    },
  }
};
</script>

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

.header_name {
  float: right;
  font-size: 13px;
  font-weight: 400;
  margin: 0 15px 0 0;
  line-height: 34px;
  background-color: transparent;
  color: #486e9b;
  text-decoration: none;
}
.header_name > a {
  background-color: transparent;
  color: #486e9b;
  text-decoration: none;
}
</style>
