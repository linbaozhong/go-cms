{{define "modify"}}
<script type="text/javascript">
$(function(){
	// $('#defaultForm').isHappy({
 //          fields: {
 //            '#form-name': {
 //            	required: true,
 //            	message: '频道名称是必须填写的'
 //            }
 //       	}
 //    });
    // $('#defaultForm').on("submit", function(){
    // 	var $this = $(this),
    // 		$data = $this.serialize(),
    // 		$thisAction = $this.attr("action")
    // 		//console.log($thisAction)
    // 	$.post($thisAction, $data, function(data){
    // 		//console.log(data.Ok,data.Key,data.Data)
    // 		if( data.Ok ){
    // 			alert("提交成功!")
    // 			window.location.href='/admin/channel';
    // 		}else{
    // 			console.log(data.Key + "：" + data.Data)
    // 		}
    // 	})
    // })
    //验证并提交表单
    $('#defaultForm').validate({
    	errorClass:'checkMessage'
    	,submitHandler:function(form){
    		$(form).ajaxSubmit(function(data){
				if( data.Ok ){
	    			alert("提交成功!")
	    			window.location.href='/admin/account';
	    		}else{
	    			console.log(data.Key + "：" + data.Data)
	    		}
			});
    	}
	});
})
</script>
{{end}}
<div class="container">
	<div class="row">
		<div class="col-md-12">
			<div class="box">
				<div class="box-header">
					<h2>
						<span class="glyphicon glyphicon-edit"></span>
						新建账户
					</h2>
				</div>
				<div class="box-content">
					<div class="row">
						<div class="col-md-6">
							<form role="form" method="post" action="/admin/account/create" id="defaultForm">
								<fieldset>
									<div class="input-group">
										<span class="input-group-addon">
											<span class="glyphicon glyphicon-tag marginright10"></span>
											登录名
										</span>
										<input type="text" class="form-control" placeholder="登录名" name="loginname" id="form-name"></div>
									<div class="input-group">
										<span class="input-group-addon">
											<span class="glyphicon glyphicon-user marginright10"></span>
											密码
										</span>
										<input type="text" class="form-control" placeholder="password" name="password" value="{{.password}}"></div>
									<div class="input-group">
										<span class="input-group-addon">
											<span class="glyphicon glyphicon-calendar marginright10"></span>
											真实姓名
										</span>
										<input type="text" class="form-control" placeholder="真实姓名" name="relname"></div>
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