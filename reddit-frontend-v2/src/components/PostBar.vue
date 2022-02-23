<template>
  <div class="height"></div>
    <div class="app-bar">
      <div class="right-links ">
        <a href="http://localhost:8080/login-register" class="app-bar-item" >LOGIN</a>
      </div>
    </div>
</template>

<script>
export default {
  name: "PostBar"
}
</script>

<style scoped>
.height {
  height: 40px;
}
.app-bar {
  height: 45px;
  width: 100%;
  position: fixed;
  z-index: 1;
  top: 0;
  left: 0;
  background-color: burlywood;
  overflow: hidden;
  display: block;
  justify-content: space-between;
}
.left-links {
  padding-left: 10%;
}
.left-links a {
  float: left;
}
a {
  float: left;
  display: block;
  color: #f2f2f2;
  text-align: center;
  padding: 14px 16px;
  text-decoration: none;
  font-size: 17px;
}
.right-links {
  float: right;
}
</style>