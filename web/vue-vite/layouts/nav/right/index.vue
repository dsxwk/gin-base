<template>
  <div class="header-ri">
    <div class="tool-bar-ri">
      <div class="header-icon">
        <Language id="language" />
        <SearchMenu id="searchMenu" />
        <Fullscreen id="fullscreen" />
      </div>
      <span class="username">超级管理员</span>
      <el-dropdown trigger="click">
        <div class="avatar">
          <img src="/images/avatar.jpg" alt="avatar" />
        </div>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item @click="openDialog('infoRef')">
              <el-icon><User /></el-icon>个人信息
            </el-dropdown-item>
            <el-dropdown-item @click="openDialog('passwordRef')">
              <el-icon><Edit /></el-icon>修改密码
            </el-dropdown-item>
            <el-dropdown-item divided @click="logout">
              <el-icon><SwitchButton /></el-icon>退出登录
            </el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
  </div>
</template>
<script setup>
import Fullscreen from '@/layouts/nav/right/fullScreen/index.vue';
import Language from '@/layouts/nav/right/language/index.vue';
import SearchMenu from '@/layouts/nav/right/searchMenu/index.vue';
import Functions from '@/utils/functions';

const funcs = new Functions();
const openDialog = (name) => {

}
const logout = () => {
  funcs.logout();
}
</script>
<style scoped lang="scss">
.avatar {
  width: 40px;
  height: 40px;
  overflow: hidden;
  cursor: pointer;
  border-radius: 50%;
  img {
    width: 100%;
    height: 100%;
  }
}
.tool-bar-ri {
  display: flex;
  align-items: center;
  justify-content: center;
  padding-right: 25px;
  .header-icon {
    display: flex;
    align-items: center;
    & > * {
      margin-left: 21px;
      color: var(--el-header-text-color);
    }
  }
  .username {
    margin: 0 20px;
    font-size: 15px;
    color: var(--el-header-text-color);
  }
}
</style>