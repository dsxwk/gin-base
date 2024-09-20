<template>
  <div class="dataVisualize-box">
    <div class="card top-box">
      <div class="top-title">{{ funcs.lang('User Data') }}</div>
      <el-tabs v-model="tabActive" class="demo-tabs">
        <el-tab-pane v-for="item in tab" :key="item.name" :label="item.label" :name="item.name"></el-tab-pane>
      </el-tabs>
      <div class="top-content">
        <el-row :gutter="40">
          <el-col class="mb40" :xs="24" :sm="12" :md="12" :lg="6" :xl="6">
            <div class="item-left sle">
              <span class="left-title">{{ funcs.lang('Access Total') }}</span>
              <div class="img-box">
                <img src="./images/book-sum.png" alt="" />
              </div>
              <span class="left-number">800.132w</span>
            </div>
          </el-col>
          <el-col class="mb40" :xs="24" :sm="12" :md="12" :lg="8" :xl="8">
            <div class="item-center">
              <div class="gitee-traffic traffic-box">
                <div class="traffic-img">
                  <img src="./images/add_person.png" alt="" />
                </div>
                <span class="item-value">{{ userCount }}</span>
                <span class="traffic-name sle">{{ funcs.lang('Total Users') }}</span>
              </div>
              <div class="gitHub-traffic traffic-box">
                <div class="traffic-img">
                  <img src="./images/add_team.png" alt="" />
                </div>
                <span class="item-value">222</span>
                <span class="traffic-name sle">{{ funcs.lang('Today Registered') }}</span>
              </div>
              <div class="today-traffic traffic-box">
                <div class="traffic-img">
                  <img src="./images/today.png" alt="" />
                </div>
                <span class="item-value">333</span>
                <span class="traffic-name sle">{{ funcs.lang('Today Visits') }}</span>
              </div>
              <div class="yesterday-traffic traffic-box">
                <div class="traffic-img">
                  <img src="./images/book_sum.png" alt="" />
                </div>
                <span class="item-value">444</span>
                <span class="traffic-name sle">{{ funcs.lang('Yesterday Visits') }}</span>
              </div>
            </div>
          </el-col>
          <el-col class="mb40" :xs="24" :sm="24" :md="24" :lg="10" :xl="10">
            <div class="item-right">
              <div class="echarts-title">Gitee / GitHub {{ funcs.lang('Visit Proportion') }}</div>
              <div class="book-echarts">
                <Pie ref="pieRef" />
              </div>
            </div>
          </el-col>
        </el-row>
      </div>
    </div>
    <div class="card bottom-box">
      <div class="bottom-title">{{ funcs.lang('Data Source') }}</div>
      <div class="bottom-tabs">
        <el-tabs v-model="tabActive" class="demo-tabs">
          <el-tab-pane v-for="item in tab" :key="item.name" :label="item.label" :name="item.name"></el-tab-pane>
        </el-tabs>
      </div>
      <div class="curve-echarts">
        <Curve ref="curveRef" />
      </div>
    </div>
  </div>
</template>
<script setup>
import { ref, onMounted } from 'vue';
import Pie from '@/views/home/components/pie.vue';
import Curve from '@/views/home/components/curve.vue';
import userModule from '@/app/modules/admin/user';
import createService from '@/utils/service';
import Functions from '@/utils/functions';

const funcs = new Functions();
const userService = createService(userModule);
const userCount = ref(0);
const tabActive = ref(1);
const tab = [
  { label: funcs.lang('Future 7 Days'), name: 1 },
  { label: funcs.lang('Last 7 Days'), name: 2 },
  { label: funcs.lang('Last Month'), name: 3 },
  { label: funcs.lang('Last 3 Months'), name: 4 },
  { label: funcs.lang('Last 6 Months'), name: 5 },
  { label: funcs.lang('Last Year'), name: 6 }
];

onMounted(async () => {
  const result = await userService.list({page: 1, pageSize: 10});
  userCount.value = result?.data?.total;
});
</script>
<style scoped lang="scss">
@import './index.scss';
</style>