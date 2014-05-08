<div class="container">
	<div class="row">
		<div class="col-md-12">
			<div class="panel panel-default">
				<div class="panel-heading clearfix">
					<div class="row">
						<div class="col-md-4">
							<!-- <p class="pull-left titleformat">分类</p>
						<select class="form-control input-sm pull-left max-width100">
							<option value="0">导航</option>
							<option value="1">文章</option>
						</select>
						-->
					</div>
					<div class="col-md-8">
						<a href="/admin/account/create" class="btn btn-primary btn-sm pull-right">
							<span class="glyphicon glyphicon-plus glyphicontext"></span>
							新建账户
						</a>
						<div class="input-group input-group-sm pull-right max-width400" style="margin-bottom:0;">
							<input type="text" class="form-control">
							<span class="input-group-btn">
								<button class="btn btn-default" type="button">
									<span class="glyphicon glyphicon-search"></span>
								</button>
							</span>
						</div>
					</div>
				</div>
			</div>
			<table class="table table-striped table-hover" id="indexlist">
				<thead>
					<tr>
						<th>选择</th>
						<th>账户名</th>
						<th>姓名</th>
						<th>角色</th>
						<th>可用?</th>
						<th>操作</th>
					</tr>
				</thead>
				<tbody>
					{{ range .accounts}}
					<tr>
						<td>
							<div class="checkbox">
								<label>
									<input type="checkbox" name="ids" value="{{.Id}}"></label>
							</div>
						</td>
						<td>
							<a href="/admin/account/edit/{{ .Id }}">{{.Loginname}}</a>
						</td>
						<td>{{.Relname}}</td>
						<td>
							{{if eq .Role 0}}读者{{else if eq .Role 1}}编辑{{else}}管理员{{end}}
						</td>
						<td>
							<input type="checkbox" class="status" data-id="{{.Id}}" {{ if eq .Status 1 }}checked="checked"{{ end }}></td>
						<td>
							<div class="btn-group btn-group-sm">
								<a href="/admin/account/edit/{{ .Id }}" class="btn btn-default" title="修改">
									<span class="glyphicon glyphicon-edit"></span>
									修改
								</a>
								<a href="/admin/account/delete/{{ .Id }}" class="btn btn-default form-del" title="删除">
									<span class="glyphicon glyphicon-edit"></span>
									删除
								</a>
							</div>
						</td>
					</tr>
					{{end}}
				</tbody>
			</table>
			
		</div>
	</div>
</div>
</div>
<script type="text/javascript">
//状态
$('.status').on('change',function(){
	var o=$(this);
	$.post(cmsapi.accountStatus,{id:o.data('id')},function(json){
		if (!json.Ok) {
			o.prop("checked",!o.is(':checked'))
			alert(json.Data)
		};
	})
})
//删除
	$("#indexlist").on("click",".form-del",function(){
		if(confirm("你确定要删除吗?")){
			var $this = $(this),$thisHref = $this.attr("href");
			$.get($thisHref,{testdata:Math.random()},function(o){
				if( o.Ok ){
					$this.parents("tr").remove()
					alert("删除成功！")
				}else{
					alert(o.Data)
				}
			})
		}
		return false
	})
</script>