<!DOCTYPE HTML>
<html>
<head>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<title>{{.page.Title}}</title>
	<link href="/static/css/bootstrap.min.css" type="text/css" rel="stylesheet"/>
</head>
<body>
	<div class="container">
		<div class="box" style="width:500px;margin:0 auto;margin-top:200px;">
			<div class="box-header">
				<h1 style="text-align:center">Orange后台登录</h1>
			</div>
			<div class="box-content">
				<div class="row">
					<div class="col-md-12">
						<form action="" method="post">
							<div class="input-group">
								<span class="input-group-addon">
									<span class="glyphicon glyphicon-tag marginright10"></span>
									用户
								</span>
								<input type="text" class="form-control" placeholder="username" name="loginname" value="{{.loginname}}">
								<input type="hidden" class="form-control" placeholder="" name="returnurl" value="{{.returnurl}}">
							<input type="hidden" class="form-control" placeholder="" name="token" value="{{.token}}"></div>
						<div class="input-group">
							<span class="input-group-addon">
								<span class="glyphicon glyphicon-tag marginright10"></span>
								密码
							</span>
							<input type="password" class="form-control" placeholder="password" name="password" value=""></div>
						<div class="checkbox">
							<label>
								<input type="checkbox" name="always">记住我的登录状态</label>
						</div>
						<button type="submit" class="btn btn-primary" style="width:120px;">登录</button>
					</form>
				</div>
			</div>
		</div>
	</div>
</div>
</body>
</html>