"use strict";
// 该文件依赖jq
// --------------------------------

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

/**
 * 判断对象是否相等
 * @param x
 * @param y
 * @returns {*}
 */
function equals( x, y ) {
    var in1 = x instanceof Object;
    var in2 = y instanceof Object;
    if(!in1||!in2){
        return x===y;
    }
    if(Object.keys(x).length!==Object.keys(y).length){
        return false;
    }
    for(var p in x){
        var a = x[p] instanceof Object;
        var b = y[p] instanceof Object;
        if(a&&b){
            if(!equals( x[p], y[p])){
                return false;
            }
           // return equals( x[p], y[p]);
        }
        else if(x[p]!==y[p]){
            return false;
        }
    }

    return true;
}

// 获取数据类型
function getType(data){
    var typeStr = Object.prototype.toString.call(data)
    typeStr = typeStr.toLocaleLowerCase()
    var reg = /\[object (.*)\]/g
    typeStr = typeStr.replace(reg,"$1");
    return typeStr
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

/**
 * alert弹窗
 * @param txt  文本内容
 * @param type icon类型 "ok","no" 或不选
 * @param fn  1秒后执行的回调
 */
function jkyAlert(txt,type,fn){
    var typeIco ={
        "yes":{
            class:"am-icon-check",
            color:"#5eb95e"
        },
        "no":{
            class:"am-icon-remove",
            color:"#dd514c"
        },
        "default":{
            class:"none",
            color:"#fff"
        }
    }
    type = type || "default"
    if($("#jky-modal-alert").length==0){

        var spanIconHtml= '<span id="ico-alert" class="' + typeIco[type].class+'" style="font-size:40px;color:'+ typeIco[type].color+';"></span>\n';

        var modalHtml ='<div class="am-modal am-modal-loading am-modal-no-btn" tabindex="-1" id="jky-modal-alert" style="border-radius: 6px;">\n' +
        '    <div class="am-modal-dialog">\n' +
        '        <div class="am-modal-hd" id="txt-alert">'+ txt +'</div>\n' +
        '        <div class="am-modal-bd">\n' +
        spanIconHtml +
        '        </div>\n' +
        '    </div>\n' +
        '</div>';
        $("body").append(modalHtml)
    }else{
        $("#jky-modal-alert").find("#ico-alert").attr("class",typeIco[type].class)
        $("#jky-modal-alert").find("#ico-alert").css("color",typeIco[type].color)
        $("#jky-modal-alert").find("#txt-alert").text(txt)

    }
    $("#jky-modal-alert").modal("open")
    setTimeout(function(){
        $("#jky-modal-alert").modal("close")
        fn && fn()
    },1000)

}

/**
 * Confirm弹窗
 * @param txt 文本内容
 * @param confirmFn 确认的回调
 * @param cancaelFn 取消的回调
 * @param confirmTxt 确定的文本内容
 * @param cancaelTxt 取消的文本内容
 */
function jkyConfirm(txt,confirmFn,cancaelFn,confirmTxt,cancaelTxt){
    var $confirmEl = null;


    if($('#jky-confirm').length>0){
        $confirmEl = $('#jky-confirm')
        $confirmEl.find("#txt-confirm").text(txt)
    }else{
        $confirmEl = $("<div class=\"am-modal am-modal-confirm\" tabindex=\"-1\" id=\"my-confirm\">\n" +
        "    <div class=\"am-modal-dialog\">\n" +
        // "        <div class=\"am-modal-hd\">Amaze UI</div>\n" +
        "        <div class=\"am-modal-bd\" id='txt-confirm'>\n" +
            txt+
        // "            你，确定要删除这条记录吗？\n" +
        "        </div>\n" +
        "        <div class=\"am-modal-footer\">\n" +
        "            <span class=\"am-modal-btn\" data-am-modal-cancel>"+(cancaelTxt || "取消")+"</span>\n" +
        "            <span class=\"am-modal-btn\" data-am-modal-confirm>"+(confirmTxt || "确定")+"</span>\n" +
        "        </div>\n" +
        "    </div>\n" +
        "</div>");

        $("body").append($confirmEl)
    }

    $confirmEl.modal({
        relatedTarget: this,
        onConfirm: function() {
            confirmFn && confirmFn()
        },
        // closeOnConfirm: false,
        onCancel: function() {
            cancaelFn && cancaelFn();
        }
    });
}


/**
 * 输入框弹窗
 * @param option{txt,valTxt 输入框的默认值,maxlength 限制长度}
 * @param confirmFn
 * @param cancaelFn
 */
function jkyPrompt(option,confirmFn,cancaelFn){
    var $promptEl = null;
    if(typeof option != "object"){
        return false
    }
    option.valTxt = option.valTxt || ""
    option.txt = option.txt || ""

    var isExist = $('#jky-prompt').length > 0;

    if(isExist){
        $promptEl = $('#jky-prompt')
        $promptEl.find("#txt-prompt").text(option.txt)
        $promptEl.find("#jky-prompt-input").val(option.valTxt)
    }else{
        $promptEl = $(' <div class="am-modal am-modal-prompt" tabindex="-1" id="jky-prompt">\n' +
        '        <div class="am-modal-dialog">\n' +
        // '            <div class="am-modal-hd">Amaze UI</div>\n' +
        '            <div class="am-modal-bd" >\n' +
                    '<div id="txt-prompt">'+option.txt +'</div>'+
             '<input type="text" class="am-modal-prompt-input" id="jky-prompt-input" value="' + option.valTxt+'">\n'+
        '            </div>\n' +
        '            <div class="am-modal-footer">\n' +
        '                <span class="am-modal-btn" data-am-modal-cancel>取消</span>\n' +
        '                <span class="am-modal-btn" data-am-modal-confirm>提交</span>\n' +
        '            </div>\n' +
        '        </div>\n' +
        '    </div>');
    }

    if(option.maxlength){
        $promptEl.find("#jky-prompt-input").attr("maxlength",option.maxlength)
    }else{
        $promptEl.find("#jky-prompt-input").removeAttr("maxlength")
    }



    if(isExist){
        var oldOption =  $promptEl.data('amui.modal')
        console.log(oldOption)
        oldOption.options.onConfirm =  function(e) {
            confirmFn && confirmFn(e.data)
        };
        oldOption.options.onCancel =  function(e) {
            cancaelFn && cancaelFn(e.data)
        };
        oldOption.toggle(this);

    }else{
        $("body").append($promptEl)
        $('#jky-prompt').modal({
            relatedTarget: this,
            onConfirm: function(e) {
                confirmFn && confirmFn(e.data)
            },
            onCancel: function(e) {
                cancaelFn && cancaelFn(e.data)
            }
        });
    }
}
