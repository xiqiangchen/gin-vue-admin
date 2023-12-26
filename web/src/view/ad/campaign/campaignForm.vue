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
          <el-switch v-model="formData.status" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="开始时间:" prop="startAt">
          <el-date-picker v-model="formData.startAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="结束时间:" prop="endAt">
          <el-date-picker v-model="formData.endAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="总预算,元:" prop="budgetTotal">
          <el-input v-model.number="formData.budgetTotal" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="每日预算,元:" prop="budgetDaily">
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
        <el-form-item label="最小点击率，单位%:" prop="ctrMax">
          <el-input-number v-model="formData.ctrMax" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="最大点击率，单位%:" prop="ctrMin">
          <el-input-number v-model="formData.ctrMin" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="投放时间段，10位表示的二进制:" prop="hours">
          <el-input v-model.number="formData.hours" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="定向包id:" prop="targetId">
          <el-input v-model.number="formData.targetId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="黑白名单id:" prop="bwlistId">
          <el-input v-model.number="formData.bwlistId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="出价策略id:" prop="policyId">
          <el-input v-model.number="formData.policyId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="出价方式:" prop="bidMethod">
          <el-select v-model="formData.bidMethod" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in bidMethodOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
       </el-form-item>
        <el-form-item label="出价策略:" prop="bidPrice">
          <el-input-number v-model="formData.bidPrice" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="出价模式:" prop="bidMode">
          <el-select v-model="formData.bidMode" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in bidModeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
       </el-form-item>
        <el-form-item label="品牌名称:" prop="brand">
          <el-input v-model="formData.brand" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="允许虚拟:" prop="allowVirtually">
          <el-switch v-model="formData.allowVirtually" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="创意方式:" prop="creativeMode">
          <el-input v-model.number="formData.creativeMode" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="曝光监测:" prop="impTrackUrl">
          <el-input v-model="formData.impTrackUrl" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="点击监测:" prop="clickTrackUrl">
          <el-input v-model="formData.clickTrackUrl" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="落地页:" prop="h5">
          <el-input v-model="formData.h5" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="deeplink字段:" prop="deeplink">
          <el-input v-model="formData.deeplink" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="universalLink字段:" prop="universalLink">
          <el-input v-model="formData.universalLink" :clearable="true" placeholder="请输入" />
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
  createCampaign,
  updateCampaign,
  findCampaign
} from '@/api/campaign'

defineOptions({
    name: 'CampaignForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const bidMethodOptions = ref([])
const bidModeOptions = ref([])
const formData = ref({
            name: '',
            desc: '',
            status: false,
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
            hours: 0,
            targetId: 0,
            bwlistId: 0,
            policyId: 0,
            bidMethod: undefined,
            bidPrice: 0,
            bidMode: undefined,
            brand: '',
            allowVirtually: false,
            creativeMode: 0,
            impTrackUrl: '',
            clickTrackUrl: '',
            h5: '',
            deeplink: '',
            universalLink: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCampaign({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.recampaign
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    bidMethodOptions.value = await getDictFunc('bidMethod')
    bidModeOptions.value = await getDictFunc('bidMode')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createCampaign(formData.value)
               break
             case 'update':
               res = await updateCampaign(formData.value)
               break
             default:
               res = await createCampaign(formData.value)
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
