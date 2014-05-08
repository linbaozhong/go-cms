<!DOCTYPE html>
<html lang="zh-CN">
<head>
	<meta charset="utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no">
	<meta name="description" content="{{.page.Description}}">
	<meta name="keywords" content="{{.page.Keywords}}">
	<meta name="author" content="{{.page.Author}}">
	<title>{{.page.Title}}</title>
	<link rel="stylesheet" href="/static/css/bootstrap.min.css" />
	<link href="/static/ueditor/themes/default/css/umeditor.min.css" type="text/css" rel="stylesheet">
	<link href="/static/css/bootstrap-datetimepicker.min.css" type="text/css" rel="stylesheet">
	<script type="text/javascript" src="/static/js/jquery-1.10.2.min.js"></script>
	<script type="text/javascript" charset="utf-8" src="/static/js/bootstrap.min.js"></script>
	<script type="text/javascript" charset="utf-8" src="/static/js/bootstrap-datetimepicker.min.js"></script>
	<script type="text/javascript" charset="utf-8" src="/static/js/bootstrap-datetimepicker.zh-CN.js"></script>
	<script type="text/javascript" charset="utf-8" src="/static/ueditor/umeditor.config.js"></script>
	<script type="text/javascript" charset="utf-8" src="/static/ueditor/umeditor.min.js"></script>
	<script type="text/javascript" charset="utf-8" src="/static/js/jquery.validate.min.js"></script>
	<script type="text/javascript" charset="utf-8" src="/static/js/messages_zh.js"></script>
	<script type="text/javascript" charset="utf-8" src="/static/js/jquery.form.min.js"></script>

	<script type="text/javascript">
		cmsapi={
			"channelSequence":"/admin/channel/sequence"
			,"channelStatus":"/admin/channel/reset"	
			,"channelChildren":"/admin/channel/children"
			,"getChannels":"/admin/channel/getall"
			,"articleSequence":"/admin/article/sequence"
			,"articleStatus":"/admin/article/reset"
			,"getArticles":"/admin/article/getall"
			,"imageSequence":"/admin/image/sequence"
			,"imageStatus":"/admin/image/reset"
			,"accountStatus":"/admin/account/reset"
		};
		var valid={}
		valid.isInt=function(s){
			return  !isNaN(parseInt(s))
		}
		// 对Date的扩展，将 Date 转化为指定格式的String 
		// 月(M)、日(d)、小时(h)、分(m)、秒(s)、季度(q) 可以用 1-2 个占位符， 
		// 年(y)可以用 1-4 个占位符，毫秒(S)只能用 1 个占位符(是 1-3 位的数字) 
		// 例子： 
		// (new Date()).Format("yyyy-MM-dd hh:mm:ss.S") ==> 2006-07-02 08:09:04.423 
		// (new Date()).Format("yyyy-M-d h:m:s.S")      ==> 2006-7-2 8:9:4.18 
		Date.prototype.Format = function(fmt) 
		{ //author: meizz 
		  var o = { 
		    "M+" : this.getMonth()+1,                 //月份 
		    "d+" : this.getDate(),                    //日 
		    "h+" : this.getHours(),                   //小时 
		    "m+" : this.getMinutes(),                 //分 
		    "s+" : this.getSeconds(),                 //秒 
		    "q+" : Math.floor((this.getMonth()+3)/3), //季度 
		    "S"  : this.getMilliseconds()             //毫秒 
		  }; 
		  if(/(y+)/.test(fmt)) 
		    fmt=fmt.replace(RegExp.$1, (this.getFullYear()+"").substr(4 - RegExp.$1.length)); 
		  for(var k in o) 
		    if(new RegExp("("+ k +")").test(fmt)) 
		  fmt = fmt.replace(RegExp.$1, (RegExp.$1.length==1) ? (o[k]) : (("00"+ o[k]).substr((""+ o[k]).length))); 
		  return fmt; 
		}
	</script>
</head>
<body>
	<header>
		<nav class="navbar navbar-default" role="navigation">
			<div class="navbar-header">
				<a class="navbar-brand" href="#">{{.page.SiteName}}</a>
			</div>
			<div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
				<ul class="nav navbar-nav">
					<li {{if eq .path "channel"}}class="active"{{end}}>
						<a href="/admin/channel">频道管理</a>
					</li>
					<li {{if eq .path "article"}}class="active"{{end}}>
						<a href="/admin/article">内容管理</a>
					</li>
					{{if .isSuperAdmin}}
					<li {{if eq .path "account"}}class="active"{{end}}>
						<a href="/admin/account">账户管理</a>
					</li>
					{{end}}
				</ul>
				<p class="navbar-text pull-right">
					{{if gt .current.Updator 0}}
					当前用户：{{.current.Name}}
					<span>&nbsp;&nbsp;</span>
					<a href="/admin/profile" class="navbar-link">帐号设置</a>
					<span>|</span>
					<a class="navbar-link" href="/home/logout">注销</a>
					{{else}}
					<a class="navbar-link" href="/home/login">登录</a>
					{{end}}
				</p>
			</div>
		</nav>
	</header>