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
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
      </el-form-item>
        <el-form-item label="平台渠道" prop="platform">
            
             <el-input v-model.number="searchInfo.platform" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="广告位" prop="spot">
            
             <el-input v-model.number="searchInfo.spot" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="bundle" prop="bundle">
            
             <el-input v-model.number="searchInfo.bundle" placeholder="搜索条件" />

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
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="平台渠道" prop="platform" width="120">
          <template #default="scope">
           <span v-if="scope.row.platform === 0">不限</span>
           <span v-else>{{ scope.row.platform }}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="广告位" prop="spot" width="120">
          <template #default="scope">
           <span v-if="scope.row.spot === 0">不限</span>
           <span v-else>{{ scope.row.spot }}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="发布者" prop="publisher" width="120">
          <template #default="scope">
           <span v-if="scope.row.publisher === 0">不限</span>
           <span v-else>{{ scope.row.publisher }}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="bundle" prop="bundle" width="120">
          <template #default="scope">
           <span v-if="scope.row.bundle === 0">不限</span>
           <span v-else>{{ scope.row.bundle }}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="广告形式" prop="format" width="120">
          <template #default="scope">
           <span v-if="scope.row.format === 0">不限</span>
           <span v-else>{{ scope.row.format }}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="设备id信息要求" prop="identity" width="120" >
          <template #default="scope">
            {{ filterDict(scope.row.identity,identityOptions) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作系统" prop="os" width="120" />
        <el-table-column align="left" label="地区" prop="region" width="120" />
        <el-table-column align="left" label="平均出价" prop="price" width="120" />
        <el-table-column align="left" label="浮动范围(%)" prop="scope" width="120" />
        <el-table-column align="left" label="操作" min-width="120">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                查看详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updatePolicyFunc(scope.row)">变更</el-button>
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
          <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="120px">
            <el-form-item label="平台渠道:"  prop="platform" >
              <el-input v-model.number="formData.platform" :clearable="true" placeholder="请输入平台渠道" />
            </el-form-item>
            <el-form-item label="广告位:"  prop="spot" >
              <el-input v-model.number="formData.spot" :clearable="true" placeholder="请输入广告位" />
            </el-form-item>
            <el-form-item label="发布者:"  prop="publisher" >
              <el-input v-model.number="formData.publisher" :clearable="true" placeholder="请输入发布者" />
            </el-form-item>
            <el-form-item label="bundle:"  prop="bundle" >
              <el-input v-model.number="formData.bundle" :clearable="true" placeholder="请输入bundle" />
            </el-form-item>
            <el-form-item label="广告形式:"  prop="format" >
              <el-input v-model.number="formData.format" :clearable="true" placeholder="请输入广告形式" />
            </el-form-item>
            <el-form-item label="设备id信息要求:"  prop="identity" >
              <el-select v-model="formData.identity" placeholder="请选择" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in identityOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="操作系统:"  prop="os" >
              <el-input v-model.number="formData.os" :clearable="true" placeholder="请输入操作系统" />
            </el-form-item>
            <el-form-item label="地区:"  prop="region" >
              <el-input v-model.number="formData.region" :clearable="true" placeholder="请输入地区" />
            </el-form-item>
            <el-form-item label="平均出价:"  prop="price" >
              <el-input-number v-model="formData.price" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="浮动范围(%):"  prop="scope" >
              <el-input-number v-model.number="formData.scope" :clearable="true" :precision="0" placeholder="请输入浮动范围(%)" />
            </el-form-item>
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
                <el-descriptions-item label="平台渠道">
                        {{ formData.platform }}
                </el-descriptions-item>
                <el-descriptions-item label="广告位">
                        {{ formData.spot }}
                </el-descriptions-item>
                <el-descriptions-item label="发布者">
                        {{ formData.publisher }}
                </el-descriptions-item>
                <el-descriptions-item label="bundle">
                        {{ formData.bundle }}
                </el-descriptions-item>
                <el-descriptions-item label="广告形式">
                        {{ formData.format }}
                </el-descriptions-item>
                <el-descriptions-item label="是否有设备id信息">
                    {{ filterDict(formData.identity,identityOptions) }}
                </el-descriptions-item>
                <el-descriptions-item label="操作系统">
                        {{ formData.os }}
                </el-descriptions-item>
                <el-descriptions-item label="地区">
                        {{ formData.region }}
                </el-descriptions-item>
                <el-descriptions-item label="平均出价">
                        {{ formData.price }}
                </el-descriptions-item>
                <el-descriptions-item label="浮动范围(%)">
                        {{ formData.scope }}
                </el-descriptions-item>
        </el-descriptions>
      </el-scrollbar>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createPolicy,
  deletePolicy,
  deletePolicyByIds,
  updatePolicy,
  findPolicy,
  getPolicyList
} from '@/api/policy'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'Policy'
})

const identityOptions = ref([])


// 自动化生成的字典（可能为空）以及字段
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
    if (searchInfo.value.identity === ""){
        searchInfo.value.identity=undefined
    }
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
  const table = await getPolicyList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
  identityOptions.value = await getDictFunc('identityType')
  console.log(identityOptions.value)
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
            deletePolicyFunc(row)
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
      const res = await deletePolicyByIds({ ids })
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
const updatePolicyFunc = async(row) => {
    const res = await findPolicy({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.repolicy
        dialogFormVisible.value = true
    }
}


// 删除行
const deletePolicyFunc = async (row) => {
    const res = await deletePolicy({ ID: row.ID })
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
  const res = await findPolicy({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.repolicy
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
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
          }
}


// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
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
        }
}
// 弹窗确定
const enterDialog = async () => {
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
                closeDialog()
                getTableData()
              }
      })
}

</script>

<style>

</style>
