
$(document).ready(function () {
    //判断是否选中
    isCheck()
    //记住密码功能
    var str = getCookie("loginInfo");
    str = str.substring(0, str.length);
    var username = str.split(",")[0];
    var password = str.split(",")[1];
    //自动填充用户名和密码
    $("input[name='Username']").val(username);
    $("input[name='Password']").val(password);
});

//获取cookie
function getCookie(c_name) {
    if (document.cookie.length > 0) {
        c_start = document.cookie.indexOf(c_name + "=")
        if (c_start != -1) {
            c_start = c_start + c_name.length + 1
            c_end = document.cookie.indexOf(";", c_start)
            if (c_end == -1) c_end = document.cookie.length
            $("input[type='checkbox']").prop("checked", true);
            return unescape(document.cookie.substring(c_start, c_end))
        }
    }
    return ""
}

function isCheck() {
    if ($("input[type='checkbox']").is(':checked')) {
        //改变value值为1
        $("input[type='checkbox']").val("1")
    } else {
        //改变value值为0
        $("input[type='checkbox']").val("0")
        delCookie("loginInfo") 

    }
}

function delCookie(name) {
    var date = new Date();
    date.setTime(date.getTime() - 10000);
    document.cookie = name + "=a; expires=" + date.toGMTString();
} 
