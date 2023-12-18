<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="描述:" prop="desc">
          <el-input v-model="formData.desc" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="广告形式:" prop="ad_format">
          <el-select v-model="formData.ad_format" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in adFormatOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
       </el-form-item>
        <el-form-item label="设备类型:" prop="device_type">
          <el-select v-model="formData.device_type" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in deviceTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
       </el-form-item>
        <el-form-item label="操作系统:" prop="os">
          <el-select v-model="formData.os" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in osOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
       </el-form-item>
        <el-form-item label="定向类型:" prop="target_type">
          <el-select v-model="formData.target_type" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in targetTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
       </el-form-item>
        <el-form-item label="行政区域:" prop="region">
          <el-input v-model="formData.region" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="性别:" prop="gender">
          <el-select v-model="formData.gender" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in genderOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
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
  createTarget,
  updateTarget,
  findTarget
} from '@/api/target'

defineOptions({
    name: 'TargetForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const targetTypeOptions = ref([])
const genderOptions = ref([])
const adFormatOptions = ref([])
const deviceTypeOptions = ref([])
const osOptions = ref([])
const formData = ref({
            name: '',
            desc: '',
            ad_format: undefined,
            device_type: undefined,
            os: undefined,
            target_type: undefined,
            region: '',
            gender: undefined,
        })
// 验证规则
const rule = reactive({
               name : [{
                   required: true,
                   message: '定向包名称必填',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findTarget({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.retarget
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    targetTypeOptions.value = await getDictFunc('targetType')
    genderOptions.value = await getDictFunc('gender')
    adFormatOptions.value = await getDictFunc('adFormat')
    deviceTypeOptions.value = await getDictFunc('deviceType')
    osOptions.value = await getDictFunc('os')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createTarget(formData.value)
               break
             case 'update':
               res = await updateTarget(formData.value)
               break
             default:
               res = await createTarget(formData.value)
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
