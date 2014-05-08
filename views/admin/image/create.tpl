{{define "modify"}}
<script type="text/javascript">
$(function(){
	$('#imgForm').ajaxForm(function(data){
		if (data.Ok) {
	        window.location.href = "/admin/image/index/"+$('#imgForm').find('input[name=articleid]').val();
        } else {
            alert(data.Data);
        }
	});	
})

function getImage(file){

	if (file.files && file.files[0] && window.FileReader) {

        var _img = file.files[0];

        var reader = new FileReader();

        reader.onload = function (evt) {
            $('.thumbnail img').attr('src',evt.target.result);
        }
        reader.readAsDataURL(_img);    
    }   
}
</script>
{{end}}

<script type="text/javascript" charset="utf-8" src="/static/js/ajaxfileupload.js"></script>
<div class="container">
	<div class="row">
		<div class="col-md-12">
			<div class="box">
				<div class="box-header">
					<h2>
						<span class="glyphicon glyphicon-edit"></span>
						发布图片
					</h2>
				</div>
				<div class="box-content">
					<div class="row">
						<div class="col-md-6">
							<form role="form" method="post" action="/admin/image/create" enctype="multipart/form-data" id="imgForm">
								<input type="hidden" name="id" value="">
								<input type="hidden" name="articleid" value="{{.articleid}}">
								<fieldset>
									<div class="input-group">
										<span class="input-group-addon">
											<span class="glyphicon glyphicon-edit marginright10"></span>
											图片标题
										</span>
										<input type="text" class="form-control" placeholder="title" name="title" value=""></div>
									<div class="input-group">
										<span class="input-group-addon">
											<span class="glyphicon glyphicon-edit marginright10"></span>
											链接地址
										</span>
										<input type="text" class="form-control" placeholder="http://" name="url" value=""></div>
									<div class="input-group">
										<input type="text" class="form-control" id="fileupload">
										<span class="input-group-btn" style="overflow:hidden;">
											<button class="btn btn-default" type="button">
												上传图片</button>
												<input type="file" style="position:absolute;right:0;top:0;padding:7px 12px;cursor:pointer;opacity:0;" onchange="getImage(this);" id="file" name="file" id="form-file">
										</span>
									</div>
									<div class="form-group">
										<div class="row">
											<div class="col-md-12">
												<div class="thumbnail">
													<img src="">
												</div>
											</div>
										</div>
									</div>
									<div class="checkbox">
										<label>
											<input type="checkbox" name="status">是否可用</label>
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