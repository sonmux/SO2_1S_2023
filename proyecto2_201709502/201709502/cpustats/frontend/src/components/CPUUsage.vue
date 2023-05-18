<script setup>
import {reactive} from 'vue'
import {Greet} from '../../wailsjs/go/main/App'


import {GetCPUUsage, GetDISKUsage, GetMemoryUsage} from '../../wailsjs/go/sys/Stats'

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
  <br>
  <div class="MEMORYUsage">
    <h3>Memory Total: {{MEMORYTotalMessage}} MB</h3>
    <h3>Memory Usage: {{MEMORYUsageMessage}} MB</h3>
    <h3>Memory Free: {{MEMORYFreeMessage}} MB</h3>
    <h3>Memory Available: {{MEMORYAvailableMessage}} MB</h3>
    <apexchart type="donut" width="480" :options="memoryoptions" :series="memoryseries"></apexchart>
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

        memoryseries: [0,100],
        memoryoptions: {
          labels: ["Used","Free"],
          title: {
            text: "Memory",
            align: "center"
          }
        },
        MEMORYTotalMessage:"DataMemoryTotal",
        MEMORYUsageMessage:"DataMemoryUsed",
        MEMORYFreeMessage:"DataMemoryFree",
        MEMORYAvailableMessage:"DataMemoryAvilable"

        /*memoryseries2: [44, 55, 67, 83],
        memoryoptions2: {
          chart: {
            height: 350,
            type: 'radialBar',
          },
          plotOptions: {
            radialBar: {
              dataLabels: {
                name: {
                  fontSize: '22px',
                },
                value: {
                  fontSize: '16px',
                },
                total: {
                  show: true,
                  label: 'MEMORY RAM',
                  formatter: function (w) {
                      GetDISKUsage().then(disk_usage => {
                      //this.message = disk_usage.avg;
                      return disk_usage.total
                    })
                  }
                }
              }
            }
          },
          labels: ['Total', 'Used', 'Free', 'Available'],
        },
        MEMORYTotalMessage:"DataMemoryTotal",
        MEMORYUsageMessage:"DataMemoryUsed",
        MEMORYFreeMessage:"DataMemoryFree",
        MEMORYAvailableMessage:"DataMemoryAvilable"*/
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
          this.DISKUsageMessage = ((disk_usage.used*100)/(disk_usage.used+disk_usage.free)).toFixed(2);
          this.DISKFreeMessage = ((disk_usage.free*100)/(disk_usage.used+disk_usage.free)).toFixed(2);
        }
      })
      GetMemoryUsage().then(memory_usage => {
        //this.message = disk_usage.avg;
        if (memory_usage) {
          //this.memoryseries2 = [ memory_usage.total, memory_usage.used, memory_usage.free, memory_usage.Available ];
          //this.MEMORYsageMessage = ((memory_usage.used*100)/(memory_usage.used+memory_usage.free)).toFixed(2);
          //this.MEMORYFreeMessage = ((memory_usage.free*100)/(memory_usage.used+memory_usage.free)).toFixed(2);
          this.memoryseries = [ memory_usage.used, memory_usage.free ];
          this.MEMORYTotalMessage = memory_usage.total;
          this.MEMORYUsageMessage = memory_usage.used;
          this.MEMORYFreeMessage = memory_usage.free;
          this.MEMORYAvailableMessage = memory_usage.Available;
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
        setInterval(()=>{
          GetMemoryUsage().then(memory_usage => {
            //this.message = cpu_usage.avg;
            if (memory_usage) {
              //this.memoryseries2 = [ memory_usage.total, memory_usage.used, memory_usage.free, memory_usage.Available ];
              this.memoryseries = [ memory_usage.used, memory_usage.free ];
              this.MEMORYTotalMessage = memory_usage.total;
              this.MEMORYUsageMessage = memory_usage.used;
              this.MEMORYFreeMessage = memory_usage.free;
              this.MEMORYAvailableMessage = memory_usage.Available;
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

.MEMORYUsage{
  float: left;
  width: 50%;
}

.MEMORYUsage2{
  float: left;
  width: 50%;
}
</style>


