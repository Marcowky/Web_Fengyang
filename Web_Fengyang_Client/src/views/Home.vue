<template>
    <!-- <div class="background">
        <img src="../assets/image/rectangle1.png" class="rectangle1" />
        <img src="../assets/image/rectangle2.png" class="rectangle2" />
        <img src="../assets/image/rectangle3.png" class="rectangle3" />
        <img src="../assets/image/rectangle4.png" class="rectangle4" />
        <img src="../assets/image/person.png" class="person" />
    </div> -->
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
    <!-- <div class="board">
        <div>
            <div @click="toRegister" class="button2">
                <div style="position: absolute;right:22px;">注册</div>
            </div>
            <div class="button1">
                <div style="position:absolute;left:22px;">登录</div>
            </div>
        </div>
        <n-form ref="formRef" :rules="rules" :model="user">
            <n-form-item path="phoneNumber" style="position:absolute;left:70px;top:150px;width:350px;">
                <n-input v-model:value="user.phoneNumber" size="large" round placeholder="手机号" />
            </n-form-item>
            <n-form-item path="password" style="position:absolute;left:70px;top:230px;width:350px;">
                <n-input v-model:value="user.password" size="large" round type="password" placeholder="密码" />
            </n-form-item>
        </n-form>
        <n-checkbox v-model:checked="user.rember" label="记住密码" style="position:absolute;left:70px;top:330px;" />
        <div @click="submit" class="button3">
            <div style="left: auto;right: auto;text-align: center;">登录</div>
        </div>
    </div> -->
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
