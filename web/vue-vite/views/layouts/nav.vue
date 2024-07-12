<script setup>
import {Functions} from '@utils/functions/functions';

let funcs = new Functions();

</script>
<template>
  <!-- Navbar -->
  <nav class="main-header navbar navbar-expand navbar-white navbar-light">
    <!-- Left navbar links -->
    <ul class="navbar-nav">
      <li class="nav-item">
        <a class="nav-link" data-widget="pushmenu" href="javascript:void(0);" role="button"><i class="fas fa-bars"></i></a>
      </li>
      <li class="nav-item d-none d-sm-inline-block">
        <a href="/" class="nav-link" target="_blank">{{ funcs.lang('Home') }}</a>
      </li>
    </ul>

    <!-- Right navbar links -->
    <ul class="navbar-nav ml-auto">
      <!-- Navbar Search -->
      <li class="nav-item">
        <a class="nav-link" data-widget="navbar-search" href="javascript:void(0);" role="button">
          <i class="fas fa-search"></i>
        </a>
        <div class="navbar-search-block">
          <form class="form-inline">
            <div class="input-group input-group-sm">
              <input class="form-control form-control-navbar" type="search" :placeholder="funcs.lang('Search')"
                     aria-label="Search">
              <div class="input-group-append">
                <button class="btn btn-navbar" type="submit">
                  <i class="fas fa-search"></i>
                </button>
                <button class="btn btn-navbar" type="button" data-widget="navbar-search">
                  <i class="fas fa-times"></i>
                </button>
              </div>
            </div>
          </form>
        </div>
      </li>

      <!-- Notifications Dropdown Menu -->
      <li class="nav-item dropdown">
        <a class="nav-link" data-toggle="dropdown" href="javascript:void(0);">
          <i class="far fa-bell"></i>
          <span class="badge badge-warning navbar-badge">15</span>
        </a>
        <div class="dropdown-menu dropdown-menu-lg dropdown-menu-right">
          <span class="dropdown-item dropdown-header">15 Notifications</span>
          <div class="dropdown-divider"></div>
          <a href="javascript:void(0);" class="dropdown-item">
            <i class="fas fa-envelope mr-2"></i> 4 new messages
            <span class="float-right text-muted text-sm">3 mins</span>
          </a>
          <div class="dropdown-divider"></div>
          <a href="javascript:void(0);" class="dropdown-item">
            <i class="fas fa-users mr-2"></i> 8 friend requests
            <span class="float-right text-muted text-sm">12 hours</span>
          </a>
          <div class="dropdown-divider"></div>
          <a href="javascript:void(0);" class="dropdown-item">
            <i class="fas fa-file mr-2"></i> 3 new reports
            <span class="float-right text-muted text-sm">2 days</span>
          </a>
          <div class="dropdown-divider"></div>
          <a href="javascript:void(0);" class="dropdown-item dropdown-footer">See All Notifications</a>
        </div>
      </li>
      <!-- Language Dropdown Menu -->
      <li class="nav-item dropdown">
        <a class="nav-link" data-toggle="dropdown" href="javascript:void(0);">
          <i class="flag-icon flag-icon-cn"></i>
        </a>
        <div class="dropdown-menu dropdown-menu-right p-0">
          <a @click="funcs.switchLang('zh-cn')" href="javascript:void(0);" class="dropdown-item active">
            <i class="flag-icon flag-icon-cn mr-2"></i> {{ funcs.lang('Chinese') }}
          </a>
          <a @click="funcs.switchLang('en-us')" href="javascript:void(0);" class="dropdown-item">
            <i class="flag-icon flag-icon-us mr-2"></i> {{ funcs.lang('English') }}
          </a>
        </div>
      </li>
      <li class="nav-item dropdown user-menu">
        <a href="javascript:void(0);" class="nav-link dropdown-toggle" data-toggle="dropdown">
          <img src="/images/user_avatar.jpg" class="user-image img-circle elevation-2" alt="User Image">
          <span class="d-none d-md-inline">{{ funcs.lang('Administrator') }}</span>
        </a>
        <ul class="dropdown-menu dropdown-menu-lg dropdown-menu-right">
          <!-- User image -->
          <li class="user-header bg-primary">
            <img src="/images/user_avatar.jpg" class="img-circle elevation-2" alt="User Image">

            <p>
              {{ funcs.lang('Administrator') }}
              <small>七天下</small>
            </p>
          </li>
          <li class="user-footer">
            <a href="javascript:void(0);" class="btn btn-default btn-flat">{{ funcs.lang('UserInfo') }}</a>
            <a href="javascript:void(0);" class="btn btn-default btn-flat float-right">{{ funcs.lang('Logout') }}</a>
          </li>
        </ul>
      </li>
      <li class="nav-item">
        <a class="nav-link" data-widget="fullscreen" href="javascript:void(0);" role="button">
          <i class="fas fa-expand-arrows-alt"></i>
        </a>
      </li>
    </ul>
  </nav>
  <!-- /.navbar -->
</template>