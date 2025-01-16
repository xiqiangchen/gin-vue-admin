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
        <el-form-item label="名称" prop="name">
         <el-input v-model="searchInfo.name" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="描述" prop="desc">
         <el-input v-model="searchInfo.desc" placeholder="搜索条件" />

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
        <el-table-column align="left" label="名称" prop="name" width="120" />
        <el-table-column align="left" label="描述" prop="desc" width="120" />
        <el-table-column align="left" label="渠道媒体" prop="adx" width="120" />
        <el-table-column align="left" label="广告形式" prop="ad_format" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.ad_format,adFormatOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="设备类型" prop="device_type" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.device_type,deviceTypeOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="操作系统" prop="os" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.os,osOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="定向类型" prop="target_type" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.target_type,targetTypeOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="行政区域" prop="region" width="120" />
        <el-table-column align="left" label="性别" prop="gender" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.gender,genderOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="操作" min-width="120">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                查看详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateTargetFunc(scope.row)">变更</el-button>
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
          <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="名称:"  prop="name" >
              <el-input v-model="formData.name" :clearable="true"  placeholder="请输入名称" />
            </el-form-item>
            <el-form-item label="描述:"  prop="desc" >
              <el-input v-model="formData.desc" :clearable="true"  placeholder="请输入描述" />
            </el-form-item>
            <el-form-item label="渠道媒体:"  prop="adx" >
              <el-checkbox-group v-model="formData.adx" @change="handleChangeMultiStrAdx" placeholder="请选择渠道媒体" style="width:100%" :clearable="true" >
                <el-checkbox-button key="unlimited" :label="0" :value="0">不限</el-checkbox-button>
                <el-checkbox-button v-for="(item,key) in adxOptions" :key="key" :label="item.value" :value="item.value">{{ item.label }}</el-checkbox-button>
              </el-checkbox-group>
            </el-form-item>
            <el-form-item label="广告形式:"  prop="ad_format" >
              <el-checkbox-group  v-model="formData.ad_format" placeholder="请选择广告形式"  @change="handleChangeMultiAdFormat" style="width:100%" :clearable="true" >
                <el-checkbox-button key="unlimited" :label="0" :value="0">不限</el-checkbox-button>
                <el-checkbox-button v-for="(item,key) in adFormatOptions" :key="key" :label="item.value" :value="item.value">{{ item.label }}</el-checkbox-button>
              </el-checkbox-group >
            </el-form-item>
            <el-form-item label="设备类型:"  prop="device_type" >
              <el-checkbox-group v-model="formData.device_type" placeholder="请选择设备类型" @change="handleChangeMultiDeviceType" style="width:100%" :clearable="true" >
                <el-checkbox-button key="unlimited" :label="0" :value="0">不限</el-checkbox-button>
                <el-checkbox-button v-for="(item,key) in deviceTypeOptions" :key="key" :label="item.value" :value="item.value">{{ item.label }}</el-checkbox-button>
              </el-checkbox-group>
            </el-form-item>
            <el-form-item label="操作系统:"  prop="os" >
              <el-checkbox-group v-model="formData.os" placeholder="请选择操作系统" @change="handleChangeMultiOs" style="width:100%" :clearable="true" >
                <el-checkbox-button key="unlimited" :label="0" :value="0">不限</el-checkbox-button>
                <el-checkbox-button v-for="(item,key) in osOptions" :key="key" :label="item.value" :value="item.value" >{{ item.label }}</el-checkbox-button>
              </el-checkbox-group>
            </el-form-item>
            <el-form-item label="定向类型:"  prop="target_type" >
              <el-checkbox-group v-model="formData.target_type" @change="handleChangeMultiTargetType" placeholder="请选择定向类型" style="width:100%" :clearable="true" >
                <el-checkbox-button key="unlimited" :label="0" :value="0">不限</el-checkbox-button>
                <el-checkbox-button v-for="(item,key) in targetTypeOptions" :key="key" :label="item.value" :value="item.value">{{ item.label }}</el-checkbox-button>
              </el-checkbox-group>
            </el-form-item>
            <el-form-item label="行政区域:"  prop="region" >
              <el-checkbox-group v-model="formData.region" @change="handleChangeMultiStr" placeholder="请选择定向地区" style="width:100%" :clearable="true" >
                <el-checkbox-button key="unlimited" :label="0" :value="0">不限</el-checkbox-button>
                <el-checkbox-button v-for="(item,key) in regionOptions" :key="key" :label="item.value" :value="item.value">{{ item.label }}</el-checkbox-button>
              </el-checkbox-group>
            </el-form-item>
            <el-form-item label="性别:"  prop="gender" >
              <el-checkbox-group v-model="formData.gender" placeholder="请选择性别" @change="handleChangeMultiGender" style="width:100%" :clearable="true" >
                <el-checkbox-button key="unlimited" :label="0" :value="0">不限</el-checkbox-button>
                <el-checkbox-button v-for="(item,key) in genderOptions" :key="key" :label="item.value" :value="item.value">{{ item.label }}</el-checkbox-button>
              </el-checkbox-group>
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
        <el-descriptions :column="1" border>
                <el-descriptions-item label="名称">
                        {{ formData.name }}
                </el-descriptions-item>
                <el-descriptions-item label="描述">
                        {{ formData.desc }}
                </el-descriptions-item>

                <el-descriptions-item label="渠道媒体">
                  <el-checkbox-group v-model="formData.adx" @change="handleChangeMultiStrAdx" placeholder="请选择渠道媒体" style="width:100%" :clearable="true" >
                    <el-checkbox-button key="unlimited" :label="0" :value="0">不限</el-checkbox-button>
                    <el-checkbox-button v-for="(item,key) in adxOptions" :key="key" :label="item.value" :value="item.value">{{ item.label }}</el-checkbox-button>
                  </el-checkbox-group>
                </el-descriptions-item>
                <el-descriptions-item label="广告形式">
                  <el-checkbox-group  v-model="formData.ad_format" placeholder="请选择广告形式"  @change="handleChangeMultiAdFormat" style="width:100%" :clearable="true" >
                    <el-checkbox-button key="unlimited" :label="0" :value="0">不限</el-checkbox-button>
                    <el-checkbox-button v-for="(item,key) in adFormatOptions" :key="key" :label="item.value" :value="item.value">{{ item.label }}</el-checkbox-button>
                  </el-checkbox-group >
                </el-descriptions-item>
                <el-descriptions-item label="设备类型">
                  <el-checkbox-group v-model="formData.device_type" placeholder="请选择设备类型" @change="handleChangeMultiDeviceType" style="width:100%" :clearable="true" >
                    <el-checkbox-button key="unlimited" :label="0" :value="0">不限</el-checkbox-button>
                    <el-checkbox-button v-for="(item,key) in deviceTypeOptions" :key="key" :label="item.value" :value="item.value">{{ item.label }}</el-checkbox-button>
                  </el-checkbox-group>
                </el-descriptions-item>
                <el-descriptions-item label="操作系统">
                  <el-checkbox-group v-model="formData.os" placeholder="请选择操作系统" @change="handleChangeMultiOs" style="width:100%" :clearable="true" >
                    <el-checkbox-button key="unlimited" :label="0" :value="0">不限</el-checkbox-button>
                    <el-checkbox-button v-for="(item,key) in osOptions" :key="key" :label="item.value" :value="item.value" >{{ item.label }}</el-checkbox-button>
                  </el-checkbox-group>
                </el-descriptions-item>
                <el-descriptions-item label="定向类型">
                  <el-checkbox-group v-model="formData.target_type" @change="handleChangeMultiTargetType" placeholder="请选择定向类型" style="width:100%" :clearable="true" >
                    <el-checkbox-button key="unlimited" :label="0" :value="0">不限</el-checkbox-button>
                    <el-checkbox-button v-for="(item,key) in targetTypeOptions" :key="key" :label="item.value" :value="item.value">{{ item.label }}</el-checkbox-button>
                  </el-checkbox-group>
                </el-descriptions-item>
                <el-descriptions-item label="行政区域">
                  <el-checkbox-group v-model="formData.region" @change="handleChangeMultiStr" placeholder="请选择定向地区" style="width:100%" :clearable="true" >
                    <el-checkbox-button key="unlimited" :label="0" :value="0">不限</el-checkbox-button>
                    <el-checkbox-button v-for="(item,key) in regionOptions" :key="key" :label="item.value" :value="item.value">{{ item.label }}</el-checkbox-button>
                  </el-checkbox-group>
                </el-descriptions-item>
                <el-descriptions-item label="性别">
                  <el-checkbox-group v-model="formData.gender" placeholder="请选择性别" @change="handleChangeMultiGender" style="width:100%" :clearable="true" >
                    <el-checkbox-button key="unlimited" :label="0" :value="0">不限</el-checkbox-button>
                    <el-checkbox-button v-for="(item,key) in genderOptions" :key="key" :label="item.value" :value="item.value">{{ item.label }}</el-checkbox-button>
                  </el-checkbox-group>
                </el-descriptions-item>
        </el-descriptions>
      </el-scrollbar>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createTarget,
  deleteTarget,
  deleteTargetByIds,
  updateTarget,
  findTarget,
  getTargetList
} from '@/api/target'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { toBitInt, toArrFromBitInt } from '@/utils/bit'
import { showDictLabel } from '@/utils/dictionary'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'Target'
})

// 自动化生成的字典（可能为空）以及字段
const targetTypeOptions = ref([])
const genderOptions = ref([])
const adFormatOptions = ref([])
const deviceTypeOptions = ref([])
const regionOptions = ref([])
const adxOptions = ref([])
const osOptions = ref([])
const formData = ref({
        name: '',
        desc: '',
        ad_format: [],
        device_type: [],
        os: [],
        target_type: [],
        region: [],
        adx: [],
        gender: [],
        })

// 验证规则
const rule = reactive({
               name : [{
                   required: true,
                   message: '定向包名称必填',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
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
  const table = await getTargetList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    targetTypeOptions.value = await getDictFunc('targetType')
    genderOptions.value = await getDictFunc('gender')
    adFormatOptions.value = await getDictFunc('adFormat')
    deviceTypeOptions.value = await getDictFunc('deviceType')
    osOptions.value = await getDictFunc('os')
    adxOptions.value = await getDictFunc('adx')
    regionOptions.value = await getDictFunc('region')
}

// 获取需要的字典 可能为空 按需保留
setOptions()



// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 需要转换的多选
const handleChangeMultiTargetType = (val) => {
  handleChangeMultiString(val, 'target_type');

}

const handleChangeMultiOs = (val) => {
  handleChangeMultiString(val, 'os');

}

const handleChangeMultiGender = (val) => {
  handleChangeMultiString(val, 'gender');

}

const handleChangeMultiAdFormat = (val) => {
  handleChangeMultiString(val, 'ad_format');

}

const handleChangeMultiDeviceType = (val) => {
  handleChangeMultiString(val, 'device_type');

}


const handleChangeMultiString = (val, key) => {
  if (val !== undefined && val.length > 0) {
    if (val[val.length-1] === 0) {
      formData.value[key] = [0];
    } else {
      formData.value[key] = formData.value[key].filter(item => item !== 0);
    }
  }
}

const handleChangeMultiStr = (val) => {
  handleChangeMultiString(val, 'region');
}

const handleChangeMultiStrAdx = (val) => {
  handleChangeMultiString(val, 'adx');
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteTargetFunc(row)
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
      const res = await deleteTargetByIds({ ids })
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
const updateTargetFunc = async(row) => {
    const res = await findTarget({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        const data = { ...res.data.retarget }
        const arrayFields = ['ad_format', 'device_type', 'os', 'target_type', 'region', 'adx', 'gender']
        arrayFields.forEach(field => {
            const value = data[field]
            if (!value || value === '') {
                data[field] = []
                return
            }
            if (typeof value !== 'string') return
            
            const values = value.indexOf(',') !== -1 ? value.split(',') : [value]
            data[field] = field === 'region' ? values : values.map(Number)
        })
        formData.value = data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteTargetFunc = async (row) => {
    const res = await deleteTarget({ ID: row.ID })
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
    const res = await findTarget({ ID: row.ID })
    if (res.code === 0) {
        const data = { ...res.data.retarget }
        const arrayFields = ['ad_format', 'device_type', 'os', 'target_type', 'region', 'adx', 'gender']
        arrayFields.forEach(field => {
            const value = data[field]
            if (!value || value === '') {
                data[field] = []
                return
            }
            if (typeof value !== 'string') return
            
            const values = value.indexOf(',') !== -1 ? value.split(',') : [value]
            data[field] = field === 'region' ? values : values.map(Number)
        })
        formData.value = data
        openDetailShow()
    }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
          name: '',
          desc: '',
          ad_format: [],
          device_type: [],
          os: [],
          target_type: [],
          adx: [],
          region: [],
          gender: [],
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
        name: '',
        desc: '',
        ad_format: [],
        device_type: [],
        os: [],
        target_type: [],
        region: [],
        adx: [],
        gender: [],
        }
}
// 弹窗确定
const enterDialog = async () => {
    elFormRef.value?.validate(async (valid) => {
        if (!valid) return
        const submitData = { ...formData.value }
        const arrayFields = ['ad_format', 'device_type', 'os', 'target_type', 'region', 'adx', 'gender']
        arrayFields.forEach(field => {
            if (Array.isArray(submitData[field])) {
                submitData[field] = submitData[field].join(',')
            }
        })
        
        let res
        switch (type.value) {
            case 'create':
                res = await createTarget(submitData)
                break
            case 'update':
                res = await updateTarget(submitData)
                break
            default:
                res = await createTarget(submitData)
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
