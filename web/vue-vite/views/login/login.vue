<template>
  <div class="login-container flx-center">
    <div class="login-box">
      <el-switch class="dark" v-model="isDark" inline-prompt :active-icon="Sunny" :inactive-icon="Moon" @change="switchDark" />
      <div class="login-left">
        <img class="login-left-img" src="@/styles/assets/images/login_left.png" alt="login" />
      </div>
      <div class="login-form">
        <div class="login-logo">
          <img class="login-icon" src="/favicon.ico" alt="" />
          <h2 class="logo-text">{{ funcs.lang('Manage Backend') }}</h2>
        </div>
        <el-form ref="loginFormRef" :model="loginForm" :rules="loginRules" size="large">
          <el-form-item prop="username">
            <el-input v-model="loginForm.username" :placeholder="funcs.lang('Home')+': admin'">
              <template #prefix>
                <el-icon class="el-input__icon">
                  <user />
                </el-icon>
              </template>
            </el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input v-model="loginForm.password" type="password" :placeholder="funcs.lang('Password')+': 123456'" show-password autocomplete="new-password">
              <template #prefix>
                <el-icon class="el-input__icon">
                  <lock />
                </el-icon>
              </template>
            </el-input>
          </el-form-item>
        </el-form>
        <div class="login-btn">
          <el-button :icon="CircleClose" round size="large" @click="resetForm(loginFormRef)"> {{ funcs.lang('Reset') }} </el-button>
          <el-button :icon="UserFilled" round size="large" type="primary" class="login-bt" @click="login(loginFormRef)">
            {{ funcs.lang('Login') }}
          </el-button>
        </div>
      </div>
    </div>
  </div>
</template>
<script setup>
import { ref, onMounted, reactive, onBeforeUnmount } from 'vue';
import { CircleClose, UserFilled, Sunny, Moon } from '@element-plus/icons-vue';
import { useRoute, useRouter } from 'vue-router';
import loginModule from '@/app/modules/admin/login';
import createService from '@/utils/service';
import Functions from '@/utils/functions';
import pnotify from '@/utils/pnotify/alert';
import {HOME_URL} from "@/config";

const funcs = new Functions();
const route = useRoute();
const router = useRouter();
const loginService = createService(loginModule);
const isDark = ref(false);
const loginForm = reactive({
  username: "",
  password: ""
});
const loginFormRef = ref();
const loginRules = reactive({
  username: [{ required: true, message: funcs.lang('Please Entry Username'), trigger: "blur" }],
  password: [{ required: true, message: funcs.lang('Please Entry Password'), trigger: "blur" }]
});
const switchDark = () => {
  const htmlElement = document.querySelector('html');
  if (isDark.value) {
    htmlElement.classList.remove('light');
    htmlElement.classList.add('dark');
  } else {
    htmlElement.classList.remove('dark');
    htmlElement.classList.add('light');
  }
};
const resetForm = (formEl) => {
  if (!formEl) {
    return;
  }

  formEl.resetFields();
};
onMounted(() => {
  // 监听enter事件调用登录
  document.onkeydown = (e) => {
    if (e.code.toLowerCase() === 'enter' || e.code === 'NumpadEnter') {
      login(loginFormRef.value);
    }
  };
});
onBeforeUnmount(() => {
  document.onkeydown = null;
});
const login = async (formEl) => {
  if (!formEl) {
    return;
  }

  formEl.validate(async valid => {
    if (!valid) {
      return;
    }
  });

  const result = await loginService.login(loginForm);

  funcs.setLogin(result.data);

  pnotify(funcs.lang('Successfully logged in'));

  setTimeout(function () {
    // 验证路径是否有redirect
    if (route.query.redirect) {
      router.push(route.query.redirect);
    } else {
      router.push(HOME_URL + '?lang=' + funcs.getLang());
    }

    return;
  }, 100);
}
</script>
<style lang="scss" scoped>
@import "./index.scss";
</style>