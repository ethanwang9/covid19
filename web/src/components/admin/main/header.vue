<template>
  <el-row class="main">
    <el-col :span="7" class="main-left">
      <div class="main-left-line">
        <dv-decoration3/>
      </div>
      <div class="main-left-btn">
        <el-button class="main-left-btn-item" size="large" text @click="changeArea">切换区域</el-button>
        <el-button class="main-left-btn-item" size="large" text @click="fullScreen">
          {{ isFull ? "退出全屏" : "全屏" }}
        </el-button>
        <el-button class="main-left-btn-item" size="large" text @click="showMenu">
          {{ isShowMenu ? "隐藏菜单" : "显示菜单" }}
        </el-button>
      </div>
    </el-col>
    <el-col :span="10" class="main-title">
      <h2 class="main-title-text">COVID19 大数据可视化系统</h2>
      <dv-decoration5 :dur="5" class="main-title-line"/>
    </el-col>
    <el-col :span="7" class="main-right">
      <dv-decoration3/>
      <div class="main-right-item">
        <p>归属地: {{ location }}</p>
        <p>{{ datetime }}</p>
      </div>
    </el-col>
  </el-row>
</template>

<script lang="ts" setup>
import screenFull from 'screenfull';
import useStore from "../../../store";
import {dateTime} from '../../../utils/tool'
import {ref} from "vue";

// store
const store = useStore().admin

// emit 事件
const emit = defineEmits([
  'changeArea',
])

// 归属地
const {location} = toRefs(useStore().app.userInfo)

// 是否全屏
const isFull = ref<boolean>(false)

// 全屏查看
function fullScreen() {
  screenFull.toggle();
  isFull.value = !isFull.value
}

// 是否显示菜单
const isShowMenu = ref<boolean>(false)

// 显示菜单
function showMenu() {
  isShowMenu.value = !isShowMenu.value
  if (isShowMenu.value) {
    store.changeShowHeader(true, true)
  } else {
    store.changeShowHeader(false)
  }
}

// 切换区域
function changeArea() {
  emit('changeArea')
}

// 当前时间
const datetime = ref<string>("")
const timer = setInterval(() => {
  let t = new Date().getTime()
  datetime.value = dateTime(t)
}, 1000)
onBeforeUnmount(() => {
  clearInterval(timer)
})
</script>

<style lang="scss" scoped>
.main {
  width: 1920px;
  height: 100px;

  &-left {
    display: flex;
    align-items: center;
    flex-direction: column;
    font-size: 22px;

    &-line {
      width: 100%;
    }

    &-btn {
      width: 100%;
      display: flex;
      justify-content: center;
      align-items: center;

      &-item {
        font-size: 22px;
        margin: 10px 50px 0 0;
        letter-spacing: 2px;
      }
    }
  }

  &-title {
    display: flex;
    align-items: center;
    flex-direction: column;
    position: relative;

    &-text {
      font-size: 28px;
      letter-spacing: 6px;
      padding: 0;
      margin: 0;
      position: absolute;
      top: 20px;
    }

    &-line {
      height: 80px;
      position: absolute;
      top: 30px;
    }
  }

  &-right {
    display: flex;
    align-items: center;
    flex-direction: column;
    font-size: 22px;

    &-item {
      display: flex;

      *:first-child {
        margin-right: 15px;
        font-weight: bold;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
        width: 290px;
      }
    }
  }
}
</style>