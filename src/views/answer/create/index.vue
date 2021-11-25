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
        <el-form ref="form" :model="question">
          <el-form-item >
            <el-input v-model="question.title"></el-input>
          </el-form-item>
          <el-form-item >
            <editor :options="editorOptions"
                    height="500px"
                    initialEditType="wysiwyg"
                    previewStyle="vertical" />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="onSubmit">立即提问</el-button>
            <el-button>取消</el-button>
          </el-form-item>
        </el-form>
      </el-card>
    </el-col>
  </el-row>
</template>

<script>
import '@toast-ui/editor/dist/toastui-editor.css';

import { Editor } from '@toast-ui/vue-editor';

export default {
  components: {
    editor: Editor
  },
  data() {
    return {
      question: {
        title: '',
        content: ''
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
