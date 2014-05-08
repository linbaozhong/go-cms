<h2>
	后台-个人账户页-修改账户信息
</h2>
<div>
	<form id="updateprofile" action="">
		<label for="">真实姓名：<input type="text" name="relname" value="{{.user.Relname}}"></label>
		<input type="hidden" name="token" value="{{.token}}">
		<button type="submit">提交</button>
	</form>
</div>