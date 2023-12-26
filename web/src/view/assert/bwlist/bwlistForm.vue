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
        <el-form-item label="平台、渠道白名单:" prop="platformWhitelist">
          <el-input v-model="formData.platformWhitelist" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="平台、渠道黑名单:" prop="platformBlacklist">
          <el-input v-model="formData.platformBlacklist" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="广告位白名单:" prop="spotWhitelist">
          <el-input v-model="formData.spotWhitelist" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="广告位黑名单:" prop="spotBlacklist">
          <el-input v-model="formData.spotBlacklist" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="应用白名单:" prop="bundleWhitelist">
          <el-input v-model="formData.bundleWhitelist" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="应用黑名单:" prop="bundleBlacklist">
          <el-input v-model="formData.bundleBlacklist" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="网站域名白名单:" prop="siteWhitelist">
          <el-input v-model="formData.siteWhitelist" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="网站域名黑名单:" prop="siteBlacklist">
          <el-input v-model="formData.siteBlacklist" :clearable="true" placeholder="请输入" />
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
  createBlackWhiteList,
  updateBlackWhiteList,
  findBlackWhiteList
} from '@/api/bwlist'

defineOptions({
    name: 'BlackWhiteListForm'
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
            name: '',
            desc: '',
            platformWhitelist: '',
            platformBlacklist: '',
            spotWhitelist: '',
            spotBlacklist: '',
            bundleWhitelist: '',
            bundleBlacklist: '',
            siteWhitelist: '',
            siteBlacklist: '',
        })
// 验证规则
const rule = reactive({
               name : [{
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
      const res = await findBlackWhiteList({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rebwlist
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
               res = await createBlackWhiteList(formData.value)
               break
             case 'update':
               res = await updateBlackWhiteList(formData.value)
               break
             default:
               res = await createBlackWhiteList(formData.value)
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
