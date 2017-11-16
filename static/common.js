"use strict";

$('.out_btn').on('click', function () {
    $.get("/pc/logout", {}, function (json, status) {
        console.log(json, status);
        if (json.result == 0) {
            window.location.href = "/pc/login"
        } else {
            alert("服务器繁忙，请稍后操作")
        }
    }, 'json');
})


// jky-menu
function setMenu(index){
    console.log(index)
    var arr = index.split("-");
    if(arr.length){
        $(".admin-sidebar-list>li").removeClass("active")
        $(".admin-sidebar-list>li").eq(parseInt(arr[0])-1).addClass("active")

        if(parseInt(arr[0])==1){
            $(".admin-sidebar-list>li").eq(0).find("ul").find("li").eq(parseInt(arr[1])).addClass("active")
        }else if(parseInt(arr[1])!=0){
            $(".admin-sidebar-list>li").eq(parseInt(arr[0])-1).find("ul").find("li").eq(parseInt(arr[1])-1).addClass("active")
        }

        if(parseInt(arr[0])==1 || parseInt(arr[0])==7){
            $(".admin-sidebar-list>li").eq(parseInt(arr[0])-1).addClass("jky-on")
            $(".admin-sidebar-list>li").eq(parseInt(arr[0])-1).find(".am-collapse").addClass("am-in")

        }else{
            $(".admin-sidebar-list>li").removeClass("jky-on")
            $(".admin-sidebar-list>li").find(".am-collapse").removeClass("am-in")
        }

    }
}


setMenu(menuIndex);

// jky 获取URL参数
function getURLParam(name) {
    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)", "i");
    var r = window.location.search.substr(1).match(reg);
    if (r != null) {
        return (r[2]);
    }
    return null;
}


// 获取当前时间 YYYY-MM-DD
function getNowDate() {
    var str = "";
    var nowdate = new Date();
    str = nowdate.getFullYear() + "-" + addZero(nowdate.getMonth() + 1) + "-" + addZero(nowdate.getDate());
    return str;
}

// 获取当前时间 YYYY-MM-DD hh:mm
function getNowDateTime () {
    var res = {},
        nowD = new Date();
    res.date = nowD.getFullYear() + "-" + addZero(nowD.getMonth() + 1) + "-" + addZero(nowD.getDate());
    res.time = addZero(nowD.getHours()) + ":" + addZero(nowD.getMinutes())
    return res
}


// 格式化时间
function fromatDate(date){
    var res = {};
    var nowD = Date();
    if(date){
        nowD= new Date(date);
    }else{
        nowD = new Date();
    }
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