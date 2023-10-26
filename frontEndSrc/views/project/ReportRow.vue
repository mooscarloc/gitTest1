<script>
export default {
  props: {
    metricId: String,
    metricVal: Number,
    deviceName: String
  },
  filters: {
    getEnterpriseStandard(deviceCode, arg) {
      if(!deviceCode) {
          return;
      }

      const codeMap = {
          AQS:{
              value: "aqsValue",
              zeroNoise: "<=0.2ppb",
              scopeNoise: "<=1.0ppb",
              linear:    "1.0%F.S",
              Precision20: "<=1.0ppb",
              Precision80: "<=2.0ppb",
              upTime: "<=120S",
              downTime: "<=120S",
          },  
          AQC:{
              value: "aqcValue",
              zeroNoise: "<=0.5ppm",
              scopeNoise: "<=1.0ppm",
              linear:    "1.0%F.S",
              Precision20: "<=1.0ppm",
              Precision80: "<=2.0ppm",
              upTime: "<=180S",
              downTime: "<=180S",
          },
      };

      return codeMap[deviceCode.substring(0,3)][arg];
    },
    metStandard(deviceCode, arg, val) {
      if(!deviceCode) {
          return;
      }
      const targetMap = {
        AQS:{
          zeroNoise: 0.2,
          scopeNoise: 1.0,
          upTime: 0.5,
          downTime: 0.7
        },
        AQC:{
          zeroNoise: 0.04,
          scopeNoise: 2.0,    
          upTime: 0.5,
          downTime: 0.7   
        }
      };
      return val <= targetMap[deviceCode.substring(0,3)][arg]  ? "Pass" : "Fail";
    },
    getMetricName(id) {
        const dict = {
            zeroNoise: "Zero Noise",
            scopeNoise: "Scope Noise",
            upTime: "Rise Time",
            downTime: "Fall Time",
        };
        return dict[id];
    }
  },
}
</script>

<template>
  <el-row>
    <el-col span='320'>
        <el-form-item :label="metricId | getMetricName">
        <span>{{ metricVal }}</span>
        </el-form-item>
    </el-col>
    <el-col span='320'>   
        <el-form-item label="Enterprise Standard">
        <span>{{ deviceName | getEnterpriseStandard(metricId) }}</span>
        </el-form-item>
    </el-col>
    <el-col span='400'>
        <el-form-item label="Test Result">
        <span>{{ deviceName | metStandard(metricId, metricVal) }}</span>
        </el-form-item>
    </el-col>
  </el-row>
</template>