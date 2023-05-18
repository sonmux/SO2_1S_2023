<script setup>
import {reactive} from 'vue'
import {Greet} from '../../wailsjs/go/main/App'


import {GetCPUUsage, GetDISKUsage} from '../../wailsjs/go/sys/Stats'

/*const data = reactive({
  name: "",
  resultText: "Please enter your name below 👇",
})

function greet() {
  Greet(data.name).then(result => {
    data.resultText = result
  })
}*/

</script>

<template>
    <!---<div id="result" class="result">{{ message }}</div>
    <div id="input" class="input-box">
      
    </div>
    <h1>{{message}}</h1>-->
      <!---<input id="name" v-model="data.name" autocomplete="off" class="input" type="text"/>
      <button class="btn" @click="getMessage">Press Me!</button>-->
  
  <div class="CPUUsage">
    <h1>CPU Usage: {{message}}%</h1>
    <apexchart type="radialBar" width="520" :options="options" :series="series"></apexchart>
  </div>
  <div class="DISKUsage">
    <h1>Disk Usage: {{DISKUsageMessage}}% and Disk Free: {{DISKFreeMessage}}%</h1>
    <apexchart type="donut" width="480" :options="diskoptions" :series="diskseries"></apexchart>
  </div>
</template>

<script>
  export default{
    data(){
      return{
        series: [0],
        options: {
          labels: ["CPU Usage"],
          title: {
            text: "CPU Usage",
            align: "center"
          }
        },
        message:"DataCPU",
        diskseries: [0,100],
        diskoptions: {
          labels: ["Used","Free"],
          title: {
            text: "Disk",
            align: "center"
          }
        },
        DISKUsageMessage:"DataDiskUsed",
        DISKFreeMessage:"DataDiskFree",
      };
    },
    /*methods: {
      getMessage: function() {
        var self = this;
        GetCPUUsage().then(cpu_usage => {
          //console.log(cpu_usage);
          self.message = cpu_usage.avg;
        });
      }
    }*/
    mounted: function() {
      GetCPUUsage().then(cpu_usage => {
        //this.message = cpu_usage.avg;
        if (cpu_usage) {
          this.series = [ cpu_usage.avg ];
          this.message = cpu_usage.avg;
        }
      })
      GetDISKUsage().then(disk_usage => {
        //this.message = disk_usage.avg;
        if (disk_usage) {
          this.diskseries = [ disk_usage.used, disk_usage.free ];
          this.DISKUsageMessage = ((disk_usage.used*100)/(disk_usage.used+disk_usage.free)).toFixed(2);;
          this.DISKFreeMessage = ((disk_usage.free*100)/(disk_usage.used+disk_usage.free)).toFixed(2);
        }
      })
      this.updateData();
    },
    methods: {
      updateData(){
        setInterval(()=>{
          GetCPUUsage().then(cpu_usage => {
            //this.message = cpu_usage.avg;
            if (cpu_usage) {
              this.series = [ cpu_usage.avg ];
              this.message = cpu_usage.avg;
            }
          })
        },1000);
        setInterval(()=>{
          GetDISKUsage().then(disk_usage => {
            //this.message = disk_usage.avg;
            if (disk_usage) {
              this.diskseries = [ disk_usage.used, disk_usage.free ];
              this.DISKUsageMessage = ((disk_usage.used*100)/(disk_usage.used+disk_usage.free)).toFixed(2);;
              this.DISKFreeMessage = ((disk_usage.free*100)/(disk_usage.used+disk_usage.free)).toFixed(2);;
            }
          })
        },1000);
      }
    }
  };
</script>

<style scoped>
h1{
  text-align: center;
}
.CPUUsage{
  float: left;
  width: 50%;
}
.DISKUsage{
  float: right;
  width: 50%;
}
</style>


