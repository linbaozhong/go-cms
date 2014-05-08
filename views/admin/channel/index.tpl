{{define "channels"}}
	{{range .chs}}
<option value="{{.Value}}" {{if .Selected}}selected="selected"{{end}}>{{str2html .Key}}</option>
{{end}}
{{end}}

{{define "channelTypes"}}
	{{range .types}}
<option value="{{.Value}}" {{if .Selected}}selected="selected"{{end}}>{{str2html .Key}}</option>
{{end}}
{{end}}
<div class="container">
	<div class="row">
		<div class="col-md-12">
			<div class="panel panel-default">
				<div class="panel-heading clearfix">
					<div class="row">
						<div class="col-md-4">
							<p class="pull-left titleformat">分类</p>
							<select class="form-control input-sm pull-left max-width100">{{template "channelTypes" .}}</select>
						</div>
						<div class="col-md-8">
							<a href="/admin/channel/create" class="btn btn-primary btn-sm pull-right">
								<span class="glyphicon glyphicon-plus glyphicontext"></span>
								新建频道
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
							<th>名称</th>
							<th>分页记录数</th>
							<th>类型</th>
							<th>顺序</th>
							<th>可用?</th>
							<th>操作</th>
						</tr>
					</thead>
					<tbody>
						{{ range .channels }}
						<tr>
							<td>
								<div class="checkbox">
									<label>
										<input type="checkbox" name="ids" value="{{.Id}}"></label>
								</div>
							</td>
							<td>
								<a href="/admin/channel/edit/{{ .Id }}">{{str2html (Indent .Name .Level)}}</a>
							</td>
							<td>
								<input type="text" data-id="{{.Id}}" size="4" class="children" value="{{ .Children }}"></td>
							<td>{{if eq .Type 0}}导航{{else}}文章{{end}}</td>
							<td>
								<input type="text" data-id="{{.Id}}" size="4" class="sequence" value="{{ .Sequence }}"></td>
							<td>
								<input type="checkbox" data-id="{{.Id}}" class="status" {{ if eq .Status 1}}checked="checked"{{ end }}></td>
							<td>
								<div class="btn-group btn-group-sm">
									<a href="/admin/channel/edit/{{ .Id }}" class="btn btn-default" title="修改">
										<span class="glyphicon glyphicon-edit"></span>
										修改
									</a>
									<a href="/admin/channel/delete/{{ .Id }}" class="btn btn-default form-del" title="删除">
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
//子项限制
	$('.children').on('change',function(){
		var o=$(this),
		s=o.val();

		if (valid.isInt(s)) {
			sq=parseInt(s);
			
			$.post(cmsapi.channelChildren,{id:o.data('id'),sq:sq},function(json){
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
//顺序
	$('.sequence').on('change',function(){
		var o=$(this),
		s=o.val();

		if (valid.isInt(s)) {
			sq=parseInt(s);
			
			$.post(cmsapi.channelSequence,{id:o.data('id'),sq:sq},function(json){
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
	$.post(cmsapi.channelStatus,{id:o.data('id')},function(json){
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