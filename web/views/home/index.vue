<template>
  <div class="dashboard-container layout-pd">
    <el-row class="website-info-row common-row-gap">
      <el-col :span="24">
        <el-card class="website-info-card">
          <div class="website-info-title">网站信息</div>
          <el-table :data="websiteInfo" border style="width: 100%;" class="website-info-table">
            <el-table-column prop="label" label="属性" width="150"></el-table-column>
            <el-table-column prop="value" label="内容"></el-table-column>
          </el-table>
        </el-card>
      </el-col>
    </el-row>
    <el-row :gutter="20" class="dashboard-cards common-row-gap">
      <el-col :span="6" v-for="card in cards" :key="card.title">
        <el-card class="dashboard-card">
          <div class="card-header">
            <span>{{ card.title }}</span>
            <el-icon :size="32" :color="card.color">
              <component :is="card.icon" />
            </el-icon>
          </div>
          <div class="card-value">{{ card.value }}</div>
          <div class="card-desc">{{ card.desc }}</div>
        </el-card>
      </el-col>
    </el-row>
    <el-row :gutter="20" class="dashboard-main common-row-gap">
      <el-col :span="8">
        <el-card class="dashboard-quick-card">
          <template #header>快捷入口</template>
          <div class="quick-actions">
            <el-row :gutter="12">
              <el-col :span="12" v-for="(action, idx) in quickActions" :key="idx">
                <el-button
                    :type="action.type"
                    :icon="action.icon"
                    plain
                    class="quick-action-btn"
                    style="width: 100%;"
                >
                  {{ action.label }}
                </el-button>
              </el-col>
            </el-row>
          </div>
        </el-card>
      </el-col>
      <el-col :span="16">
        <el-card class="dashboard-chart-card">
          <template #header>访问趋势</template>
          <div class="chart-wrapper large-chart">
            <v-chart v-if="chartOption" :option="chartOption" autoresize style="width:100%;height:100%;" />
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup>
import { ElIcon } from 'element-plus';
import { User, Document, Setting, Warning, UserFilled } from '@element-plus/icons-vue';
import { ref } from 'vue';
import VChart from 'vue-echarts';

const cards = [
  { title: '用户数', value: 1024, desc: '较昨日 +10', icon: User, color: '#409EFF' },
  { title: '访问量', value: 23000, desc: '较昨日 +500', icon: Document, color: '#67C23A' },
  { title: '活跃用户', value: 800, desc: '较昨日 +20', icon: User, color: '#909399' },
  { title: '待处理事项', value: 12, desc: '较昨日 +2', icon: Warning, color: '#E6A23C' }
];

const websiteInfo = [
  { label: '网站名称', value: 'Gin-Base后台管理系统' },
  { label: '当前版本', value: 'v1.0.0' },
  { label: '管理员', value: '张三' },
  { label: '联系方式', value: 'dsx.email@qq.com' },
  { label: '运行环境', value: '生产环境' }
];

const quickActions = [
  { label: '新增用户', type: 'primary', icon: 'el-icon-plus' },
  { label: '菜单管理', type: 'success', icon: 'el-icon-document' },
  { label: '用户管理', type: 'warning', icon: 'el-icon-setting' },
  { label: '角色管理', type: 'info', icon: 'el-icon-user' },
  { label: '字典管理', type: 'danger', icon: 'el-icon-warning' }
];

const chartOption = ref({
  tooltip: { trigger: 'axis' },
  xAxis: { type: 'category', data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'] },
  yAxis: { type: 'value' },
  series: [
    { data: [120, 200, 150, 80, 70, 110, 130], type: 'line', smooth: true }
  ]
});
</script>

<style scoped>
.dashboard-container {
  padding-top: 32px;
  padding-left: 16px;
  padding-right: 16px;
  padding-bottom: 0;
  min-height: 100vh;
  background: #f5f6fa;
}
.common-row-gap {
  margin-top: 0;
  margin-bottom: 0;
  padding-top: 10px;
  padding-bottom: 10px;
}
.dashboard-cards {
  /* 其他样式可保留 */
}
.dashboard-card {
  margin-top: 0;
}
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 16px;
  margin-bottom: 10px;
}
.card-value {
  font-size: 28px;
  font-weight: bold;
  margin-bottom: 5px;
}
.card-desc {
  color: #909399;
  font-size: 14px;
}
.website-info-row {
}
.website-info-card {
}
.website-info-title {
  font-size: 18px;
  font-weight: bold;
  margin-bottom: 8px;
}
.website-info-table {
  margin-top: 8px;
  border-radius: 6px;
  overflow: hidden;
}
.dashboard-main {
  min-height: 420px;
}
.dashboard-chart-card,
.dashboard-quick-card {
  height: 420px;
  display: flex;
  flex-direction: column;
}
.chart-wrapper {
  flex: 1;
  min-height: 0;
  height: 100%;
  width: 100%;
  display: flex;
}
.large-chart {
  height: 360px !important;
}
.quick-actions {
  flex: 1;
  margin-top: 8px;
}
.quick-action-btn {
  width: 100%;
  margin-bottom: 12px;
  box-sizing: border-box;
}
</style>