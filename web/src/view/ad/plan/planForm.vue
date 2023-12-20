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
        <el-form-item label="状态:" prop="status">
          <el-select v-model="formData.status" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in adStatusOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
       </el-form-item>
        <el-form-item label="投放方式:" prop="mode">
          <el-select v-model="formData.mode" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in adModeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
       </el-form-item>
        <el-form-item label="时区:utc:" prop="timezone">
          <el-select v-model="formData.timezone" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in timezoneOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
       </el-form-item>
        <el-form-item label="开始时间:" prop="startAt">
          <el-date-picker v-model="formData.startAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="结束时间:" prop="endAt">
          <el-date-picker v-model="formData.endAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="总预算,千分:" prop="budgetTotal">
          <el-input v-model.number="formData.budgetTotal" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="每日预算,千分:" prop="budgetDaily">
          <el-input v-model.number="formData.budgetDaily" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="总曝光数:" prop="impTotal">
          <el-input v-model.number="formData.impTotal" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="每日曝光数:" prop="impDaily">
          <el-input v-model.number="formData.impDaily" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="曝光频制:" prop="impFrequency">
          <el-input v-model.number="formData.impFrequency" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="曝光频控周期:" prop="impFrequencyMinute">
          <el-input v-model.number="formData.impFrequencyMinute" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="点击频控:" prop="clkFrequency">
          <el-input v-model.number="formData.clkFrequency" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="点击频控周期:" prop="clkFrequencyMinute">
          <el-input v-model.number="formData.clkFrequencyMinute" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="最小点击率，单位0.1%:" prop="ctrMax">
          <el-input v-model.number="formData.ctrMax" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="最大点击率，单位0.1%:" prop="ctrMin">
          <el-input v-model.number="formData.ctrMin" :clearable="true" placeholder="请输入" />
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
  createPlan,
  updatePlan,
  findPlan
} from '@/api/plan'

defineOptions({
    name: 'PlanForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const adStatusOptions = ref([])
const adModeOptions = ref([])
const timezoneOptions = ref([])
const formData = ref({
            name: '',
            desc: '',
            status: undefined,
            mode: undefined,
            timezone: undefined,
            startAt: new Date(),
            endAt: new Date(),
            budgetTotal: 0,
            budgetDaily: 0,
            impTotal: 0,
            impDaily: 0,
            impFrequency: 0,
            impFrequencyMinute: 0,
            clkFrequency: 0,
            clkFrequencyMinute: 0,
            ctrMax: 0,
            ctrMin: 0,
        })
// 验证规则
const rule = reactive({
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               status : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               mode : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               timezone : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               startAt : [{
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
      const res = await findPlan({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.replan
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    adStatusOptions.value = await getDictFunc('adStatus')
    adModeOptions.value = await getDictFunc('adMode')
    timezoneOptions.value = await getDictFunc('timezone')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createPlan(formData.value)
               break
             case 'update':
               res = await updatePlan(formData.value)
               break
             default:
               res = await createPlan(formData.value)
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
