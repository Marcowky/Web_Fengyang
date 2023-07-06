<template>
  <el-dialog model-value="showModal" :title="dialogTitle" width="25%" show-close=false @closed='handleClose' center>

    <el-form ref="formRef"  :model="carousel" label-width="100px" :label-position="labelPosition"
             style="min-width: 250px; padding-right: 40px;">
      <el-form-item label="ID" prop="ID" v-if="dialogTitle == editCarouselTitle">
        <el-input v-model="carousel.ID" disabled />
      </el-form-item>
      <el-form-item label="顺序" prop="order">
        <el-input v-model="carousel.order" />
      </el-form-item>
      <el-form-item label="图片">
        <!-- 无封面时 -->
        <div v-if="!newImage" style="width: 80%; margin: auto;">
          <el-upload drag :before-upload="imageCheck" :http-request="customRequest" multiple>
            <el-icon class="el-icon--upload"><upload-filled /></el-icon>
            <div class="el-upload__text">
              拖动文件 或 <em>点击上传</em>
            </div>
            <template #tip>
              <div class="el-upload__tip">
                仅支持png/jpg/jpeg格式的图片
              </div>
            </template>
          </el-upload>
        </div>
        <!-- 有封面时 -->
        <div v-else style="width: 80%; margin: auto;">
          <el-image :src="serverUrl + carousel.url"></el-image>
          <el-button class="delete-button" size="large" @click="deleteImage" type="danger" :icon="Delete" circle />
        </div>
      </el-form-item>
      <el-row>
        <el-col :span="11" style="padding-left: 15px;">
          <el-form-item label="轮播图类型" prop="category">
            <el-tag class="ml-2" type="success">{{ carousel.category }}</el-tag>
          </el-form-item>
        </el-col>
      </el-row>
    </el-form>

    <div class="button_area">
      <el-button @click="submitForm(formRef)" type="primary" style="margin-right: 30px;">
        确认
      </el-button>
      <el-button @click="handleClose" type="danger">
        取消
      </el-button>
    </div>

  </el-dialog>
</template>

<script setup>

import {ref, reactive, watch, inject} from 'vue';
import { showMessage } from '../../../components/Message';
import {axios} from "../../../main.js";
import { UploadFilled, Delete } from '@element-plus/icons-vue'
import {
  carouselRegister,
  carouselUpdate,
  carouselUpload,
  imageCheck,
  imageDelete,
} from "../../../api/image.js";

const props = defineProps({
  dialogTitle: {
    type: String,
    default: '',
    required: true
  },
  dialogTableValue: {
    type: Object,
    default: () => { }
  },
  dialogCarouselType: {
    type: String,
    default: '',
    required: true
  }
})
const serverUrl = inject("serverUrl") // 网络请求
const editCarouselTitle = "编辑轮播图"
const addCarouselTitle = "添加轮播图"
const formRef = ref(null)
const carousel = ref({
  ID: "",
  order: "",
  category: "",
  url: ""
})

const newImage = ref(false)

watch(() => props.dialogTableValue, () => (
        carousel.value.ID = props.dialogTableValue.ID,
        carousel.value.order = props.dialogTableValue.order,
        carousel.value.category = props.dialogTableValue.category,
        carousel.value.url = props.dialogTableValue.url,
        newImage.value = props.dialogTableValue.ID ? true : false
), { deep: true, immediate: true })

watch(() => props.dialogCarouselType, () => (
    carousel.value.category = props.dialogCarouselType
), { deep: true, immediate: true })

const emits = defineEmits(['update:modelValue', 'updateCarouselList'])

const handleClose = () => {
  emits('update:modelValue', false)
}

// // 表单验证规则
// const rules = reactive({
//   userName: [
//     { required: true, message: "请输入轮播图名", trigger: "blur" },
//     { min: 3, max: 20, message: "轮播图名长度在 3 到 20 个字符", trigger: "blur" },
//   ],
//   phoneNumber: [
//     { required: true, message: "请输入手机号", trigger: "blur" },
//     { min: 11, max: 11, message: "手机号为 11 位", trigger: "blur" },
//   ],
//   password: [
//     { required: true, message: "请输入密码", trigger: "blur" },
//     { min: 6, max: 20, message: "密码长度在 6 到 20 个字符", trigger: "blur" },
//   ]
// })


// 提交表单
const submitForm = async (formEl) => {
  if (!formEl) return

  await formEl.validate((valid, fields) => {
    if (valid) {
      if (props.dialogTitle == addCarouselTitle) {
        register()
      }
      else if (props.dialogTitle == editCarouselTitle) {
        updateCarousels()
      }
    }
    // else {
    //   if (dialogType.value === '登录') {
    //     showMessage('提交失败', 'error')
    //   } else if (dialogType.value === '注册') {
    //     showMessage('注册失败', 'error')
    //   }
    // }
  })
}
// 添加轮播图
const register = async () => {
  let tempCarousel = ({
    ID: carousel.value.ID,
    order: carousel.value.order,
    category: carousel.value.category,
    url: carousel.value.url,

  })
  carouselRegister(tempCarousel).then(result => {
    if (result == null) {
      emits('updateCarouselList')
      handleClose()
    }
  })
}

const updateCarousels = async () => {
  carouselUpdate(carousel.value).then(result => {
    if (result == null) {
      emits('updateCarouselList')
      handleClose()
    }
  })
}


// 上传轮播图图片
const customRequest = async (file) => {

  carouselUpload(file).then(result => {
    if (result != null) {
      carousel.url = result.data.data.filePath
      newImage.value = true
    }
  })
}

// 删除图片
const deleteImage = async () => {

  imageDelete(carousel.url).then(result => {
    if (result == null) {
      carousel.url = ""
      newImage.value = false
    }
  })
}



</script>

<style scoped>
.button_area {
  display: flex;
  justify-content: center;
}
</style>
