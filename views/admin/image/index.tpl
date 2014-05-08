<div class="container">
	<div class="row">
		<div class="col-md-12">
			<div class="panel panel-default">
				<div class="panel-heading clearfix">
					<div class="row">
						<div class="col-md-4">
							<p class="pull-left titleformat">所属文章</p>
							<span class="pull-left titleformat">{{.article.Title}}</span>
						</div>
						<div class="col-md-8">
							<a href="/admin/image/create/{{.article.Id}}" class="btn btn-primary btn-sm pull-right">
								<span class="glyphicon glyphicon-plus glyphicontext"></span>
								新图片
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
				<table class="table table-striped table-hover imglist" id="indexlist">
					<thead>
						<tr>
							<th>选择</th>
							<th>图片</th>
							<th>顺序</th>
							<th>可用?</th>
							<th>操作</th>
						</tr>
					</thead>
					<tbody>
						{{ range .images}}
						<tr>
							<td>
								<div class="checkbox">
									<label>
										<input type="checkbox" name="ids" value="{{.Id}}"></label>
								</div>
							</td>

							<td>
								<img src="{{htmlquote .Path}}" alt="{{.Title}}"  />
							</td>
							<td>
								<input type="text" data-id="{{.Id}}" class="sequence" value="{{ .Sequence }}"></td>
							<td>
								<input type="checkbox" data-id="{{.Id}}" class="status" {{ if eq .Status 1 }}checked="checked"{{ end }}></td>
							<td>
								<div class="btn-group btn-group-sm">
									<a href="/admin/image/edit/{{ .Id }}" class="btn btn-default" title="修改">
										<span class="glyphicon glyphicon-edit"></span>
										修改
									</a>
									<a href="/admin/image/delete/{{ .Id }}" class="btn btn-default form-del" title="删除">
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
//顺序
	$('.sequence').on('change',function(){
		var o=$(this),
		s=o.val();

		if (valid.isInt(s)) {
			sq=parseInt(s);
			
			$.post(cmsapi.imageSequence,{id:o.data('id'),sq:sq},function(json){
				if (json.Ok) {
					o.val(sq);
				}else{
					o.val(s);
				};
			})
		}else{
			o.val(0)
		};
	});
	//状态
$('.status').on('change',function(){
	var o=$(this);
	$.post(cmsapi.imageStatus,{id:o.data('id')},function(json){
		if (!json.Ok) {
			o.prop("checked",!o.is(':checked'))
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