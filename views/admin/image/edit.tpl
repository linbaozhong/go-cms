<div class="container">
	<div class="row">
		<div class="col-md-12">
			<div class="box">
				<div class="box-header">
					<h2>
						<span class="glyphicon glyphicon-edit"></span>
						编辑图片
					</h2>
				</div>
				<div class="box-content">
					<div class="row">
						<div class="col-md-6">
							<form role="form" method="post" action="/admin/image/edit" enctype="multipart/form-data" id="imgForm">
								<input type="hidden" name="id" value="{{.image.Id}}">
								<input type="hidden" name="articleid" value="{{.image.Articleid}}">
								<fieldset>
									<div class="input-group">
										<span class="input-group-addon">
											<span class="glyphicon glyphicon-edit marginright10"></span>
											图片标题
										</span>
										<input type="text" class="form-control" placeholder="title" name="title" value="{{.image.Title}}"></div>
									<div class="input-group">
										<span class="input-group-addon">
											<span class="glyphicon glyphicon-edit marginright10"></span>
											链接地址
										</span>
										<input type="text" class="form-control" placeholder="http://" name="url" value="{{.image.Url}}"></div>
									<div class="input-group">
										<input type="text" class="form-control" id="fileupload">
										<span class="input-group-btn" style="overflow:hidden;">
											<button class="btn btn-default" type="button">
												上传图片
												</button><input type="file" style="position:absolute;right:0;top:0;padding:7px 12px;cursor:pointer;opacity:0;" onchange="getImage(this);" name="file">
										</span>
									</div>
									<div class="form-group">
										<div class="row">
											<div class="col-md-12">
												<div class="thumbnail">
													<img src="{{.image.Path}}" alt="{{.image.Title}}" width="100" height="100" />
												</div>
											</div>
										</div>
									</div>
									<div class="checkbox">
										<label>
											<input type="checkbox" name="status"  {{if eq .image.Status 1}}checked="checked"{{end}}>是否可用</label>
										<input type="hidden" name="token" value="{{.token}}"></div>
									<button type="submit" class="btn btn-primary" style="width:120px;">提交内容</button>
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