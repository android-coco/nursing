"use strict";

//退出登录
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


// 设置侧边栏当前菜单
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
try{
    if(menuIndex){
        setMenu(menuIndex);
    }
}catch(e){
    console.log("menuIndex不存在")
}


//侧边栏子菜单收放效果
$('#jky-menu').on('open.collapse.amui', function(e) {
    $(e.target).parents("li").addClass("jky-on")
})
$('#jky-menu').on('close.collapse.amui', function(e) {
    $(e.target).parents("li").removeClass("jky-on")
})

// 限制长度




// ================tool======================

/**
 * 获取URL参数
 * @param name
 * @returns {*}
 */
function getURLParam(name) {
    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)", "i");
    var r = window.location.search.substr(1).match(reg);
    if (r != null) {
        return (r[2]);
    }
    return null;
}
// 获取当前时间 hh:mm:ss
function getNowTime() {
    var str = "";
    var nowD = new Date();
    str = addZero(nowD.getHours()) + ":" + addZero(nowD.getMinutes()) + ":" +addZero(nowD.getSeconds());
    return str;
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


/**
 * 格式化时间
 * @param date //可 时间戳,时间字符串文本,null(返回当前时间)
 * @returns {data:dataStr,time:timeStr}  {data:"2017-10-10",time:"00:00"}
 */
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

/**
 * 格式化时间
 * @param date //可 时间戳,时间字符串文本,null(返回当前时间)
 * @returns {date:29,hours:19,minutes:39,month:10,seconds:0,year:2017}
 */
function fromatDateObj(date){
    var res = {};
    var nowD = Date();
    if(date){
        nowD= new Date(date);
    }else{
        nowD = new Date();
    }
    res.date = nowD.getDate()
    res.hours= nowD.getHours()
    res.month = nowD.getMonth() + 1
    res.minutes = nowD.getMinutes()
    res.seconds = nowD.getSeconds()
    res.year = nowD.getFullYear()

        // nowD.getFullYear() + "-" + addZero(nowD.getMonth() + 1) + "-" + addZero(nowD.getDate());
    // res.time = addZero(nowD.getHours()) + ":" + addZero(nowD.getMinutes())
    return res
}

/**
 * 不足10补0
 * @param i
 * @returns {*}
 */
function addZero(i) {
    if (i < 10) {
        return "0" + i;
    } else {
        return i;
    }
}

/**
 * 获取当前时间戳
 * @type {(() => number) | Function}
 */
var getNowTimeStamp = Date.now || function() {
    return new Date().getTime()
}

/**
 * 节流函数
 * @param func 执行函数
 * @param wait 等待时间
 * @param option 配置项
 * @returns {Function}
 */
function throttle(func, wait, option) {
    var previous = 0, //之前的时间
        timeout, self, args, result;

    var later = function() {
        previous = getNowTimeStamp();
        timeout = null;
        result = func.apply(self, args);
        !timeout && (self = args = null);
    }
    return function() {
        var now = getNowTimeStamp();

        if (!previous && option.leading === false) previous = now;
        var diff = wait - (now - previous);

        args = arguments;
        self = this;
        if (diff <= 0) {
            if (timeout) {
                clearTimeout(timeout);
                timeout = null;
            }
            result = func.apply(self, args);
            previous = now;
            !timeout && (self = args = null); //如果没有堆积的任务 则清空 self 和args
        } else if (!timeout && option.trailing !== false) {
            timeout = setTimeout(later, diff)
        } else {
            console.log('被抛弃了')
        }
        return result;

    }
}
