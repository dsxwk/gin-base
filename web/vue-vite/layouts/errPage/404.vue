<script setup>
import { onMounted } from 'vue';

onMounted(() => {
  let str = document.getElementsByTagName('div')[0].innerHTML.toString();
  let i = 0;
  document.getElementsByTagName('div')[0].innerHTML = "";

  setTimeout(function() {
    let se = setInterval(function() {
      i++;
      document.getElementsByTagName('div')[0].innerHTML = str.slice(0, i) + "|";
      if (i == str.length) {
        clearInterval(se);
        document.getElementsByTagName('div')[0].innerHTML = str;
      }
    }, 10);
  },0);
});
</script>
<template>
  <div class="page-404">
    <div class="bg">404</div>
    <div>
      <p>> <span>Error Code</span>: "<i>HTTP 404 Not Found</i>"</p>
      <p>> <span>错误描述</span>: "<i>服务器上你访问的页面不存在！</i>"</p>
      <p>> <span>可能由以下原因引起的错误</span>: [<b>你可能访问了错误的连接地址,请检查你访问的地址是否正确！</b>...]</p>
      <p>> <span>页面走丢了,请打开正确的页面重试！</span>: [<a href="/" target="_blank">首页</a>, <a href="/" target="_blank">联系我们</a>, <a href="/" target="_blank">Dashboard</a>...]</p><p>> <span>祝你愉快！ :-)</span></p>
    </div>
  </div>
</template>
<style lang="less">
* {
  margin: 0;
  padding: 0;
}
.page-404{
  background-image: linear-gradient(120deg, #4f0088 0%, #000000 100%);
  height: 100vh;
  user-select: none;
  .bg {
    font-size: 25vw;
    text-align: center;
    position: fixed;
    width: 100%;
    z-index: 1;
    color: #ffffff26;
    text-shadow: 0 0 50px rgba(0, 0, 0, 0.07);
    top: 50%;
    -webkit-transform: translateY(-50%);
    transform: translateY(-50%);
    font-family: "Montserrat", monospace;
    font-weight: bold;
  }

  div {
    background: rgba(0, 0, 0, 0);
    width: 70vw;
    position: relative;
    top: 50%;
    -webkit-transform: translateY(-50%);
    transform: translateY(-50%);
    margin: 0 auto;
    padding: 30px 30px 10px;
    box-shadow: 0 0 150px -20px rgba(0, 0, 0, 0.5);
    z-index: 3;
  }

  P {
    font-family: "Share Tech Mono", monospace;
    color: #f5f5f5;
    margin: 0 0 20px;
    font-size: 17px;
    line-height: 1.2;
  }

  span {
    color: #f0c674;
  }

  i {
    color: #8abeb7;
  }

  div a {
    text-decoration: none;
  }

  b {
    color: #81a2be;
  }

  a.avatar {
    position: fixed;
    bottom: 15px;
    right: -100px;
    -webkit-animation: slide 0.5s 4.5s forwards;
    animation: slide 0.5s 4.5s forwards;
    display: block;
    z-index: 4
  }

  a.avatar img {
    border-radius: 100%;
    width: 44px;
    border: 2px solid white;
  }
}
@-webkit-keyframes slide {
  from {
    right: -100px;
    -webkit-transform: rotate(360deg);
    transform: rotate(360deg);
    opacity: 0;
  }
  to {
    right: 15px;
    -webkit-transform: rotate(0deg);
    transform: rotate(0deg);
    opacity: 1;
  }
}

@keyframes slide {
  from {
    right: -100px;
    -webkit-transform: rotate(360deg);
    transform: rotate(360deg);
    opacity: 0;
  }
  to {
    right: 15px;
    -webkit-transform: rotate(0deg);
    transform: rotate(0deg);
    opacity: 1;
  }
}
</style>