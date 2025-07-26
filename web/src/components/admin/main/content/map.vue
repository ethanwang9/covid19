<template>
  <div class="MMain">
    <div id="MMap" ref="MMap"></div>
  </div>
</template>

<script lang="ts" setup>
import * as echarts from 'echarts'
import {EChartOption} from 'echarts'
import useStore from "../../../../store";
import {GeoMap, GeoMapData} from "../../../../api/panel";

// store
const store = useStore().app

// 地图
const MMap = ref()

// 设置echarts
const initMap = async () => {
  // 初始化地图
  const map = echarts.init(MMap.value)
  map.showLoading()

  // 获取Geo地理位置信息
  await GeoMap({
    province: store.userInfo.location
  }).then(res => {
    echarts.registerMap("map", res.data)
  }).catch(err => {
    ElNotification({
      title: "错误",
      message: err.message,
    })
  })

  // 获取地图数据
  let d: any = null
  let center: number[] = []
  await GeoMapData({
    province: store.userInfo.location
  }).then(res => {
    d = res.data.value
    center = res.data.center
  }).catch(err => {
    ElNotification({
      title: "错误",
      message: err.message,
    })
  })

  // 设置属性
  const option: EChartOption = {
    tooltip: {
      show: true,
      formatter: function (d) {
        if (d["data"].value === -99999999) {
          let s: string = "该地区暂未公布"
          return (
              `${d["name"]}<br>确诊人数: ${s}<br>高风险地区: ${d["data"].h}<br>中风险地区: ${d["data"].m}<br>低风险地区: ${d["data"].l}<br>`
          )
        }
        return (
            `${d["name"]}<br>确诊人数: ${d["data"].value}<br>高风险地区: ${d["data"].h}<br>中风险地区: ${d["data"].m}<br>低风险地区: ${d["data"].l}<br>`
        )
      }
    },
    geo: {
      type: 'map',
      map: 'map',
      zoom: 6,
      roam: true,
      center: center,
      label: {
        show: true,
        fontSize: 16,
      },
    },
    series: [
      {
        geoIndex: 0,
        type: 'map',
        data: d,
      },
    ],
    visualMap: {
      show: true,
    },
  }

  map.setOption(option)
  map.hideLoading()
}

// 设置图表
onMounted(() => {
  initMap()
})
</script>

<style lang="scss" scoped>
.MMain, #MMap {
  width: 900px;
  height: 570px;
  //background: darkgreen;
}
</style>
