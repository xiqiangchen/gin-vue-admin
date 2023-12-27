<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="planId字段:" prop="planId">
          <el-input v-model.number="formData.planId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="campaignId字段:" prop="campaignId">
          <el-input v-model.number="formData.campaignId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="materialId字段:" prop="materialId">
          <el-input v-model.number="formData.materialId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="title字段:" prop="title">
          <el-input v-model="formData.title" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="desc字段:" prop="desc">
          <el-input v-model="formData.desc" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="button字段:" prop="button">
          <el-input v-model="formData.button" :clearable="true" placeholder="请输入" />
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
  createCreative,
  updateCreative,
  findCreative
} from '@/api/creative'

defineOptions({
    name: 'CreativeForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            planId: 0,
            campaignId: 0,
            materialId: 0,
            title: '',
            desc: '',
            button: '',
        })
// 验证规则
const rule = reactive({
               planId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               campaignId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               materialId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCreative({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.recreative
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
               res = await createCreative(formData.value)
               break
             case 'update':
               res = await updateCreative(formData.value)
               break
             default:
               res = await createCreative(formData.value)
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
