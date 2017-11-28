


var colorObj = {
    // '0': '#008000', //呼吸线的颜色
    // '1': '#675bba', //脉搏线的颜色

    '0': 'red', //呼吸线的颜色
    '1': 'red', //脉搏线的颜色

    '2': '#000', //腋温线的颜色
    '3': '#000', //口温线的颜色
    '4': '#000' //肛温线的颜色

    // '2': '#fe9903', //腋温线的颜色
    // '3': '#fe2617', //口温线的颜色
    // '4': '#5793f3' //肛温线的颜色


}
option = {
    title: {
        show: false,
        text: 'title',
        left: 'center'
    },
    legend: {
        left: 'left',
        data: [' ', ' '],
        show: true,
        top: 20,
        left: 0,
        // width: 129,
        // padding: [0,0,0,10],
        // height:'100%',
        formatter: '{name}\n\n\n',
        itemWidth: 0,
        itemGap: 8,
        orient: 'horizontal',
        align: 'right',
        tooltip: {
            show: true
        },
    },
    tooltip: { //移动上去显示的提示
        show: true,
        trigger: 'item',
        formatter: '{a}<br/>{b} : {c}'
    },
    grid: {
        show: true,
        borderWidth: 0,
        // borderColor: "#000",
        shadowBlur: 0,
        // shadowColor: "#000",
        // width: 901.25,
        width:805.875,
        // height: 1000,
        height:  1000,
        // left: 128.75,
        left :116.125,
        top: 0,

        // top:120

    },
    xAxis: {
        type: 'category',
        data: ['4', '8', '12', '4', '8', '12', '4', '8', '12', '4', '8', '12', '4', '8', '12', '4', '8', '12', '4', '8', '12', '4', '8', '12', '4', '8', '12', '4', '8', '12', '4', '8', '12', '4', '8', '12', '4', '8', '12', '4', '8', '12'],
        axisLine: {
            show: false,
        },
        axisTick: {
            show: false
        },
        splitLine: {
            show: true,
            lineStyle: {
                // color: '#000'
                color: 'transparent',
            }
        },
        position: 'top'
    },

    // yAxis: {
    //     type: 'log',
    //     name: 'y',
    //     axisLine:{
    //         show:false
    //     },
    //     axisTick:{
    //         show:false
    //     },
    //     splitLine: {
    //         show: true,
    //         lineStyle: {
    //             color: 'red',
    //         }
    //     },
    // },
    yAxis: [
        {
            // type: 'category',
            type: 'value',
            name: '温度',
            min: 33,
            max: 42,
            position: 'left',
            offset: 40,
            // splitNumber: 50,
            splitNumber: 10,
            axisLine: {
                show: true,
                lineStyle: {
                    color: '#000'//colors[0]

                }
            },
            // interval:50,
            axisTick: {
                show: false
            },
            axisLabel: {
                formatter: '{value}°C',
                margin: 10,
                showMinLabel: false,
                showMaxLabel: false,
                inside: true
            },
            splitLine: {
                show: true,
                lineStyle: {
                    color: '#000',
                    // color: 'transparent',
                    width:2,
                },
                // interval:.5
            },

            // data: ['32', '33', '34', '35', '36', '37', '38', '39', '40', '41', '42']

        },
        {
            type: 'value',
            name: '口温',
            min: 33,
            max: 42,
            position: 'left',
            offset: 10,
            splitNumber: 50,
            // splitNumber: 10,
            nameTextStyle: {
                color: 'green'
            },
            axisLine: {
                show: false,
                lineStyle: {
                    color: '#000'
                }
            },
            axisTick: {
                show: false
            },
            axisLabel: {
                show: false,
                formatter: '{value} °C',

            },
            splitLine: {
                show: true,
                lineStyle: {
                    color: '#000',
                    // color: 'transparent',
                }
            },
            nameTextStyle: {
                color: '#000'
            },
            data: ['32', '33', '34', '35', '36', '37', '38', '39', '40', '41', '42']

        },
        {
            // type: 'category',
            type: 'value',
            name: '肛温',
            min: 33,
            max: 42,
            position: 'left',
            offset: 10,
            splitNumber: 50,
            // splitNumber: 10,
            nameTextStyle: {
                color: 'green'
            },
            axisLine: {
                show: false,
                lineStyle: {
                    // color: '#000',
                    color: 'transparent',
                }
            },
            axisTick: {
                show: false
            },
            axisLabel: {
                show: false,
                formatter: '{value} °C',

            },
            splitLine: {
                show: false,
                lineStyle: {
                    color: '#000',
                }
            },
            nameTextStyle: {
                color: '#000'
            },
            data: ['32', '33', '34', '35', '36', '37', '38', '39', '40', '41', '42']

        },
        {
            type: 'value',
            name: '脉搏',
            nameLocation:"end",
            nameTextStyle: {
                color: 'red',
            },
            min: 0,
            max: 180,
            position: 'left',
            offset: 72,
            axisLine: {
                show: false,
                lineStyle: {
                    color: 'red'//colors[1]
                    // color: 'transparent',
                }
            },
            axisLabel: {
                // color:'red',
                formatter: '{value}',
                showMinLabel: false,
                showMaxLabel: false,
                inside: true,
            },
            splitNumber: 10,

            axisTick: {
                show: false
            },
            splitLine: {
                show: false
            }
        },
        {
            type: 'value',
            name: '心率',
            nameTextStyle: {
                color: '#000'
            },
            min: 0,
            max: 180,
            position: 'left',
            offset: 109,
            axisLine: {
                show: false,
                lineStyle: {
                    color: '#000'//colors[1]
                }
            },
            axisLabel: {
                // formatter: '{value}\n',
                // showMinLabel: true,
                // showMaxLabel: false,
                // inside: true
                show:false
            },
            splitNumber: 10,
            axisTick: {
                show: false
            },
            splitLine: {
                show: false
            }
        },
    ],

    series: [
        {
            name: '肛温',
            type: 'line',
            smooth: true,
            lineStyle: {
                normal: {
                    color: colorObj[4],
                    width: 1
                }
            },
            itemStyle: {
                normal: {
                    color: colorObj[4],
                    // width: 1
                }
            },
            symbol: "emptyCircle",
            symbolSize: 8,
            zlevel:0,
            data: temp3,
            markLine: {
                symbolSize: 1,
                silent: true,
                label: {
                    normal: {
                        show: false
                    }
                },
                lineStyle: {
                    normal: {
                        type: 'solid',
                        color: 'red',
                        width: 2,

                    }
                },
                data: [
                    {
                        yAxis: 37,
                    },
                    // {
                    //     yAxis: 35
                    // },
                    // {
                    //     yAxis: 34
                    // },
                ]
            }
        },
        {
            name: '腋温',
            type: 'line',
            smooth: true,
            lineStyle: {
                normal: {
                    color: colorObj[2],
                    width: 1
                }
            },
            itemStyle: {
                normal: {
                    color: colorObj[2],
                    // width: 1
                }
            },
            symbol: "image:///static/dist/images/admin_pc/icon34.png",
            symbolSize: 12,
            data: temp1,
            zlevel:1,


            // markLine: {
            //     symbolSize: 1,
            //     silent: true,
            //     label: {
            //         normal: {
            //             show: false
            //         }
            //     },
            //     lineStyle: {
            //         normal: {
            //             type: 'solid',
            //             color: 'red',
            //             width: 1,

            //         }
            //     },
            //     data: [
            //         {
            //             yAxis: 37,
            //         },
            //         {
            //             yAxis: 35
            //         },
            //         {
            //             yAxis: 34
            //         },
            //     ]
            // }
        },
        {
            name: '口温',
            type: 'line',
            smooth: true,
            lineStyle: {
                normal: {
                    color: colorObj[3],
                    width: 1
                }
            },
            itemStyle: {
                normal: {
                    color: colorObj[3],
                    // width: 1
                }
            },
            zlevel:1,


            symbol: "circle",
            symbolSize: 8,
            data: temp2,
            // markLine: {
            //     symbolSize: 1,
            //     silent: true,
            //     label: {
            //         normal: {
            //             show: false
            //         }
            //     },
            //     lineStyle: {
            //         normal: {
            //             type: 'solid',
            //             color: 'red',
            //             width: 1,
            //         }
            //     },
            //     data: [
            //         {
            //             yAxis: 37,
            //         },
            //         {
            //             yAxis: 35
            //         },
            //         {
            //             yAxis: 34
            //         },
            //     ]
            // }
        },
        {
            name: '脉搏',
            yAxisIndex: 3,
            type: 'line',
            smooth: true,
            symbolSize: 8,
            symbol: 'circle',
            lineStyle: {
                normal: {
                    color: colorObj[1],
                    width: 1
                }
            },
            itemStyle: {
                normal: {
                    color: colorObj[1],
                    // width: 1
                }
            },

            zlevel:1,

            data: pulse,

        },
        {
            name: '心率',
            yAxisIndex: 4,
            symbol: "emptyCircle",
            type: 'line',
            smooth: true,
            symbolSize: 8,
            lineStyle: {
                normal: {
                    color: colorObj[0],
                    width: 1
                }
            },
            itemStyle: {
                normal: {
                    color: colorObj[0],
                    // width: 1
                }
            },
            data: heartrate,
            zlevel:1,

        }
        // {
        //     name: 'C',
        //     yAxisIndex: 2,
        //     type: 'line',
        //     data: [1, 2, 4, 8, 16, 32, 64, 128, 256, 1, 2, 4, 8, 16, 32, 64, 128],
        //     lineStyle: "yellow"
        // }
    ],
};



//生成从minNum到maxNum的随机数
function randomNum(minNum, maxNum) {
    return parseInt(Math.random() * (maxNum - minNum + 1) + minNum, 10);
}

function createArr(len, minNum, maxNum) {
    var tmpArr = [];
    for (var i = 0; i < len; i++) {
        tmpArr.push(randomNum(minNum, maxNum));
    }
    return tmpArr;

}

var hashCls = {
    '0': '腋温',
    '1': '口温',
    '2': '肛温',
    '3': '脉搏',
    '4': '心率',
}

// myChart.dispatchAction({
//       type: 'legendUnSelect',
//     // 图例名称
//     name: "口温"
// })

// $('.slb-bar').on('click', 'span', function () {
//     var $this = $(this);
//     if ($this.hasClass('sel')) {
//         $this.removeClass('sel')
//         myChart.dispatchAction({
//             type: 'legendUnSelect',
//             // 图例名称
//             name: hashCls[$this.data('val')]
//         })
//     } else {
//         $this.addClass('sel')
//         myChart.dispatchAction({
//             type: 'legendSelect',
//             // 图例名称
//             name: hashCls[$this.data('val')]
//         })
//     }
// })
$('.wd-tit').on('click', function () {
    var $this = $(this);
    if ($this.hasClass('sel')) {
        $this.removeClass('sel');
        $('.slb-bar').find('span').removeClass('sel')
        for (var i = 0; i < 3; i++) {
            myChart.dispatchAction({
                type: 'legendUnSelect',
                // 图例名称
                name: hashCls[i]
            })
        }
    } else {
        $this.addClass('sel');
        $('.slb-bar').find('span').addClass('sel')
        for (var i = 0; i < 3; i++) {
            myChart.dispatchAction({
                type: 'legendSelect',
                // 图例名称
                name: hashCls[i]
            })
        }
    }
})

// $('.slb-bar').on('mouseover', 'span', function () {

//     myChart.dispatchAction({
//         type: 'highlight',
//         seriesName: hashCls[$(this).data('val')],
//     })
// })
// $('.slb-bar').on('mouseout', 'span', function () {

//     myChart.dispatchAction({
//         type: 'downplay',
//         seriesName: hashCls[$(this).data('val')],

//     })
// })
$('.tit-bottom').on('mouseover','.kw,.yw,.gw,.mb,.xl', function () {
    myChart.dispatchAction({
        type: 'highlight',
        seriesName: hashCls[$(this).data('val')],
    })
})

$('.tit-bottom').on('mouseout', '.kw,.yw,.gw,.mb,.xl', function () {

    myChart.dispatchAction({
        type: 'downplay',
        seriesName: hashCls[$(this).data('val')],

    })
})
$('.tit-bottom').on('click','.kw,.yw,.gw,.mb,.xl',function(){
    var $this = $(this);
    if ($this.hasClass('sel')) {
        $this.removeClass('sel');
        myChart.dispatchAction({
            type: 'legendUnSelect',
            // 图例名称
            name: hashCls[$(this).data('val')]
        })
    }else{
        $this.addClass('sel');
        myChart.dispatchAction({
            type: 'legendSelect',
            // 图例名称
            name: hashCls[$(this).data('val')]
        })
    }
})

$('.mb-tit').on('mouseover', function () {
    myChart.dispatchAction({
        type: 'highlight',
        seriesName: hashCls[$(this).data('val')],
    })
})

$('.mb-tit').on('mouseout',  function () {

    myChart.dispatchAction({
        type: 'downplay',
        seriesName: hashCls[$(this).data('val')],

    })
})
$('.mb-tit').on('click',function(){
    var $this = $(this);
    if ($this.hasClass('sel')) {
        $this.removeClass('sel');
        myChart.dispatchAction({
            type: 'legendUnSelect',
            // 图例名称
            name: hashCls[$(this).data('val')]
        })
    }else{
        $this.addClass('sel');
        myChart.dispatchAction({
            type: 'legendSelect',
            // 图例名称
            name: hashCls[$(this).data('val')]
        })
    }
})
