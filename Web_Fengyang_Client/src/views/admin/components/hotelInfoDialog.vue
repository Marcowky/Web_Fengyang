<template>
  <el-dialog model-value="showModal" :title="dialogTitle" width="30%" show-close=false @closed='handleClose'>
    <el-form ref="formRef" :rules="rules" :model="hotel" label-width="100px"  style="min-width: 250px">
      <el-form-item label="酒店名" prop="name"  >
        <el-input v-model="hotel.name" />
      </el-form-item>
      <el-form-item label="酒店图片地址" prop="image">
        <el-input v-model="hotel.image" />
      </el-form-item>
      <el-form-item label="电话" prop="telephone">
        <el-input v-model="hotel.telephone" />
      </el-form-item>
      <el-form-item label="价格"  prop="price">
        <el-input v-model="hotel.price" />
      </el-form-item>
      <el-form-item label="网站" prop="website">
        <el-input v-model="hotel.website" />
      </el-form-item>
      <el-form-item label="经纬度" prop="center"  >
        <el-input v-model="hotel.center[0]" />
        <el-input v-model="hotel.center[1]" />
      </el-form-item>
      <el-form-item label="地址" prop="placeAddress">
        <el-input v-model="hotel.placeAddress" />
      </el-form-item>
    </el-form>

    <el-button @click="submitForm(formRef)" class="button3" type="primary"
               style="left: auto; right: auto; text-align: center; margin-top: 10px;">
      确认
    </el-button>
  </el-dialog>
</template>

<script setup>
import { ref, reactive, inject, watch } from 'vue';
import {ElMessage} from "element-plus";
const axios = inject("axios")
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
  placeAddress:""
})
const gethotaldetail =(rowtitle,content,image,id)=> {
  var priceMatch = content.match(/price: (.*?)(\n|$)/)
  var websiteMatch = content.match(/website: "(.*?)"(\n|$)/)
  var centerMatch = content.match(/center: \[(.*?)\](\n|$)/)
  var telephoneMatch = content.match(/telephone: (.*?)(\n|$)/)
  var titleMatch = rowtitle.match(/(.*) "/)
  var placeAddressMatch = rowtitle.match(/\"([^"]+)\"/)
  // 获取匹配结果中的值
  hotel.value.price = priceMatch ? parseInt(priceMatch[1]) : null
  hotel.value.website = websiteMatch ? websiteMatch[1] : null
  hotel.value.center = centerMatch ? centerMatch[1].split(',').map(Number) : null // 分割数字并将其转换为数字类型
  hotel.value.telephone = telephoneMatch ? parseInt(telephoneMatch[1]) : null
  hotel.value.name = titleMatch ? titleMatch[1] : null
  hotel.value.placeAddress = placeAddressMatch ? placeAddressMatch[1] : null
  hotel.value.image=image
  hotel.value.id=id
  console.log(hotel.value)
}

watch(() => props.dialogTableValue, () => {
  if (props.dialogTableValue.test != null) {
    // console.log("添加酒店");
  } else {
    raw_hotel.value = props.dialogTableValue;
    console.log(raw_hotel.value);
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

const submitForm = async (formEl) => {
  const new_title=`${hotel.value.name} "${hotel.value.placeAddress}"`
  const new_content=`price: ${hotel.value.price}\ntelephone: ${hotel.value.telephone}\nwebsite: "${hotel.value.website}"\ncenter: [${hotel.value.center[0]}, ${hotel.value.center[1]}]`
  // 验证表单
  const valid = await formEl.validate()
  if (valid) {
    // 表单验证通过，执行 POST 或 GET 请求
    // 剩余代码省略
  } else {
    // 表单验证不通过，可以进行错误处理
    ElMessage({
      message: '表单验证不通过，请检查输入',
      type: 'error',
      offset: 80
    })
  }
  // 添加酒店
  if(hotel.value.id!=="")
  {
    let res = await axios.put(`article/update?articleType=hotelarticle&id=${hotel.value.id}`, {
      title: new_title,
      content: new_content,
      head_image: hotel.value.image,
      article_type: "hotelarticle",
      category_id: 1,
    })
    if (res.data.code == 200) {
      ElMessage({
        message: res.data.msg,
        type: 'success',
        offset: 80
      })
      emits('updateHotelsList')
      handleClose()
    } else {
      ElMessage({
        message: res.data.msg,
        type: 'error',
        offset: 80
      })
    }
  }
  //更改酒店
  else {
    let res = await axios.post("/article/create", {
      title: new_title,
      content: new_content,
      head_image: hotel.value.image,
      article_type: "hotelarticle",
      category_id: 1,
    })
    if (res.data.code == 200) {
      ElMessage({
        message: res.data.msg,
        type: 'success',
        offset: 80
      })
      emits('updateHotelsList')
      handleClose()
    } else {
      ElMessage({
        message: res.data.msg,
        type: 'error',
        offset: 80
      })
    }
  }

}
</script>

<style scoped>

</style>