<template>
    <div @click="toRegister" class="button" style="position:absolute;left:0px;top:150px;width:100px;">
        <div style="position: absolute;right:22px;">注册</div>
    </div>
    <div @click="toLogin" class="button" style="position:absolute;left:300px;top:150px;width:100px;">
        <div style="position: absolute;right:22px;">登录</div>
    </div>
    <div @click="toArticlePublish" class="button" style="position:absolute;left:600px;top:150px;width:100px;">
        <div style="position: absolute;right:22px;">发布文章</div>
    </div>
    <div @click="toArticleHome" class="button" style="position:absolute;left:900px;top:150px;width:100px;">
        <div style="position: absolute;right:22px;">文章中心</div>
    </div>

</template>

<script setup>

import { useRouter } from 'vue-router'
const router = useRouter()

const toRegister = () => {
    router.push("/user/register")
}

const toLogin = () => {
    router.push("/user/login")
}

const toArticlePublish = () => {
    router.push("/article/publish")
}

const toArticleHome = () => {
    router.push("/article")
}

</script>

<style lang="scss" scoped>
.button {
    position: absolute;
    width: 80px;
    height: 40px;
    border-radius: 20px;
    background-color: #7B3DE0;
    line-height: 40px;
    font-size: 16px;
    color: white;
    cursor: pointer;
}
</style>
