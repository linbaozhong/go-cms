<div class="container">
	<div class="row">
		<div class="col-md-12">
			<div class="box">
				<div class="box-header">
					<h2>
						<span class="glyphicon glyphicon-edit"></span>
						频道修改
					</h2>
				</div>
				<div class="box-content">
					<div class="row">
						<div class="col-md-6">

							<form role="form" method="post" action="/admin/channel/edit" id="defaultForm">
								<fieldset>
									<input type="hidden" name="id" value="{{.channel.Id}}">
									<div class="input-group">
										<span class="input-group-addon">
											<span class="glyphicon glyphicon-tag marginright10"></span>
											所属频道
										</span>
										<select class="form-control" name="pid">{{template "channels" .}}</select>
									</div>

									<div class="input-group">
										<span class="input-group-addon">
											<span class="glyphicon glyphicon-user marginright10"></span>
											频道名称
										</span>
										<input type="text" class="form-control" placeholder="频道名称" name="name" value="{{.channel.Name}}" required></div>
									<div class="input-group">
										<span class="input-group-addon">
											<span class="glyphicon glyphicon-calendar marginright10"></span>
											英文名称
										</span>
										<input type="text" class="form-control" placeholder="英文名称" name="enname" value="{{.channel.Enname}}"></div>
									<div class="input-group">
										<span class="input-group-addon">
											<span class="glyphicon glyphicon-calendar marginright10"></span>
											子项数量
										</span>
										<input type="text" class="form-control" placeholder="子项数量" name="children" value="{{.channel.Children}}"></div>
									<div class="input-group">
										<span class="input-group-addon">
											<span class="glyphicon glyphicon-bold marginright10"></span>
											类型
										</span>
										<select class="form-control" placeholder="类型" name="type">{{template "channelTypes" .}}</select>
									</div>
									<div class="checkbox">
										<label>
											<input type="checkbox" name="status" {{if eq .channel.Status 1}}checked="checked"{{end}}>是否可用</label>
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