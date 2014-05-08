<%@ Page Language="C#" AutoEventWireup="true" CodeFile="ajaxValidator.aspx.cs" Inherits="ajaxValidator" %>

<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">

<html xmlns="http://www.w3.org/1999/xhtml" >
<head runat="server">
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
<title>jQuery formValidator表单验证插件 -- by:猫冬，email:wzmaodong@126.com</title>
<meta name="Author" content="猫冬" />
<meta name="description" content="jQuery formValidator表单验证插件" />
<meta name="keywords" content="jQuery,formValidator,插件,表单,验证,插件,javascript,表单验证,提示层" />
<script src="jquery-1.4.4.min.js" type="text/javascript"></script>
<script src="formValidator-4.1.3.js" type="text/javascript" charset="UTF-8"></script>
<script src="formValidatorRegex.js" type="text/javascript" charset="UTF-8"></script>
<script type="text/javascript">
$(document).ready(function(){
	$.formValidator.initConfig({formID:"form1",theme:"ArrowSolidBox",submitOnce:true,
		onError:function(msg,obj,errorlist){
			$("#errorlist").empty();
			$.map(errorlist,function(msg){
				$("#errorlist").append("<li>" + msg + "</li>")
			});
			alert(msg);
		},
		ajaxPrompt : '有数据正在异步验证，请稍等...'
	});

	$("#us").formValidator({onShowText:"请输入用户名",onShow:"请输入用户名,只有输入\"maodong\"才是对的",onFocus:"用户名至少5个字符,最多10个字符",onCorrect:"该用户名可以注册"}).inputValidator({min:5,max:10,onError:"你输入的用户名非法,请确认"})//.regexValidator({regExp:"username",dataType:"enum",onError:"用户名格式不正确"})
	    .ajaxValidator({
		dataType : "json",
		async : true,
		url : "http://www.yhuan.com/Handler.ashx",
		success : function(data){
            if( data == "0" ) return true;
			return "该用户名不可用，请更换用户名";
		},
		buttons: $("#button"),
		error: function(jqXHR, textStatus, errorThrown){alert("服务器没有返回数据，可能服务器忙，请重试"+errorThrown);},
		onError : "该用户名不可用，请更换用户名",
		onWait : "正在对用户名进行合法性校验，请稍候..."
	});
});
</script>
</head>
<body>
    <form id="form1" runat="server">
    <input id="btnA" type="button" value="Ajax提交" style="margin-left: 70px;" />&nbsp;
        <input ID="Button1" type="submit" Text="普通提交" /><br />
      <a href="ajaxValidator.aspx?us=123456">编辑用户名为:123456</a><br />
      用户名:<asp:TextBox ID="us" runat="server" style="width:120px"></asp:TextBox><div id="usTip" style="width:280px">
    </form>
    <script src="http://s13.cnzz.com/stat.php?id=1123054&web_id=1123054&show=pic" language="JavaScript"></script>
</body>
</html>
