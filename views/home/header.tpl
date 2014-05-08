<!DOCTYPE HTML>
<html>
<head>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<title>{{.page.Title}}</title>
	<meta name="description" content="{{.page.Description}}">
	<meta name="keywords" content="{{.page.Keywords}}">
	<meta name="author" content="{{.page.Author}}">
	<link href="/static/css/index.css" type="text/css" rel="stylesheet"/>
</head>

<body>
	<div class="top">
		<div class="top_cen">
			<h1>
				<img src="/static/images/log.jpg"/>
			</h1>
			<div class="Rbut">
				<p>
					<a style="padding-right:5px;" href="#">联系我们</a>
					<a href="#">关注我们</a>
				</p>
			</div>
		</div>
	</div>
	<div class="nav">
		{{Navibar .action 1}}
		<!-- <ul>
			{{$action := .action}}
			{{range $index,$nav := .navs}}
			<li {{if eq $index 0}} class="lbor" {{end}}>
				<a {{if eq $action .Enname}}class="hover"{{end}} href="/home/cn/{{.Enname}}">{{.Name}}</a>
			</li>
			{{end}}
		</ul> -->
	</div>