<div class="container">
	<div class="row">
		<div class="col-md-12">
			<div class="box">
				<div class="box-header">
					<h2>
						<span class="glyphicon glyphicon-edit"></span>
						修改账户
					</h2>
				</div>
				<div class="box-content">
					<div class="row">
						<div class="col-md-6">
							<form role="form" method="post" action="/admin/account/edit" id="defaultForm">
								<fieldset>
									<div class="input-group">
										<input type="hidden" name="id" value="{{.user.Id}}">
										<span class="input-group-addon">
											<span class="glyphicon glyphicon-tag marginright10"></span>
											登录名
										</span>
										<input type="text" class="form-control" placeholder="loginname" name="loginname" value="{{.user.Loginname}}" id="form-name"></div>
									<div class="input-group">
										<span class="input-group-addon">
											<span class="glyphicon glyphicon-calendar marginright10"></span>
											真实姓名
										</span>
										<input type="text" class="form-control" placeholder="relname" name="relname" value="{{.user.Relname}}"></div>
									<div class="checkbox">
										<label>
											<input type="checkbox">是否可用</label>
										<input type="hidden" name="token" value="{{.token}}"></div>
									<button type="submit" class="btn btn-primary" style="width:120px;">提交</button>
								</fieldset>
							</form>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
</div>
{{template "modify" .}}