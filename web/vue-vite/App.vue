<script setup>
import {onMounted} from "vue";
import $ from 'jquery';
import {Functions} from "@utils/functions/functions";
document.title = '后台管理';

let funcs = new Functions();

onMounted(async function () {
  setTimeout(updateMenuActive, 100);
});

// 更新菜单选中
function updateMenuActive() {
  $('body a').each(function () {
    //设置所有连接带语言
    $(this).attr('href', funcs.setUrlParam($(this).attr('href'), 'lang', funcs.getLang()));
  });

  let sidebarMenu = $('#SidebarMenu');
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
  <router-view></router-view>
</template>
