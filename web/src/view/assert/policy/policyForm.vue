<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="平台渠道:" prop="platform">
          <el-input v-model.number="formData.platform" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="广告位:" prop="spot">
          <el-input v-model.number="formData.spot" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="发布者:" prop="publisher">
          <el-input v-model.number="formData.publisher" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="bundle:" prop="bundle">
          <el-input v-model.number="formData.bundle" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="广告形式:" prop="format">
          <el-input v-model.number="formData.format" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="是否有设备id信息:" prop="identity">
          <el-switch v-model="formData.identity" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="操作系统:" prop="os">
          <el-input v-model.number="formData.os" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="地区:" prop="region">
          <el-input v-model.number="formData.region" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="平均出价:" prop="price">
          <el-input-number v-model="formData.price" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="浮动范围(%):" prop="scope">
          <el-input v-model.number="formData.scope" :clearable="true" placeholder="请输入" />
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
  createPolicy,
  updatePolicy,
  findPolicy
} from '@/api/policy'

defineOptions({
    name: 'PolicyForm'
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
            platform: 0,
            spot: 0,
            publisher: 0,
            bundle: 0,
            format: 0,
            identity: 0,
            os: 0,
            region: 0,
            price: 0,
            scope: 0,
        })
// 验证规则
const rule = reactive({
               price : [{
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
      const res = await findPolicy({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.repolicy
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
               res = await createPolicy(formData.value)
               break
             case 'update':
               res = await updatePolicy(formData.value)
               break
             default:
               res = await createPolicy(formData.value)
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
