<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAt">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期" :style="{ width: '185px' }" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期" :style="{ width: '185px' }" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
      </el-form-item>
        <el-form-item label="名称" prop="name">
         <el-input v-model="searchInfo.name" placeholder="搜索条件" :style="{ width: '80px' }" />

        </el-form-item>
           <el-form-item label="状态" prop="status">
            <el-select v-model="searchInfo.status" clearable placeholder="请选择" :style="{ width: '80px' }" @clear="()=>{searchInfo.status=undefined}">
              <el-option v-for="(item,key) in adStatusOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-popover v-model:visible="deleteVisible" :disabled="!multipleSelection.length" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button type="primary" link @click="deleteVisible = false">取消</el-button>
                <el-button type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
            </template>
            </el-popover>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="ID" prop="ID" width="80" />
        <el-table-column align="left" label="名称" prop="name" width="120">
          <template #default="scope">
            <a :href="'#' + campaignPath + '?pid=' + scope.row.ID" >{{ scope.row.name }}</a>
          </template>
        </el-table-column>
        <el-table-column align="left" label="描述" prop="desc" width="120" />
        <el-table-column align="left" label="状态" prop="status" width="120">
            <template #default="scope">
              <el-switch v-model="scope.row.status" @change="handleSwitchChange(scope.row)"></el-switch>
            </template>
        </el-table-column>
        <el-table-column align="left" label="投放方式" prop="mode" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.mode,adModeOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="时区" prop="timezone" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.timezone,timezoneOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="总预算" prop="budgetTotal" width="120" />
        <el-table-column align="left" label="每日预算" prop="budgetDaily" width="120" />
        <el-table-column align="left" label="总曝光数" prop="impTotal" width="120" />
        <el-table-column align="left" label="每日曝光数" prop="impDaily" width="120" />
        <el-table-column align="left" label="点击率" prop="ctrMax" width="120">
          <template #default="scope">[{{ (scope.row.ctrMax) }}%, {{ (scope.row.ctrMin) }}%]</template>
        </el-table-column> />
        <el-table-column align="left" label="操作" min-width="120">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                查看详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updatePlanFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="type==='create'?'添加':'修改'" destroy-on-close>
      <el-scrollbar height="500px">
          <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="160px">
            <el-form-item label="名称:"  prop="name" >
              <el-input v-model="formData.name" :clearable="true"  placeholder="请输入名称" style="width:80%" />
            </el-form-item>
            <el-form-item label="描述:"  prop="desc" >
              <el-input v-model="formData.desc" :clearable="true"  placeholder="请输入描述" style="width:80%" />
            </el-form-item>
            <el-form-item label="状态:"  prop="status" >
              <el-switch v-model="formData.status"></el-switch>
            </el-form-item>
            <el-form-item label="投放方式:"  prop="mode" >
              <el-select v-model="formData.mode" placeholder="请选择投放方式" style="width:30%" :clearable="true" >
                <el-option v-for="(item,key) in adModeOptions" :key="key" :label="item.label" :value="item.value" />
                
              </el-select>
            </el-form-item>
            <el-form-item label="时区"  prop="timezone" >
              <el-select v-model="formData.timezone" placeholder="请选择时区" style="width:30%" :clearable="true" >
                <el-option v-for="(item,key) in timezoneOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-row>
              <el-col :span="10" class="grid-cell">
                <el-form-item label="开始时间:"  prop="startAt" >
                  <el-date-picker v-model="formData.startAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
                </el-form-item>
              </el-col>
              <el-col :span="10" class="grid-cell">
                <el-form-item label="结束时间:"  prop="endAt" >
                  <el-date-picker v-model="formData.endAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
                </el-form-item>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="10" class="grid-cell">
                <el-form-item label="总预算(元):"  prop="budgetTotal" >
                  <el-input-number v-model.number="formData.budgetTotal" :clearable="true" placeholder="请输入总预算,元" />
                </el-form-item>
              </el-col>
              <el-col :span="10" class="grid-cell">
                <el-form-item label="每日预算(元):"  prop="budgetDaily" >
                  <el-input-number v-model.number="formData.budgetDaily" :clearable="true" placeholder="请输入每日预算,元" />
                </el-form-item>
              </el-col>
            </el-row>

            <el-row>
              <el-col :span="10" class="grid-cell">
                <el-form-item label="总曝光数:"  prop="impTotal" >
                  <el-input-number v-model.number="formData.impTotal" :clearable="true" placeholder="请输入总曝光数" />
                </el-form-item>
              </el-col>
              <el-col :span="10" class="grid-cell">
                <el-form-item label="每日曝光数:"  prop="impDaily" >
                  <el-input-number v-model.number="formData.impDaily" :clearable="true" placeholder="请输入每日曝光数" />
                </el-form-item>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="10" class="grid-cell">
                <el-form-item label="曝光频制:"  prop="impFrequency" >
                  <el-input-number v-model.number="formData.impFrequency" :clearable="true" placeholder="请输入曝光频制" />
                </el-form-item>
              </el-col>
              <el-col :span="10" class="grid-cell">
                <el-form-item label="曝光频控周期(分钟):"  prop="impFrequencyMinute" >
                  <el-input-number v-model.number="formData.impFrequencyMinute" :clearable="true" placeholder="请输入曝光频控周期" />
                </el-form-item>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="10" class="grid-cell">
                <el-form-item label="点击频控:"  prop="clkFrequency" >
                  <el-input-number v-model.number="formData.clkFrequency" :clearable="true" placeholder="请输入点击频控" />
                </el-form-item>
              </el-col>
              <el-col :span="10" class="grid-cell">
                <el-form-item label="点击频控周期(分钟):"  prop="clkFrequencyMinute" >
                  <el-input-number v-model.number="formData.clkFrequencyMinute" :clearable="true" placeholder="请输入点击频控周期" />
                </el-form-item>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="10" class="grid-cell">
                <el-form-item label="最小点击率(%):"  prop="ctrMax" >
                  <el-input-number v-model="formData.ctrMax" :precision="2" :clearable="true"  />
                </el-form-item>
              </el-col>
              <el-col :span="10" class="grid-cell">
                <el-form-item label="最大点击率(%):"  prop="ctrMin" >
                  <el-input-number v-model="formData.ctrMin" :precision="2" :clearable="true"  />
                </el-form-item>
              </el-col>
            </el-row>
          </el-form>
      </el-scrollbar>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog v-model="detailShow" style="width: 800px" lock-scroll :before-close="closeDetailShow" title="查看详情" destroy-on-close>
      <el-scrollbar height="550px">
        <el-descriptions column="1" border>
                <el-descriptions-item label="名称">
                        {{ formData.name }}
                </el-descriptions-item>
                <el-descriptions-item label="描述">
                        {{ formData.desc }}
                </el-descriptions-item>
                <el-descriptions-item label="状态">
                        <el-switch v-model="formData.status"></el-switch>
                </el-descriptions-item>
                <el-descriptions-item label="投放方式">
                        {{ filterDict(formData.mode,adModeOptions) }}
                </el-descriptions-item>
                <el-descriptions-item label="时区">
                        {{ filterDict(formData.timezone,timezoneOptions) }}
                </el-descriptions-item>
                <el-descriptions-item label="开始时间">
                      {{ formatDate(formData.startAt) }}
                </el-descriptions-item>
                <el-descriptions-item label="结束时间">
                      {{ formatDate(formData.endAt) }}
                </el-descriptions-item>
                <el-descriptions-item label="总预算(元)">
                        {{ formData.budgetTotal }}
                </el-descriptions-item>
                <el-descriptions-item label="每日预算(元)">
                        {{ formData.budgetDaily }}
                </el-descriptions-item>
                <el-descriptions-item label="总曝光数">
                        {{ formData.impTotal }}
                </el-descriptions-item>
                <el-descriptions-item label="每日曝光数">
                        {{ formData.impDaily }}
                </el-descriptions-item>
                <el-descriptions-item label="曝光频制">
                        {{ formData.impFrequency }}
                </el-descriptions-item>
                <el-descriptions-item label="曝光频控周期">
                        {{ formData.impFrequencyMinute }}
                </el-descriptions-item>
                <el-descriptions-item label="点击频控">
                        {{ formData.clkFrequency }}
                </el-descriptions-item>
                <el-descriptions-item label="点击频控周期">
                        {{ formData.clkFrequencyMinute }}
                </el-descriptions-item>
                <el-descriptions-item label="最小点击率(%)">
                        {{ formData.ctrMax }}
                </el-descriptions-item>
                <el-descriptions-item label="最大点击率(%)">
                        {{ formData.ctrMin }}
                </el-descriptions-item>
        </el-descriptions>
      </el-scrollbar>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createPlan,
  deletePlan,
  deletePlanByIds,
  updatePlan,
  findPlan,
  getPlanList
} from '@/api/plan'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'


defineOptions({
    name: 'Plan'
})

// 自动化生成的字典（可能为空）以及字段
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
        endAt: undefined,
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
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               status : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               mode : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               timezone : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               startAt : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
const campaignPath = ref('')

const router = useRouter()

const setCampaignPath = () => {

  for (let r of router.getRoutes()) {
    if (r.name === 'campaign') {
      campaignPath.value = r.path
      break
    }
  }
}

setCampaignPath()

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getPlanList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    adStatusOptions.value = await getDictFunc('adStatus')
    adModeOptions.value = await getDictFunc('adMode')
    timezoneOptions.value = await getDictFunc('timezone')
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deletePlanFunc(row)
        })
    }


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
      const res = await deletePlanByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        deleteVisible.value = false
        getTableData()
      }
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updatePlanFunc = async(row) => {
    const res = await findPlan({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.replan
        dialogFormVisible.value = true
    }
}

// 定义选项改变时的事件
const handleSwitchChange = async(row) => {
  const res = await findPlan({ ID: row.ID })
  if (res.code === 0) {
      res.data.replan.status = row.status
      const res2 = await updatePlan(res.data.replan)
      if (res2.code === 0) {
        console.log("修改成功")
        //getTableData()
      }
  }
  
}

// 删除行
const deletePlanFunc = async (row) => {
    const res = await deletePlan({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)


// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findPlan({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.replan
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
          name: '',
          desc: '',
          status: true,
          mode: undefined,
          timezone: undefined,
          startAt: new Date(),
          endAt: undefined,
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
          }
}


// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
    console.log(adStatusOptions.value[0])
    //formData.value.status = adStatusOptions.value[0].value
    formData.value.timezone = timezoneOptions.value[0].value
    formData.value.mode = adModeOptions.value[0].value
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        name: '',
        desc: '',
        status: undefined,
        mode: undefined,
        timezone: undefined,
        startAt: new Date(),
        endAt: undefined,
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
        }
}
// 弹窗确定
const enterDialog = async () => {
      console.log(formData.value.status)
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
                closeDialog()
                getTableData()
              }
      })
}

</script>

<style>

</style>
