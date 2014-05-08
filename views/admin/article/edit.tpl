<div class="container">
	<div class="row">
		<div class="col-md-12">
			<div class="box">
				<div class="box-header">
					<h2>
						<span class="glyphicon glyphicon-edit"></span>
						发布内容
					</h2>
				</div>
				<div class="box-content">
					<div class="row">
						<div class="col-md-12">
							<form role="form" method="post" action="/admin/article/edit" id="defaultForm">
								<input type="hidden" name="id" value="{{.article.Id}}">
								<fieldset>
									<div class="input-group">
										<span class="input-group-addon">
											<span class="glyphicon glyphicon-tag marginright10"></span>
											所属频道
										</span>
										<select class="form-control" name="channelid">{{template "channels" .}}</select>
									</div>
									<div class="input-group">
										<span class="input-group-addon">
											<span class="glyphicon glyphicon-tag marginright10"></span>
											标题
										</span>
										<input type="text" class="form-control" placeholder="title" name="title" value="{{.article.Title}}" id="form-title" maxlength="150" required></div>
									<div class="input-group">
										<span class="input-group-addon">
											<span class="glyphicon glyphicon-tag marginright10"></span>
											副标题
										</span>
										<input type="text" class="form-control" placeholder="subtitle" name="subtitle" value="{{.article.Subtitle}}" maxlength="150"></div>

									<div class="form-group">
										<div class="input-group date form_datetime" data-date-format="yyyy-dd-mm HH:ii p" data-link-field="dtp_input1">
											<span class="input-group-addon">
												<span class="glyphicon glyphicon-calendar marginright10"></span>
												日期
											</span>
											<input class="form-control" size="16" type="text" id="form-date" name="published" value="" readonly>
										</div>
										<input type="hidden" id="dtp_input1" value="" />
									</div>
									<div class="input-group">
										<span class="input-group-addon">
											<span class="glyphicon glyphicon-user marginright10"></span>
											作者
										</span>
										<input type="text" class="form-control" placeholder="admin" name="author" value="{{.article.Author}}" maxlength="100"></div>
									<div class="input-group">
										<span class="input-group-addon">
											<span class="glyphicon glyphicon-bold marginright10"></span>
											关键字
										</span>
										<input type="text" class="form-control" placeholder="keywords" name="keywords"  value="{{.article.Keywords}}" maxlength="255"></div>
									<div class="form-group">
										<label class="control-label">内容简介</label>
										<div class="input-group">
											<textarea class="form-control" rows="3" name="intro"  maxlength="255">{{.article.Intro}}</textarea>
										</div>
									</div>
									<div class="form-group">
										<label class="control-label">内容</label>
										<div class="input-group">
											<textarea class="form-control" rows="6" name="content" id="form-content" style="width:100%;height:240px;" required>{{.article.Content}}</textarea>
										</div>
									</div>
									<div class="checkbox">
										<label>
											<input type="checkbox" name="status"  {{if eq .article.Status 1}}checked="checked"{{end}}>是否可用</label>
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