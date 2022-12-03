<template>
  <el-row :gutter="20" class="main" justify="center">
    <el-col :span="6" class="main-item-1">
      <el-row>
        <ExtantDiagnosis></ExtantDiagnosis>
      </el-row>
      <el-row>
        <Added></Added>
      </el-row>
      <el-row>
        <Cumulative></Cumulative>
      </el-row>
    </el-col>
    <el-col :span="12" class="main-item-2">
      <el-col class="m-1">
        <Statistics></Statistics>
      </el-col>
      <el-col class="m-2">
        <Map></Map>
      </el-col>
      <el-col class="m-3">
        <HotMsg></HotMsg>
      </el-col>
    </el-col>
    <el-col :span="6" class="main-item-3">
      <el-row class="r-1">
        <Outbreak></Outbreak>
      </el-row>
      <el-row class="r-2">
        <DangerousAreas></DangerousAreas>
      </el-row>
    </el-col>
  </el-row>
</template>

<script lang="ts" setup>
import ExtantDiagnosis from "./content/extant_diagnosis.vue"
import Added from "./content/added.vue"
import Cumulative from "./content/cumulative.vue"
import Outbreak from "./content/outbreak.vue"
import DangerousAreas from "./content/dangerous_areas.vue"
import Statistics from "./content/statistics.vue"
import Map from "./content/map.vue"
import HotMsg from "./content/hot_msg.vue"</script>

<style lang="scss" scoped>
.main {
  padding: 20px;

  &-item-1 {

    * {
      width: 450px;
      height: 308px;
    }

    *:not(:last-child) {
      margin-bottom: 10px;
    }
  }

  &-item-2 {

    .m-1 {
      width: 950px;
      height: 120px;
      margin-bottom: 10px;
    }

    .m-2 {
      width: 950px;
      height: 570px;
      margin-bottom: 10px;
    }

    .m-3 {
      width: 950px;
      height: 230px;
    }
  }

  &-item-3 {

    .r-1 {
      width: 450px;
      height: 365px;
      margin-bottom: 10px;
    }

    .r-2 {
      width: 450px;
      height: 570px;
    }
  }
}
</style>