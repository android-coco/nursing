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


setMenu(menuIndex)
