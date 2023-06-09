<template>
  <div style="border: 1px solid #ccc;">
    <Toolbar :editor="editorRef" :defaultConfig="toolbarConfig" :mode="mode" style="border-bottom: 1px solid #ccc" />
    <Editor :defaultConfig="editorConfig" :mode="mode" v-model="valueHtml" style="height: 630px; overflow-y: hidden"
      @onCreated="handleCreated" @onChange="handleChange" />
  </div>
</template>
  
<script setup>
import '@wangeditor/editor/dist/css/style.css' // 引入 css
import { ref, inject, onMounted, onBeforeUnmount, shallowRef } from 'vue'
import { Editor, Toolbar } from '@wangeditor/editor-for-vue'
import { showMessage } from '../components/Message.js'

// 服务端地址
const serverUrl = inject("serverUrl")

// 编辑器实例，必须用 shallowRef
const editorRef = shallowRef();
// 模式
const mode = ref("default")
// 内容HTML
const valueHtml = ref("")
//菜单栏配置
const toolbarConfig = { excludeKeys: ['group-video', 'fullScreen'] };
// 编辑器配置
const editorConfig = { placeholder: '请输入内容...', MENU_CONF: {} };
// 上传图片
editorConfig.MENU_CONF = {}
editorConfig.MENU_CONF['uploadImage'] = {
  maxFileSize: 1024 * 1024 * 1024, // 取消原先的文件限制
  // 自定义文件限制函数
  onBeforeUpload(file) {
    console.log(Object.keys(file))
    if (file[Object.keys(file)[Object.keys(file).length-1]].size >= 2 * 1024 * 1024) {
      showMessage("图片不能超过2M", 'warning')
      return false
    }
    else {
      return file
    }
  },
  server: serverUrl + "/image/upload/rich_editor_upload",
}
// 插入图片
editorConfig.MENU_CONF['insertImage'] = {
  parseImageSrc: (src) => {
    if (src.indexOf("http") != 0) {
      return `${serverUrl}${src}`
    }
    return src
  }
}
// 定义属性进行双向绑定
const props = defineProps({
  modelValue: {
    type: String,
    default: ""
  }
})
// 定义抛出事件
const emit = defineEmits(["update:model-value"])
// 模拟 ajax 异步获取内容
onMounted(() => {
  setTimeout(() => {
    valueHtml.value = props.modelValue
    initFinished = true
  }, 10)
})
let initFinished = false
// 组件销毁时，也及时销毁编辑器
onBeforeUnmount(() => {
  const editor = editorRef.value
  if (editor == null) return
  editor.destroy()
})

// 编辑器回调函数
const handleCreated = (editor) => {
  editorRef.value = editor // 记录 editor 实例，重要！
}
const handleChange = (editor) => {
  if (initFinished) {
    emit("update:model-value", valueHtml.value) // 输入时往外抛
  }
};

</script>
  
<style lang="scss" scoped></style>
  