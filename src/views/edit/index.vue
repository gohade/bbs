<template>
  <el-row type="flex" justify="center" align="middle">
    <el-col :span="8">
      <el-card class="box-card" shadow="never">
        <el-form ref="form" :model="question" label-width="80px">
          <el-form-item label="问题标题">
            <el-input v-model="question.title">{{question.title}}</el-input>
          </el-form-item>
          <el-form-item label="问题描述" v-if="question.context">
            <editor :options="editorOptions"
                    :initialValue="question.context"
                    height="500px"
                    initialEditType="wysiwyg"
                    previewStyle="vertical" />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="onSubmit">更新</el-button>
          </el-form-item>
        </el-form>
      </el-card>
    </el-col>
  </el-row>
</template>

<script>
import '@toast-ui/editor/dist/toastui-editor.css';

import { Editor } from '@toast-ui/vue-editor';
import request from "../../utils/request";

export default {
  components: {
    editor: Editor
  },
  created() {
    if (this.$route.query.id) {
      let id = parseInt(this.$route.query.id)
      this.getDetail(id);
    }
  },
  data() {
    return {
      question: {
        title: '',
        context: ''
      },
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
    onSubmit: function () {

    }
  }
}
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
</style>
