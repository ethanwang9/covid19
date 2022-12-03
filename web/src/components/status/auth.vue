<template>
  <el-row class="error">
    <el-col class="error-icon">
      <i-ep-CircleCloseFilled class="error-icon-size"/>
    </el-col>
    <el-col class="error-text">
      <p>权限不足无法查看</p>
    </el-col>
  </el-row>
</template>

<script lang="ts" setup>

</script>

<style lang="scss" scoped>
.error {
  display: flex;
  justify-content: center;
  flex-direction: column;
  overflow: hidden;

  &-icon {
    padding: 50px 0 0;
    text-align: center;

    &-size {
      font-size: 60px;
      color: red;
      margin: 0;
      padding: 0;
    }
  }

  &-text {
    font-size: 20px;
    font-weight: bold;
    text-align: center;
    padding: 20px 0 40px;
    letter-spacing: 2px;
  }
}
</style>