<script setup>
import {Functions} from '@utils/functions/functions';
import {onMounted, nextTick, ref} from "vue";

let funcs = new Functions();

// 菜单
const menu = ref([
  {
    name: funcs.lang('Dashboard'), uri: 'javascript:void(0);', icon: 'fas fa-tachometer-alt', path:'',
    children: [
      {name: funcs.lang('Dashboard') + ' v1', uri: '/', icon: 'far fa-circle', path:''},
    ]
  },
  {
    name: funcs.lang('User Manage'), uri: 'javascript:void(0);', icon: 'far fa-circle', path:'',
    children: [
      {name: funcs.lang('User List'), uri: '/user', icon: 'far fa-circle', path:''},
    ]
  }
]);

onMounted(async () => {
  await nextTick();
  updateMenuActive();
});

// 更新菜单选中
async function updateMenuActive() {
  $('body a').each(function () {
    //设置所有连接带语言
    $(this).attr('href', funcs.setUrlParam($(this).attr('href'), 'lang', funcs.getLang()));
  });

  let sidebarMenu = $('#SidebarMenu');
  console.log(sidebarMenu);
  let url = window.location;
  url = url.pathname + url.hash;
  if (url.indexOf('?') == -1) {
    url = url + '?lang=' + funcs.getLang();
  }

  sidebarMenu.find('a[href="' + url + '"]').addClass('active');
  sidebarMenu.find('a[href="' + url + '"]').parent('li').parent().parent().addClass('menu-open');
  sidebarMenu.find('a[href="' + url + '"]').parent().parent().prev().addClass('active');
}
</script>
<template>
  <!-- Main Sidebar Container -->
  <aside class="main-sidebar sidebar-dark-primary elevation-4">
    <!-- Brand Logo -->
    <a href="/" class="brand-link">
      <img src="/favicon.ico" alt="AdminLTE Logo" class="brand-image img-circle elevation-3" style="opacity: .8;width: 30px;height: 30px;">
      <span class="brand-text font-weight-light">{{funcs.lang('Manage Backend')}}</span>
    </a>

    <!-- Sidebar -->
    <div class="sidebar">
      <!-- SidebarSearch Form -->
      <div class="form-inline">
        <div class="input-group" data-widget="sidebar-search">
          <input class="form-control form-control-sidebar" type="search" :placeholder="funcs.lang('Search')" aria-label="Search">
          <div class="input-group-append">
            <button class="btn btn-sidebar">
              <i class="fas fa-search fa-fw"></i>
            </button>
          </div>
        </div>
      </div>

      <!-- Sidebar Menu -->
      <nav class="mt-2">
        <ul id="SidebarMenu" class="nav nav-pills nav-sidebar flex-column" data-widget="treeview" role="menu" data-accordion="false">
          <!-- Add icons to the links using the .nav-icon class
               with font-awesome or any other icon font library -->
          <li v-for="(item, key) in menu" class="nav-item">
            <a :href="item.uri" class="nav-link">
              <i class="nav-icon" :class="item.icon"></i>
              <p>
                {{item.name}}
                <i class="right fas fa-angle-left"></i>
              </p>
            </a>
            <ul class="nav nav-treeview">
              <li v-for="(i, k) in item.children" class="nav-item">
                <a :href="i.uri" class="nav-link">
                  <i class="nav-icon" :class="i.icon"></i>
                  <p>{{i.name}}</p>
                </a>
              </li>
            </ul>
          </li>
        </ul>
      </nav>
      <!-- /.sidebar-menu -->
    </div>
    <!-- /.sidebar -->
  </aside>
</template>