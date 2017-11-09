"use strict";

$('.s-btn2').on('click', function () {
    if ($(this).hasClass('active')) {
        $(this).removeClass('active');
        $('.left_history_menu').removeClass('none');
        $('.admin').removeClass("active");

    } else {
        $(this).addClass('active');

        $('.left_history_menu').addClass('none');
        $('.admin').addClass("active");

    }
});


var nowDateObj = nowDate();

function nowDate(day) {
    var res = {},
        nowD = new Date();
    day = day || 0
    var timeinterval = day*86400000;
    nowD = new Date(nowD.getTime()-timeinterval)
    res.date = nowD.getFullYear() + "-" + addZero(nowD.getMonth() + 1) + "-" + addZero(nowD.getDate());
    res.time = addZero(nowD.getHours()) + ":" + addZero(nowD.getMinutes())
    return res
}


function addZero(i) {
    if (i < 10) {
        return "0" + i;
    } else {
        return i;
    }
}

laydate.render({
    elem: '#date3', //指定元素
    showBottom: false,
    value: nowDate(2).date,
    // done: function (value, date, endDate) {
    //     console.log(value); //得到日期生成的值，如：2017-08-18
    //     console.log(date); //得到日期时间对象：{year: 2017, month: 8, date: 18, hours: 0, minutes: 0, seconds: 0}
    //     console.log(endDate); //得结束的日期时间对象，开启范围选择（range: true）才会返回。对象成员同上。
    // }
});
laydate.render({
    elem: '#date4', //指定元素
    showBottom: false,
    value: nowDate(0).date,
    // done: function (value, date, endDate) {
    //     console.log(value); //得到日期生成的值，如：2017-08-18
    //     console.log(date); //得到日期时间对象：{year: 2017, month: 8, date: 18, hours: 0, minutes: 0, seconds: 0}
    //     console.log(endDate); //得结束的日期时间对象，开启范围选择（range: true）才会返回。对象成员同上。
    // }
});

// $("body").on('click','.tr-item',function(){
//     $(this).parents(".top_tabel").find(".tr-item").removeClass("active")
//     $(this).addClass("active")
//
//     const patient = $(this).data('pid')
//     console.log("病人ID："+patient)
// })

$('#history-search-btn').on('click', function () {
    var parameters = {}
    if ($('#timeSel').is(':checked')) {
        parameters['staDate'] = $('#date3').val() + " 00:00:00"
        parameters['endDate'] = $('#date4').val() + " 23:59:59"
    } else if ($('#id1').is(':checked')) {
        if ($('#rgtNum').val().length == 0) {
            alert("请输入完整的登记号")
            return
        } else {
            parameters['rgtNum'] = $('#rgtNum').val()
        }
    } else if ($('#id2').is(':checked')) {
        if ($('#hospNum').val().length == 0) {
            alert("请输入完整的住院号")
            return
        } else {
            parameters['hospNum'] = $('#hospNum').val()
        }
    } else if ($('#name').is(':checked')) {
        if ($('#patName').val().length == 0) {
            alert("请输入姓名或者拼音首字母")
            return
        } else {
            parameters['name'] = $('#patName').val()
        }
    }
    parameters['did'] = $(this).data('did')

    $.get('/pc/history/search',parameters,function (json, status) {
        if (json.result == 0) {
            var resArr = json.datas;
            var html = '';
            for (var i = 0; i < resArr.length; i++) {
                const patient = resArr[i]
                html += "<div class=\"tr-item\"  data-pid=\""+patient.patientId+"\">\n" +
                    "<div class=\"td-1\">"+patient.name+"</div>\n" +
                    "<div class=\"td-2\">"+patient.rgtNum+"</div>\n" +
                    "<div class=\"td-3\">"+patient.hospNum+"</div>\n" +
                    "<div class=\"td-4\">"+patient.dcgDate+"</div>\n" +
                    "</div>";
            }
            $(".search-patient").html(html)

        } else {
            alert(json.errmsg)
        }
    }, 'json')

});



