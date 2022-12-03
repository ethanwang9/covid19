<template>
  <div class="CUMMain">
    <dv-border-box12>
      <div class="CUMMain-title">
        <p>累计</p>
        <dv-decoration1 class="CUMMain-title-line"/>
      </div>
      <div ref="cEcharts" class="CUMMain-echarts"></div>
    </dv-border-box12>
  </div>
</template>

<script lang="ts" setup>
import * as echarts from 'echarts'
import {EChartOption} from 'echarts'
import roma from '../../../../assets/theme/echarts.json'
import useStore from "../../../../store";
import {HistoryCurrConf} from "../../../../api/panel";
// 获取DOM
const cEcharts = ref()

// store
const store = useStore().app

// 获取历史数据
const GetHistory = async (fun: Function) => {
  await HistoryCurrConf({
    province: store.userInfo.location
  }).then(res => {
    let x: string[] = []
    let y1: number[] = []
    let y2: number[] = []
    let y3: number[] = []
    res.data.forEach(v => {
      let a = v.dateId.toString().substring(0, 4)
      let b = v.dateId.toString().substring(4, 6)
      let c = v.dateId.toString().substring(6, 8)
      x.push(`${a}-${b}-${c}`)
      y1.push(v.confirmedCount)
      y2.push(v.deadCount)
      y3.push(v.curedCount)
    })
    return fun(x, y1, y2, y3)
  }).catch(err => {
    ElNotification({
      title: "错误",
      message: err.message,
    })
  })
}

// 设置图表
onMounted(() => {
  const charts = echarts.init(cEcharts.value, roma)
  charts.showLoading()
  GetHistory((x: any, y1: any, y2: any, y3: any) => {
    const options: EChartOption = {
      tooltip: {
        trigger: 'axis',
        axisPointer: {
          type: 'cross',
        },
      },
      toolbox: {
        show: true,
        feature: {
          saveAsImage: {show: true}
        }
      },
      legend: {show: true},
      grid: {
        left: 55,
        x2: 10,
        y1: 10,
        y2: 70,
        height: 140,
      },
      xAxis: {
        type: 'category',
        data: x,
        axisLabel: {
          interval: 180,
        }
      },
      yAxis: {
        type: 'value'
      },
      series: [
        {
          name: '累计确诊',
          data: y1,
          type: 'line',
          smooth: true,
          areaStyle: {},
        },
        {
          name: '累计死亡',
          data: y2,
          type: 'line',
          smooth: true,
          areaStyle: {},
        },
        {
          name: '累计治愈',
          data: y3,
          type: 'line',
          smooth: true,
          areaStyle: {},
        }
      ],
      dataZoom: [
        {
          type: 'slider'
        }
      ]
    }
    charts.setOption(options)
    charts.hideLoading()
  })
})
</script>

<style lang="scss" scoped>
.CUMMain {
  width: 450px;
  height: 308px;

  &-title {
    display: flex;
    align-items: center;
    vertical-align: middle;

    p {
      color: #a5d5f6;
      font-weight: bold;
      font-size: 20px;
      margin: 18px 0 0 25px;
      text-align: left;
    }

    &-line {
      width: 200px;
      height: 24px;
      padding: 6px 0 0 4px;
    }
  }

  &-echarts {
    width: 430px;
    height: 252px;
    margin-left: 10px;
    margin-top: 6px;
  }
}

</style>