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
        <el-form-item label="计划" prop="plan_id">
         <el-select v-model="searchInfo.plan_id" placeholder="计划" >
            <!-- <el-option :key="key" :label="plan.name" :value="plan.ID" /> -->
            <el-option v-for="(item,key) in plans" :key="key" :label="item.name" :value="item.ID" />
          </el-select>
        </el-form-item>
            <el-form-item label="状态" prop="status">
            <el-select v-model="searchInfo.status" clearable placeholder="请选择">
                <el-option
                    key="true"
                    label="是"
                    value="true">
                </el-option>
                <el-option
                    key="false"
                    label="否"
                    value="false">
                </el-option>
            </el-select>
            </el-form-item>
           <el-form-item label="出价方式" prop="bid_method">
            <el-select v-model="searchInfo.bid_method" clearable placeholder="请选择" @clear="()=>{searchInfo.bid_method=undefined}">
              <el-option v-for="(item,key) in bidMethodOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
           <el-form-item label="出价模式" prop="bid_mode">
            <el-select v-model="searchInfo.bid_mode" clearable placeholder="请选择" @clear="()=>{searchInfo.bid_mode=undefined}">
              <el-option v-for="(item,key) in bidModeOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
        <el-form-item label="品牌名称" prop="brand">
         <el-input v-model="searchInfo.brand" placeholder="搜索条件" />

        </el-form-item>
         <!-- <el-form-item label="允许虚拟" prop="allow_virtually">
            <el-select v-model="searchInfo.allow_virtually" clearable placeholder="请选择">
                <el-option
                    key="true"
                    label="是"
                    value="true">
                </el-option>
                <el-option
                    key="false"
                    label="否"
                    value="false">
                </el-option>
            </el-select>
        </el-form-item> -->
        <el-form-item label="创意方式" prop="creative_mode">
            
             <el-input v-model.number="searchInfo.creative_mode" placeholder="搜索条件" />

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
        <el-table-column align="left" label="计划" prop="plan.name" width="120"/>
        <el-table-column align="left" label="名称" prop="name" width="120" >
          <template #default="scope">
            <a :href="'#' + creativePath + '?cid=' + scope.row.ID" >{{ scope.row.name }}</a>
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作" min-width="250">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                查看详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateCampaignFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="plus" class="table-button" @click="createCreativeFunc(scope.row)">添加创意</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        <el-table-column align="left" label="描述" prop="desc" width="120" />
        <el-table-column align="left" label="状态" prop="status" width="120">
            <template #default="scope">
              <el-switch v-model="scope.row.status" @change="handleSwitchChange(scope.row)"></el-switch>
            </template>
        </el-table-column>
        <!-- <el-table-column align="left" label="虚拟活动" prop="is_virtually" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.is_virtually) }}</template>
        </el-table-column>
        <el-table-column align="left" label="允许混量" prop="allow_virtually" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.allow_virtually) }}</template>
        </el-table-column> -->
         <el-table-column align="left" label="开始时间" width="180">
            <template #default="scope">{{ formatDateTo(scope.row.start_at, 'yyyy-MM-dd') }}</template>
         </el-table-column>
         <el-table-column align="left" label="结束时间" width="180">
            <template #default="scope">{{ formatDateTo(scope.row.end_at, 'yyyy-MM-dd') }}</template>
         </el-table-column>
        <el-table-column align="left" label="总预算(元)" prop="budget_total" width="120" />
        <el-table-column align="left" label="每日预(元)" prop="budget_daily" width="120" />
        <el-table-column align="left" label="总曝光数" prop="imp_total" width="120" />
        <el-table-column align="left" label="每日曝光数" prop="imp_daily" width="120" />
        <el-table-column align="left" label="曝光频制" prop="imp_frequency" width="120" />
        <el-table-column align="left" label="曝光频控周期" prop="imp_frequency_minute" width="120" />
        <el-table-column align="left" label="点击频控" prop="clk_frequency" width="120" />
        <el-table-column align="left" label="点击频控周期" prop="clk_frequency_minute" width="120" />
        <el-table-column align="left" label="最小点击率(%)" prop="ctr_max" width="120" />
        <el-table-column align="left" label="最大点击率(%)" prop="ctr_min" width="120" />
        <el-table-column align="left" label="出价策略id" prop="policy_id" width="120" />
        <el-table-column align="left" label="出价方式" prop="bid_method" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.bid_method,bidMethodOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="出价策略" prop="bid_price" width="120" />
        <el-table-column align="left" label="出价模式" prop="bid_mode" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.bid_mode,bidModeOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="品牌名称" prop="brand" width="120" />
        <el-table-column align="left" label="创意方式" prop="creative_mode" width="120" />
        
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
          <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="140px">
            <el-row>
              <el-col :span="8" class="grid-cell">
                <el-form-item label="计划:"  prop="plan_id" >
                  <el-select v-model="formData.plan_id" placeholder="计划" >
                    <!-- <el-option :key="key" :label="plan.name" :value="plan.ID" :disabled="true" /> -->
                    <el-option v-for="(item,key) in plans" :key="key" :label="item.name" :value="item.ID" />
                  </el-select>
                </el-form-item>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="15" class="grid-cell">
                <el-form-item label="名称:"  prop="name" >
                  <el-input v-model="formData.name" :clearable="true"  placeholder="请输入名称" style="width:100%" />
                </el-form-item>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="15" class="grid-cell">
                <el-form-item label="描述:"  prop="desc" >
                  <el-input v-model="formData.desc" :clearable="true"  placeholder="请输入描述" style="width:100%" />
                </el-form-item>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="6" class="grid-cell">
                <el-form-item label="状态:"  prop="status" >
                  <el-switch v-model="formData.status" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
                </el-form-item>
              </el-col>

              <el-col :span="6" class="grid-cell">
                <el-form-item label="虚拟活动:"  prop="is_virtually" >
                  <el-switch v-model="formData.is_virtually" active-text="是" inactive-text="否" clearable ></el-switch>
                </el-form-item>
              </el-col>
              <el-col :span="6" class="grid-cell">
                <el-form-item label="混量:"  prop="allow_virtually" >
                  <el-switch v-model="formData.allow_virtually"  active-text="是" inactive-text="否" clearable ></el-switch>
                </el-form-item>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="10" class="grid-cell">
                <el-form-item label="开始时间:"  prop="start_at" >
                  <el-date-picker v-model="formData.start_at" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
                </el-form-item>
              </el-col>
              <el-col :span="10" class="grid-cell">
                <el-form-item label="结束时间:"  prop="end_at" >
                  <el-date-picker v-model="formData.end_at" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
                </el-form-item>
              </el-col>
            </el-row>

            <el-row>
              <el-col :span="10" class="grid-cell">
                <el-form-item label="总预算(元):"  prop="budget_total" >
                  <el-input-number v-model.number="formData.budget_total" :clearable="true" placeholder="请输入总预算,元" />
                </el-form-item>
              </el-col>
              <el-col :span="10" class="grid-cell">
                <el-form-item label="每日预算(元):"  prop="budget_daily" >
                  <el-input-number v-model.number="formData.budget_daily" :clearable="true" placeholder="请输入每日预算,元" />
                </el-form-item>
              </el-col>
            </el-row>

            <el-row>
              <el-col :span="10" class="grid-cell">
                <el-form-item label="总曝光数:"  prop="imp_total" >
                  <el-input-number v-model.number="formData.imp_total" :clearable="true" placeholder="请输入总曝光数" />
                </el-form-item>
              </el-col>
              <el-col :span="10" class="grid-cell">
                <el-form-item label="每日曝光数:"  prop="imp_daily" >
                  <el-input-number v-model.number="formData.imp_daily" :clearable="true" placeholder="请输入每日曝光数" />
                </el-form-item>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="10" class="grid-cell">
                <el-form-item label="曝光频制:"  prop="imp_frequency" >
                  <el-input-number v-model.number="formData.imp_frequency" :clearable="true" placeholder="请输入曝光频制" />
                </el-form-item>
              </el-col>
              <el-col :span="10" class="grid-cell">
                <el-form-item label="曝光频控周期(分钟):"  prop="imp_frequency_minute" >
                  <el-input-number v-model.number="formData.imp_frequency_minute" :clearable="true" placeholder="请输入曝光频控周期" />
                </el-form-item>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="10" class="grid-cell">
                <el-form-item label="点击频控:"  prop="clk_frequency" >
                  <el-input-number v-model.number="formData.clk_frequency" :clearable="true" placeholder="请输入点击频控" />
                </el-form-item>
              </el-col>
              <el-col :span="10" class="grid-cell">
                <el-form-item label="点击频控周期(分钟):"  prop="clk_frequency_minute" >
                  <el-input-number v-model.number="formData.clk_frequency_minute" :clearable="true" placeholder="请输入点击频控周期" />
                </el-form-item>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="10" class="grid-cell">
                <el-form-item label="最小点击率(%):"  prop="ctr_max" >
                  <el-input-number v-model="formData.ctr_max" :precision="2" :clearable="true"  />
                </el-form-item>
              </el-col>
              <el-col :span="10" class="grid-cell">
                <el-form-item label="最大点击率(%):"  prop="ctr_min" >
                  <el-input-number v-model="formData.ctr_min" :precision="2" :clearable="true"  />
                </el-form-item>
              </el-col>
            </el-row>
            <!--
            <el-form-item label="投放时间段:"  prop="hours" >
              <el-input v-model.number="formData.hours" :clearable="true" placeholder="请输入投放时间段" />
            </el-form-item>
            <el-form-item label="时间段预览:"  prop="__hours" >
                <pre class="bg-gray-100 p-2 rounded">{{ JSON.stringify(selectedHours) }}</pre>
              </el-form-item>
            -->
            <el-row>
              <el-col :span="24" class="grid-cell">
                <el-form-item label="投放时间段:"  prop="_hours" >
                  <button
                    v-for="(selected, index) in selectedHours"
                    :key="index"
                    @click="toggleHour(index)"
                    :class="[
                      'w-8 h-8 text-xs rounded flex items-center justify-center',
                      selected ? 'bg-blue-500 text-white hover:bg-blue-600' : 'bg-gray-200 hover:bg-gray-300'
                    ]"
                  >
                    {{ String(index).padStart(2, '0') }}
                  </button>
                </el-form-item>
              </el-col>
            </el-row>
            <el-form-item label="定向包id:"  prop="target_id" >
              <el-input v-model.number="formData.target_id" :clearable="true" placeholder="请输入定向包id" />
            </el-form-item>
            <el-form-item label="黑白名单id:"  prop="bwlist_id" >
              <el-input v-model.number="formData.bwlist_id" :clearable="true" placeholder="请输入黑白名单id" />
            </el-form-item>
            <el-form-item label="出价策略id:"  prop="policy_id" >
              <el-input v-model.number="formData.policy_id" :clearable="true" placeholder="请输入出价策略id" />
            </el-form-item>

            <el-row>
              <el-col :span="12" class="grid-cell">
                <el-form-item label="出价方式:"  prop="bid_method" >
                  <el-select v-model="formData.bid_method" placeholder="出价方式" :clearable="true" >
                    <el-option v-for="(item,key) in bidMethodOptions" :key="key" :label="item.label" :value="item.value" />
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="12" class="grid-cell">
                <el-form-item label="出价模式:"  prop="bid_mode" >
                  <el-select v-model="formData.bid_mode" placeholder="出价模式"  :clearable="true" >
                    <el-option v-for="(item,key) in bidModeOptions" :key="key" :label="item.label" :value="item.value" />
                  </el-select>
                </el-form-item>
              </el-col>
            </el-row>

            <el-row>
              <el-col :span="12" class="grid-cell">
                <el-form-item label="出价(元):"  prop="bid_price">
                  <el-input-number v-model="formData.bid_price"  :precision="2" :clearable="true"  />
                </el-form-item>
              </el-col>
              <el-col :span="12" class="grid-cell">
                <el-form-item label="出价率:"  prop="bid_rate" >
                  <el-input-number v-model="formData.bid_rate"  :precision="2" :clearable="true"  />%
                </el-form-item>
              </el-col>
            </el-row>
            
            <el-form-item label="品牌名称:"  prop="brand" >
              <el-input v-model="formData.brand" :clearable="true"  placeholder="请输入品牌名称" />
            </el-form-item>
            <el-form-item label="创意方式:"  prop="creative_mode" >
              <el-input v-model.number="formData.creative_mode" :clearable="true" placeholder="请输入创意方式" />
            </el-form-item>
            <el-form-item label="曝光监测:"  prop="imp_track_url" >
              <el-input v-model="formData.imp_track_url" :clearable="true"  placeholder="请输入曝光监测" />
            </el-form-item>
            <el-form-item label="点击监测:"  prop="click_track_url" >
              <el-input v-model="formData.click_track_url" :clearable="true"  placeholder="请输入点击监测" />
            </el-form-item>
            <el-form-item label="落地页h5:"  prop="h5" >
              <el-input v-model="formData.h5" :clearable="true"  placeholder="请输入落地页" />
            </el-form-item>
            <el-form-item label="deeplink:"  prop="deeplink" >
              <el-input v-model="formData.deeplink" :clearable="true"  placeholder="请输入deeplink字段" />
            </el-form-item>
            <el-form-item label="universal_link:"  prop="universal_link" >
              <el-input v-model="formData.universal_link" :clearable="true"  placeholder="请输入universal_link字段" />
            </el-form-item>
            <el-form-item label="动态对应的url:"  prop="_adm" >
              <el-input v-model="_adm" type="text"  :clearable="true"  placeholder="请输入动态代码url" />
            </el-form-item>
            <el-form-item label="动态代码:"  prop="adm" >
              <el-input v-model="formData.adm" type="textarea" :rows="10" :clearable="true"  placeholder="请输入动态代码" />
            </el-form-item>
          </el-form>
      </el-scrollbar>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
          <el-button type="primary" @click="openDialogCreative">确定并添加创意</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog v-model="detailShow" style="width: 800px" lock-scroll :before-close="closeDetailShow" title="查看详情" destroy-on-close>
      <el-scrollbar height="550px">
        <el-descriptions column="1" border>
                <el-descriptions-item label="计划">
                        {{ formData.plan.name }}
                </el-descriptions-item>
                <el-descriptions-item label="名称">
                        {{ formData.name }}
                </el-descriptions-item>
                <el-descriptions-item label="描述">
                        {{ formData.desc }}
                </el-descriptions-item>
                <el-descriptions-item label="状态">
                    {{ formatBoolean(formData.status) }}
                </el-descriptions-item>
                <el-descriptions-item label="是否虚拟活动">
                    {{ formatBoolean(formData.is_virtually) }}
                </el-descriptions-item>
                <el-descriptions-item label="允许虚拟混量">
                    {{ formatBoolean(formData.allow_virtually) }}
                </el-descriptions-item>
                <el-descriptions-item label="开始时间">
                      {{ formatDate(formData.start_at) }}
                </el-descriptions-item>
                <el-descriptions-item label="结束时间">
                      {{ formatDate(formData.end_at) }}
                </el-descriptions-item>
                <el-descriptions-item label="总预算,元">
                        {{ formData.budget_total }}
                </el-descriptions-item>
                <el-descriptions-item label="每日预算,元">
                        {{ formData.budget_daily }}
                </el-descriptions-item>
                <el-descriptions-item label="总曝光数">
                        {{ formData.imp_total }}
                </el-descriptions-item>
                <el-descriptions-item label="每日曝光数">
                        {{ formData.imp_daily }}
                </el-descriptions-item>
                <el-descriptions-item label="曝光频制">
                        {{ formData.imp_frequency }}
                </el-descriptions-item>
                <el-descriptions-item label="曝光频控周期">
                        {{ formData.imp_frequency_minute }}
                </el-descriptions-item>
                <el-descriptions-item label="点击频控">
                        {{ formData.clk_frequency }}
                </el-descriptions-item>
                <el-descriptions-item label="点击频控周期">
                        {{ formData.clk_frequency_minute }}
                </el-descriptions-item>
                <el-descriptions-item label="最小点击率(%)">
                        {{ formData.ctr_max }}
                </el-descriptions-item>
                <el-descriptions-item label="最大点击率(%)">
                        {{ formData.ctr_min }}
                </el-descriptions-item>
                <el-descriptions-item label="投放时间段">
                        {{ formData.hours }}
                </el-descriptions-item>
                <el-descriptions-item label="定向包id">
                        {{ formData.target_id }}
                </el-descriptions-item>
                <el-descriptions-item label="黑白名单id">
                        {{ formData.bwlist_id }}
                </el-descriptions-item>
                <el-descriptions-item label="出价策略id">
                        {{ formData.policy_id }}
                </el-descriptions-item>
                <el-descriptions-item label="出价方式">
                        {{ filterDict(formData.bid_method,bidMethodOptions) }}
                </el-descriptions-item>
                <el-descriptions-item label="出价策略">
                        {{ formData.bid_price }}
                </el-descriptions-item>
                <el-descriptions-item label="出价模式">
                        {{ filterDict(formData.bid_mode,bidModeOptions) }}
                </el-descriptions-item>
                <el-descriptions-item label="品牌名称">
                        {{ formData.brand }}
                </el-descriptions-item>
                <el-descriptions-item label="创意方式">
                        {{ formData.creative_mode }}
                </el-descriptions-item>
                <el-descriptions-item label="曝光监测">
                        {{ formData.imp_track_url }}
                </el-descriptions-item>
                <el-descriptions-item label="点击监测">
                        {{ formData.click_track_url }}
                </el-descriptions-item>
                <el-descriptions-item label="落地页">
                        {{ formData.h5 }}
                </el-descriptions-item>
                <el-descriptions-item label="deeplink">
                        {{ formData.deeplink }}
                </el-descriptions-item>
                <el-descriptions-item label="universal_link">
                        {{ formData.universal_link }}
                </el-descriptions-item>
        </el-descriptions>
      </el-scrollbar>
    </el-dialog>

    <el-dialog v-model="dialogFormCreative" :before-close="closeDialogCreative" :title="'添加创意'" destroy-on-close>
      <el-scrollbar height="500px">
          <el-form :model="creatives" label-position="right" ref="elFormRefCreative" :rules="rule" label-width="140px">
            <el-form-item label="计划:"  prop="plan_id" >
              <el-input v-model="creatives.plan_id" placeholder="计划" />
            </el-form-item>
            <el-form-item label="活动:"  prop="campaign_id" >
              <el-input v-model="creatives.campaign_id" placeholder="活动" />
            </el-form-item>
            <el-form-item label="标题:"  prop="title" >
              <el-input v-model="creatives.title" :clearable="true"  placeholder="请输入标题" />
            </el-form-item>
            <el-form-item label="描述:"  prop="desc" >
              <el-input v-model="creatives.desc" :clearable="true"  placeholder="请输入描述" />
            </el-form-item>
            <el-form-item label="行动语:"  prop="desc" >
              <el-input v-model="creatives.button" :clearable="true"  placeholder="请输入行动语" />
            </el-form-item>
            <el-form-item label="图片:"  prop="images">
              <SelectMaterial
                v-model="creatives.images"
                multiple=true
                 file-type="image"
                @selection-change="handleMaterialChangeImages"
                />
            </el-form-item>
            <el-form-item label="视频:"  prop="videos">
              <SelectMaterial
                v-model="creatives.videos"
                multiple=true
                 file-type="video"
                @selection-change="handleMaterialChangeVideos"
                />
            </el-form-item>
          </el-form>
      </el-scrollbar>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialogCreative">取 消</el-button>
          <el-button type="primary" @click="enterDialogCreative">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createCampaign,
  deleteCampaign,
  deleteCampaignByIds,
  updateCampaign,
  findCampaign,
  getCampaignList
} from '@/api/campaign'

import {
  createCreatives
} from '@/api/creative'

import {
  findPlan,
  getPlanList,
} from '@/api/plan'

import SelectMaterial from '@/components/selectMaterial/selectMaterial.vue'
// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatDateTo, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive, watchEffect } from 'vue'
import { useRouter } from 'vue-router'
import { toBitInt, toArrFromBitInt } from '@/utils/bit'



defineOptions({
    name: 'Campaign'
})

const plans = ref([])
const materials = ref([])

        
const initCreative = ref({
  plan_id: 0,
  campaign_id: 0,
  material_id: 0,
  title: '',
  desc: '',
  button: '',
})

// 时间选择器
const selectedHours = ref(new Array(24).fill(0))

const toggleHour = (index) => {
  selectedHours.value[index] = selectedHours.value[index] === 0 ? 1 : 0
  formData.value.hours = binaryArrayToDecimal(selectedHours.value)
  console.log(toBitInt(selectedHours.value))
  //console.log(toArrFromBitInt(toBitInt(selectedHours.value)))
  console.log(binaryArrayToDecimal(selectedHours.value))
}

const binaryArrayToDecimal = (arr) => { 
  return arr.reduce((acc, curr, index) => {
    return acc + curr * Math.pow(2, arr.length - 1 - index)
  }, 0)
}

// 将十进制数转为二进制数组
const decimalToBinaryArray = (decimal) => {
  const binaryStr = decimal.toString(2).padStart(24, '0')
  return Array.from(binaryStr).map(Number)
}


// 自动化生成的字典（可能为空）以及字段
const bidMethodOptions = ref([])
const bidModeOptions = ref([])
const formData = ref({
        plan_id: 0,
        plan: {},
        name: '',
        desc: '',
        status: true,
        is_virtually: false,
        allow_virtually: false,
        start_at: new Date(),
        end_at: new Date(),
        budget_total: 0,
        budget_daily: 0,
        imp_total: 0,
        imp_daily: 0,
        imp_frequency: 0,
        imp_frequency_minute: 0,
        clk_frequency: 0,
        clk_frequency_minute: 0,
        ctr_max: 0,
        ctr_min: 0,
        hours: 0,
        target_id: 0,
        bwlist_id: 0,
        policy_id: 0,
        bid_method: undefined,
        bid_price: 0,
        bid_mode: undefined,
        brand: '',
        creative_mode: 0,
        imp_track_url: '',
        click_track_url: '',
        h5: '',
        deeplink: '',
        universal_link: '',
        creatives: [initCreative],
        })

const creatives = ref({
  plan_id: 0,
  campaign_id: 0,
  title: '',
  desc: '',
  button: '',
  videos: [],
  images: [],
})

const _adm = ref('');
// 监听 urlInput 的变化并发送请求
watchEffect(() => {
      if (_adm.value.startsWith('http')) {
        console.log(_adm.value)
        fetch(_adm.value)
          .then(response => {
            if (!response.ok) {
              throw new Error('Network response was not ok');
            }
            return response.text();
          })
          .then(data => {
            formData.value.adm = data; // 将响应的 HTML 赋值给 htmlInput
          })
          .catch(error => {
            console.error('Error fetching URL:', error);
            formData.value.adm = 'Error fetching URL'; // 处理错误
          });
      }
    });

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
  }]
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
const elFormRefCreative = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

const creativePath = ref('')

const router = useRouter()

const setcreativePath = () => {

  for (let r of router.getRoutes()) {
    if (r.name === 'creative') {
      creativePath.value = r.path
      break
    }
  }
}

setcreativePath()

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
    if (searchInfo.value.status === ""){
        searchInfo.value.status=null
    }
    if (searchInfo.value.allow_virtually === ""){
        searchInfo.value.allow_virtually=null
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

const handleMaterialChangeImages = (val) => {
  console.log(val)
  if (props.multiple && val && val.length > 0) {
    for (let i = 0; i < val.length; i++) {
      creatives.value.images.push(val[i].id)
    }
  } else if (!props.multiple) {
    creatives.value.images.push(val.id)
  }
  /* 
  if (formData.value.creatives && formData.value.creatives.length > 0) {
    if (formData.value.creatives.length === 1) {
      formData.value.creatives[0].material_id = val.id
    } else {
      for ((c, idx) in formData.value.creatives) {

      }
    }
  } */
}
const handleMaterialChangeVideos = (val) => {
  console.log(val)
  if (props.multiple && val && val.length > 0) {
    for (let i = 0; i < val.length; i++) {
      creatives.value.videos.push(val[i].id)
      /* if (i > formData.value.creatives.length - 1) {
        let a = ref({
          plan_id: initCreative.value.plan_id,
          campaign_id: initCreative.value.campaign_id,
          material_id: val[i].id,
          title: initCreative.value.title,
          desc: initCreative.value.desc,
          button: initCreative.value.button,
        })
        formData.value.creatives.push(a)
      } else {
        formData.value.creatives[i].value.material_id = c.id
      } */
    }
  } else if (!props.multiple) {
    creatives.value.videos.push(val.id)
  }
  /* 
  if (formData.value.creatives && formData.value.creatives.length > 0) {
    if (formData.value.creatives.length === 1) {
      formData.value.creatives[0].material_id = val.id
    } else {
      for ((c, idx) in formData.value.creatives) {

      }
    }
  } */
}

// 查询
const getTableData = async() => {
  const table = await getCampaignList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    bidMethodOptions.value = await getDictFunc('bidMethod')
    bidModeOptions.value = await getDictFunc('bidMode')
}

// 获取需要的字典 可能为空 按需保留
setOptions()

// 获得计划
const getPlanDetails = async (pid) => {
  const res = await findPlan({ ID: pid })
  if (res.code === 0) {
    formData.value.plan_id = res.data.replan.ID
    searchInfo.value.plan_id = res.data.replan.ID
    plans.value.push(res.data.replan)
    getTableData()
  }
}

const getPlans = async() => {
  const table = await getPlanList({ page: page.value, pageSize: pageSize.value})
  if (table.code === 0) {
    plans.value = table.data.list
    formData.value.plan_id = plans.value[0].ID
  }
}

const setPlan = () => {
  const router = useRouter()
  const pid = router.currentRoute.value.query['pid']
  if (pid && pid !== '') {
    getPlanDetails(pid)
  } else {
    getPlans()
    getTableData()
  }
} 

setPlan()

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
            deleteCampaignFunc(row)
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
      const res = await deleteCampaignByIds({ ids })
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


// 定义选项改变时的事件
const handleSwitchChange = async(row) => {
  const res = await findCampaign({ ID: row.ID })
  if (res.code === 0) {
      res.data.recampaign.status = row.status
      const res2 = await updateCampaign(res.data.recampaign)
      if (res2.code === 0) {
        console.log("修改成功")
        //getTableData()
      }
  }
  
}

// 更新行
const updateCampaignFunc = async(row) => {
    const res = await findCampaign({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.recampaign
        selectedHours.value = decimalToBinaryArray(formData.value.hours)
        console.log(selectedHours.value)
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteCampaignFunc = async (row) => {
    const res = await deleteCampaign({ ID: row.ID })
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
const dialogFormCreative = ref(false)


// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findCampaign({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.recampaign
    selectedHours.value = decimalToBinaryArray(formData.value.hours)
    console.log(selectedHours.value)
    console.log(res.data.recampaign.plan.name)
    console.log(res.data.recampaign.materials)
    if (plans.length > 0) {
      formData.value.plan_id = plans.value[0].ID
    }
    openDetailShow()
  }
}

// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
          plan_id: 0,
          plan: {},
          name: '',
          desc: '',
          status: true,
          allow_virtually: false,
          is_virtually: false,
          start_at: new Date(),
          end_at: new Date(),
          budget_total: 0,
          budget_daily: 0,
          imp_total: 0,
          imp_daily: 0,
          imp_frequency: 0,
          imp_frequency_minute: 0,
          clk_frequency: 0,
          clk_frequency_minute: 0,
          ctr_max: 0,
          ctr_min: 0,
          hours: 0,
          target_id: 0,
          bwlist_id: 0,
          policy_id: 0,
          bid_method: undefined,
          bid_price: 0,
          bid_mode: undefined,
          brand: '',
          creative_mode: 0,
          imp_track_url: '',
          click_track_url: '',
          h5: '',
          deeplink: '',
          universal_link: '',
          creatives: [initCreative],
          }
}


// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}


// 打开弹窗
const openDialogCreative = async () => {
    let succ = await enterDialog()
    console.log(succ)
    if (succ) {
      openCreativeDialog()
    }
}

const createCreativeFunc = (row) => {
  creatives.value.plan_id = row.plan_id
  creatives.value.campaign_id = row.ID
  openCreativeDialog()
}

const openCreativeDialog = () => {
  type.value = 'create'
  dialogFormCreative.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        plan_id: 0,
        plan: {},
        name: '',
        desc: '',
        status: true,
        is_virtually: false,
        allow_virtually: false,
        start_at: new Date(),
        end_at: new Date(),
        budget_total: 0,
        budget_daily: 0,
        imp_total: 0,
        imp_daily: 0,
        imp_frequency: 0,
        imp_frequency_minute: 0,
        clk_frequency: 0,
        clk_frequency_minute: 0,
        ctr_max: 0,
        ctr_min: 0,
        hours: 0,
        target_id: 0,
        bwlist_id: 0,
        policy_id: 0,
        bid_method: undefined,
        bid_price: 0,
        bid_mode: undefined,
        brand: '',
        creative_mode: 0,
        imp_track_url: '',
        click_track_url: '',
        h5: '',
        deeplink: '',
        universal_link: '',
        creatives: [initCreative],
        }
}


// 关闭弹窗
const closeDialogCreative = () => {
    dialogFormCreative.value = false
    formData.value.creatives = []
}


// 弹窗确定
const enterDialog = async () => {
     return elFormRef.value?.validate( async (valid) => {
             if (!valid) return false
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
                console.log(res)
                if (res.data.recampaign) {
                  creatives.value.plan_id = res.data.recampaign.plan_id
                  creatives.value.campaign_id = res.data.recampaign.ID
                } else if (formData.value.ID) {
                  creatives.value.plan_id = formData.value.plan_id
                  creatives.value.campaign_id = formData.value.ID
                }
                
                console.log(creatives.value)

                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
                closeDialog()
                getTableData()
                return true
              }
              return false
      })
}

// 弹窗确定
const enterDialogCreative = async () => {
     elFormRefCreative.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createCreatives(creatives.value)
                  break
                default:
                  res = await createCreatives(creatives.value)
                  break
              }
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
                closeDialogCreative()
                //getTableData()
              }
      })
}

</script>

<style>

</style>
