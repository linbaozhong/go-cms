<h2>
	后台-个人账户页-修改密码
</h2>
<div>
	<h3>修改密码</h3>
	<form id="updatepassword" method="post">
		<label for="">旧密码：<input type="password" name="oldpassword"></label>
		<label for="">新密码：<input type="password" name="newpassword"></label>
		<label for="">重新输入密码：<input type="password" name="repassword"></label>
		<input type="hidden" name="token" value="{{.token}}">
		<button type="submit">提交</button>
		<button type="reset">重置</button>
	</form>
</div>