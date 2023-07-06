<template>
<el-card style="border-radius: 20px;margin-left: 120px;margin-right: 120px;margin-top: 60px">
  <el-button @click="router.go(-1)">返回</el-button>
  <el-container>
    <el-header  class="info-header">
      <h2 class="info-title">{{ articleInfo.title }}</h2>
      <div>
        <span>文章来源：未知</span>
        <el-divider direction="vertical"></el-divider>
        <span>作者：未知</span>
        <el-divider direction="vertical"></el-divider>
        <span>发布时间：{{ articleInfo.created_at }}</span>
      </div>
    </el-header>
    <el-main class="info-main">
      <div class="article-content" v-html="articleInfo.content"></div>
      <div style="height: 20px"></div>
    </el-main>
  </el-container>


</el-card>
</template>

<script setup>
import { ref, inject, onMounted } from 'vue'

// 导入路由
import { useRouter, useRoute } from 'vue-router'
const router = useRouter()
const route = useRoute()

// 网络请求
const axios = inject("axios")

// 定义变量
const articleInfo = ref({})
const categoryName = ref("")
const user = ref({})
const self = ref(false)
const activeIndex = ref('0')

// 挂载页面时触发
onMounted(() => {
  loadArticle()
})
const loadArticle = async () => {
  // 获取文章信息
  let resArticle = await axios.get(`article/detail?articleType=infoArticle&id=${route.query.id}`)
  if (resArticle.data.code === 200) {
    articleInfo.value = resArticle.data.data.article
  }
}
</script>

<style lang="scss" scoped>
img{
  height:inherit;
  width: 100%
}
.info-header{
  text-align: center;
  margin: 20px;
}
.info-title{
  position: relative;
  font-weight:600;
  text-align: center;
  margin: 20px;
  line-height: 1.6em;
}
.info-main{
  margin-left: 180px;
  margin-right: 180px;
  margin-top: 50px;
}
.info-align{
  position: relative;
  font-size: 18px;
  line-height: 2.0em;
}

</style>