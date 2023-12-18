<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="链接:" prop="url">
          <SelectImage v-model="formData.url" file-type="image"/>
       </el-form-item>
        <el-form-item label="图片格式:" prop="format">
          <el-input v-model="formData.format" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="宽:" prop="width">
          <el-input v-model.number="formData.width" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="高:" prop="height">
          <el-input v-model.number="formData.height" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="备注:" prop="comment">
          <el-input v-model="formData.comment" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createImage,
  updateImage,
  findImage
} from '@/api/image'

defineOptions({
    name: 'ImageForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
import SelectImage from '@/components/selectImage/selectImage.vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            url: "",
            format: '',
            width: 0,
            height: 0,
            comment: '',
        })
// 验证规则
const rule = reactive({
               url : [{
                   required: true,
                   message: '图片链接不能为空',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findImage({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reimage
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createImage(formData.value)
               break
             case 'update':
               res = await updateImage(formData.value)
               break
             default:
               res = await createImage(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
