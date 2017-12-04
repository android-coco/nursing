"use strict";

function nrl_page_control(ty) {
    // 起止时间
    var hospitalDate = "{{.PInfo.HospitalDate}}";
    var nowDate = nowDateObj.date,nowTime = nowDateObj.time;

    var date1, date2 = 0;
    var a = new Date(hospitalDate);
    date1 = a.getTime();
    var b = new Date(nowDate+" "+nowTime);
    date2 = b.getTime();


    var sdate = getURLParam("sdate");
    var edate = getURLParam("edate");
    var stime="",etime = "";
    console.log("----", sdate, edate);
    if (sdate == null || sdate == "0" || edate == null || edate == "0") {
        sdate = hospitalDate.substr(0, 10);
        edate = nowDate;
        stime = hospitalDate.substr(11,5);
        etime = nowDateObj.time.substr(0,5);
    } else {
        date1 = sdate;
        date2 = edate;
        // alert(fromatDate(parseInt(sdate)).time)
        stime =fromatDate(parseInt(sdate)).time;
        etime = fromatDate(parseInt(edate)).time;

        sdate = fromatDate(parseInt(sdate)).date;
        edate = fromatDate(parseInt(edate)).date;

        // console.log(stime,etime)
        // alert(stime,etime)
    }


    laydate.render({
        elem: '#date1', //指定元素
        showBottom: false,
        value: sdate,

        done: function (value, date, endDate) {
            console.log(value); //得到日期生成的值，如：2017-08-18
            console.log(date); //得到日期时间对象：{year: 2017, month: 8, date: 18, hours: 0, minutes: 0, seconds: 0}
            console.log(endDate); //得结束的日期时间对象，开启范围选择（range: true）才会返回。对象成员同上。

            var timeStr = $("#time1").val()+":00"
            var valStr = value + " " + timeStr
            var a = new Date(valStr);
            date1 = a.getTime();
            reloadSheet();
        }
    });
    laydate.render({
        elem: '#time1', //指定元素
        type: 'time',
        format: 'HH:mm',
        min:"00:00:00",
        max:"23:59:00",
        value: stime,
        done: function (value, date, endDate) {
            var timeStr = value+":00";
            var valueStr = $("#date1").val()+" " + timeStr
            var a = new Date(valueStr);
            date1 = a.getTime();
            reloadSheet();
        }
    });
    laydate.render({
        elem: '#time2', //指定元素
        type: 'time',
        format: 'HH:mm',
        value: etime,
        done: function (value, date, endDate) {
            var timeStr = value+":00";
            var valueStr = $("#date2").val()+" " + timeStr
            var a = new Date(valueStr);
            date2 = a.getTime();
            reloadSheet();
        }
    });

    laydate.render({
        elem: '#date2', //指定元素
        showBottom: false,
        value: edate,
        done: function (value, date, endDate) {

            var timeStr = $("#time2").val()+":00"
            var valStr = value + " " + timeStr

            var a = new Date(valStr);
            date2 = a.getTime();
            console.log("date ", date1, date2)
            // reloadSheet();
        }
    });


    // 总页数
    var pagenum = {{.PageNum}};
    // 当前页数
    var index = {{.PageIndex}};
    var pid = parseInt({{.PInfo.VAA01}});

    var baseurl = "/pc/record/nrl" + ty + "?";
    var tempindex = index;

    $("#leftLast-btn").on("click", function () {
        if (index == 1) return;
        tempindex = 1;
        reloadSheet()
    });

    $("#left-btn").on("click", function () {
        if (index <= 1) {
            return
        } else {
            tempindex = index - 1;
        }
        reloadSheet();
    });

    $("#right-btn").on("click", function () {
        if (index >= pagenum) {
            // tempindex = pagenum;
            return
        } else {
            tempindex = index + 1;
        }
        reloadSheet();
    });


    $("#rightLast-btn").on("click", function () {
        if (index == pagenum) return;
        tempindex = pagenum;
        reloadSheet()
    });

    $("#go-btn").on("click", function () {
        var val = $(".pageNum").val();
        tempindex = val;
        reloadSheet()
    });

    function reloadSheet() {
        if ($(".col-cont.active").length > 0) {
            alert("请先保存数据！");
            return false;
        }
        index = tempindex;
        window.location.href = baseurl + "pid=" + pid + "&num=" + index + "&sdate=" + date1 + "&edate=" + date2;
    };


    $(".user-list").on("click", "li", function () {
        var pid = $(this).data('msg');
        console.log($(this).data('msg'));
        window.location.href = "/pc/record/nrl" + ty + "?pid=" + pid;
    })
}