<template xmlns="http://www.w3.org/1999/html">
  <el-dialog model-value="showModal" :title="dialogTitle" width="30%" show-close=false @closed='handleClose'>
    <el-form ref="formRef" :rules="rules" :model="hotel" label-width="100px"  style="min-width: 250px">

<!--      <el-form-item label="酒店图片地址" prop="image">-->
<!--        <el-input v-model="hotel.image" />-->
<!--      </el-form-item>-->
      <!-- 无封面时 -->
      <div v-if="!newHeadImage" style="width: 80%; margin: auto;">
        <el-upload drag :before-upload="imageCheck" :http-request="customRequest" multiple>
          <el-icon class="el-icon--upload"><upload-filled /></el-icon>
          <div class="el-upload__text">
            拖动文件 或 <em>点击上传</em>
          </div>
          <template #tip>
            <div class="el-upload__tip">
              仅支持png/jpg/jpeg格式的图片,照片不为空
            </div>
          </template>
        </el-upload>
      </div>
      <!-- 有封面时 -->
      <div v-else style="width: 80%; margin: auto;">
        <el-image :src="serverUrl + hotel.image"></el-image>
        <el-button class="delete-button" size="large" @click="deleteImage" type="danger" :icon="Delete" circle />
      </div>
      <div class="content-text">
      <el-form-item label="酒店名" prop="name"  >
        <el-input v-model="hotel.name" class="item_style"/>
      </el-form-item>
      <el-form-item label="电话" prop="telephone">
        <el-input v-model="hotel.telephone" class="item_style"/>
      </el-form-item>
      <el-form-item label="价格"  prop="price">
        <el-input v-model="hotel.price" class="item_style"/>
      </el-form-item>
      <el-form-item label="网站" prop="website">
        <el-input v-model="hotel.website" class="item_style"/>
      </el-form-item>
      <el-form-item label="东经" prop="longitude"  >
        <el-input v-model="hotel.longitude" class="item_style"/>
      </el-form-item>
      <el-form-item label="北纬" prop="latitude"  >
        <el-input v-model="hotel.latitude" class="item_style"/>
      </el-form-item>
      <el-form-item label="地址" prop="placeAddress">
        <el-input v-model="hotel.placeAddress" class="item_style"/>
      </el-form-item>
      </div>
    </el-form>
    <div class="button_box">
      <el-button @click="submitForm(formRef)" class="button3" type="primary">确认</el-button>
    <el-button type="danger" @click="handleClose">取消</el-button>
    </div>
  </el-dialog>
</template>

<script setup>
import { ref, reactive, inject, watch } from 'vue';
import {ElMessage} from "element-plus";
import { showMessage } from '../../../components/Message';
import {articlePost, articleUpdate} from '../../../api/article.js'
import {imageCheck, imageDelete, imageUpload} from "../../../api/image.js";
import {Delete, UploadFilled} from "@element-plus/icons-vue";
const axios = inject("axios")
const serverUrl = inject("serverUrl")
const props = defineProps({
  dialogTitle: {
    type: String,
    default: '',
    required: true
  },
  dialogTableValue: {
    tyoe: Object,
    default: () => {}
  },
})
const newHeadImage= ref(false)
const formRef = ref(null)
const raw_hotel = ref(null)
const hotel = ref({
  id:"",
  name:"",
  image:"",
  telephone:"",
  price: "",
  website: "",
  center: "",
  longitude:"",
  latitude:"",
  placeAddress:""
})
const gethotaldetail =(rowtitle,content,image,id)=> {
  var priceMatch = content.match(/price: (.*?)##/);
  var websiteMatch = content.match(/website: "(.*?)"##/);
  var centerMatch = content.match(/center: \[(.*?)\]##/);
  var telephoneMatch = content.match(/telephone: (.*?)##/);
  var placeAddressMatch = content.match(/location: "(.*?)"/);
  // 获取匹配结果中的值
  hotel.value.price = priceMatch ? parseInt(priceMatch[1]) : null
  hotel.value.website = websiteMatch ? websiteMatch[1] : null
  hotel.value.center = centerMatch ? centerMatch[1].split(',').map(Number) : null // 分割数字并将其转换为数字类型
  hotel.value.telephone = telephoneMatch ? parseInt(telephoneMatch[1]) : null
  hotel.value.name = rowtitle
  hotel.value.placeAddress = placeAddressMatch ? placeAddressMatch[1] : null
  hotel.value.image=image
  hotel.value.id=id
  hotel.value.longitude=hotel.value.center[0]
  hotel.value.latitude=hotel.value.center[1]
  console.log(hotel.value)
}

watch(() => props.dialogTableValue, () => {
  if (props.dialogTableValue.test != null) {
    // console.log("添加酒店");
  } else {
    raw_hotel.value = props.dialogTableValue;
    newHeadImage.value=true;
    gethotaldetail(raw_hotel.value.title, raw_hotel.value.content, raw_hotel.value.head_image, raw_hotel.value.id);
  }
}, { deep: true, immediate: true });


const emits = defineEmits(['update:modelValue', 'updateHotelsList'])
// 如果是取消,南无传递给外部不用刷新表单
const handleClose = () => {
  emits('update:modelValue',false)
}

// 表单验证规则
const rules = reactive({
  name: [
    { required: true, message: "请输入酒店名", trigger: "blur" },
    { min: 3, max: 5, message: "酒店名称在 3 到 5 个字符", trigger: "blur" },
  ],
  telephone: [
    { required: true, message: "请输入酒店电话", trigger: "blur" },
    { min: 11, max: 11, message: "手机号为 11 位", trigger: "blur" },
  ],
  price: [
    { required: true, message: "请输入价格", trigger: "blur" },
  ],
})

const customRequest = async (file) => {
  imageUpload(file).then(result => {
    if (result != null) {
      hotel.value.image = result.data.data.filePath
      newHeadImage.value = true
    }
  })
}

const deleteImage = async () => {
  console.log("delete"+hotel.value.image)
  imageDelete(hotel.value.image).then(result => {
    if (result == null) {
      hotel.value.image = ""
      newHeadImage.value = false
    }
  })
}
const submitForm = async (formEl) => {
  const new_title=`${hotel.value.name}`
  const new_content=`price: ${hotel.value.price}##telephone: ${hotel.value.telephone}##website: "${hotel.value.website}"##center: [${hotel.value.longitude}, ${hotel.value.latitude}]##location: "${hotel.value.placeAddress}"`
  if(newHeadImage.value==false)
  {
    // 表单验证不通过，可以进行错误处理
    showMessage('照片验证不能为空，请添加图片', 'error')
    return;
  }
  // 验证表单
  const valid = await formEl.validate()
  if (valid) {
    // 表单验证通过，执行 POST 或 GET 请求
    // 剩余代码省略
  } else {
    // 表单验证不通过，可以进行错误处理
    showMessage('表单验证不通过，请检查输入', 'error')
    return;
  }
  // 更改酒店
  if(hotel.value.id!=="")
  {
    let updateArticle = {
      articleType: 'hotelArticle',
      id: hotel.value.id,
      title: new_title,
      content: new_content,
      headImage: hotel.value.image,
      categoryId: '1'
    }
    await update(updateArticle)
  }
  //添加酒店
  else {
    let addArticle = {
      articleType: 'hotelArticle',
      title: new_title,
      content: new_content,
      headImage: hotel.value.image,
      categoryId: 1
    }
    await submit(addArticle)


    // let res = await axios.post("/article/create", {
    //   title: new_title,
    //   content: new_content,
    //   head_image: hotel.value.image,
    //   article_type: "hotelArticle",
    //   category_id: 1,
    // })
    // if (res.data.code == 200) {
    //   ElMessage({
    //     message: res.data.msg,
    //     type: 'success',
    //     offset: 80
    //   })
    //   emits('updateHotelsList')
    //   handleClose()
    // } else {
    //   ElMessage({
    //     message: res.data.msg,
    //     type: 'error',
    //     offset: 80
    //   })
    // }
  }

}

const submit = async (addArticle) => {
  articlePost(addArticle).then(result => {
    if (result == null) {
      // router.go(-1)
      emits('updateHotelsList')
      handleClose()
    }
  })
}

const update = async (updateArticle) => {
  articleUpdate(updateArticle).then(result => {
    if (result == null) {
      // router.go(-1)
      emits('updateHotelsList')
      handleClose()
    }
  })
}
</script>

<style scoped>
.delete-button {
  position: absolute;
  right: 0px;
  bottom: 0px;
  box-shadow: 0 0 6px rgba(50, 50, 50, 0.26);
}
.content-text{
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100%;
}
.item_style{
  width: 370px;
  padding-right: 50px ;
}
.button_box{
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>