{{define "channels"}}
	{{range .chs}}
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
							<p class="pull-left titleformat">所属频道</p>
							<select class="form-control input-sm pull-left max-width100" id="channelId">{{template "channels" .}}</select>
						</div>
						<div class="col-md-8">
							<a href="/admin/article/create" class="btn btn-primary btn-sm pull-right">
								<span class="glyphicon glyphicon-plus glyphicontext"></span>
								新文章
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
							<th width="5%">选择</th>
							<th width="50%">标题</th>
							<th width="10%">发布日期</th>
							<th width="5%">顺序</th>
							<th width="10%">可用?</th>
							<th width="20%">操作</th>
						</tr>
					</thead>
					<tbody></tbody>
				</table>
				<div class="clearfix news-page">
					<ul class="pagination pull-right" id="newspage"></ul>
				</div>
			</div>
		</div>
	</div>
</div>
<script type="text/javascript">
//顺序
	$('#indexlist tbody').on('change','.sequence',function(){
		var o=$(this),
		s=o.val();
		if (valid.isInt(s)) {
			sq=parseInt(s);
			$.post(cmsapi.articleSequence,{id:o.data('id'),sq:sq},function(json){
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
$('#indexlist tbody').on('change','.status',function(){
	var o=$(this);
	$.post(cmsapi.articleStatus,{id:o.data('id')},function(json){
		if (!json.Ok) {
			o.prop("checked",!o.is(':checked'))
		};
	})
})


$(function(){
	var index = 1, // 当前页
	size = 10 // 每页
	count = 0 //总记录数

	$("#channelId").change(function(){
		var $this = $(this),
			channelId = $this.val();
			
		getNewsList(channelId,index,size);
	})
	
	getNewsList(0,index,size)
	//分页新闻
	$("#newspage").on("click", "li", function(){
		var index = $(this).data("id"),
			channelId = $("#channelId").val();
		if (index>0) {
			if($(this).hasClass("pagemore")){
				$("#newspage").empty().append(showPages(count,index,size));
			}else{
				getNewsList(channelId,index,size)
			}		
		}

		return false
	})
	//格式化日期
	function formatDate(date){
		var o = new Date(date);
		return o.Format('yyyy-MM-dd hh:mm:ss')
	}
	//分页
	function showPages(count,index,size){
		var page = Math.ceil( count / size ), // 页数
				pageHtml = "";

		if (page > 0) {
	        //页码按钮数
	        var buttons = 5;
	        var start = 1;
	        if (index >= buttons)
	        {
	            if (page - index > buttons)
	            {
	                start = index - 2;
	            }
	            else if (page - index < buttons)
	            {
	                start = page - buttons;
	            }
	            else if (page - index == buttons)
	            {
	                start = page - buttons - 1;
	            }
	        }
	        if (index > 1)
            {
                pageHtml += '<li data-id="1" title="首页"><a href="#">|<</a></li>'
                pageHtml += '<li data-id="'+(index - 1)+'" title="上一页"><a href="#"><</a></li>'
            }
            else
            {
                pageHtml += '<li disabled="disabled" data-id="0" title="首页"><a href="javascript:;">|<</a></li>'
                pageHtml += '<li disabled="disabled" data-id="0" title="上一页"><a href="javascript:;"><</a></li>'
            }

			for( var i = start; i < page + 1; i++ ){ 
				if (i - start > 5)
                {
                    break;
                }
                if ((page > buttons) && (index >= buttons && i == start) || (index <= page - buttons && i == start + buttons))
                {
                	pageHtml += '<li class="pagemore" data-id="' + i + '"><a href="#" >...</a></li>'
                }
                else
                {
                	if(i==index){
                		pageHtml += '<li class="active" data-id="' + i + '"><a href="#">' + i + '</a></li>'
                	}else{
                		pageHtml += '<li data-id="' + i + '"><a href="#">' + i + '</a></li>'
                	}
                }
				
			}

			if (index < page)
            {
                pageHtml += '<li data-id="'+(index + 1)+'" title="下一页"><a href="#">></a></li>'
                pageHtml += '<li title="末页" data-id="'+page+'"><a href="#">>|</a></li>'
            }
            else
            {
                pageHtml += '<li disabled="disabled" title="下一页" data-id="0"><a href="javascript:;">></a></li>'
                pageHtml += '<li disabled="disabled" title="末页" data-id="0"><a href="javascript:;">>|</a></li>'
            }
		}
		return pageHtml;
	}
	// 获取新闻
	function getNewsList(channelId,index,size){
		html = ""
		$.getJSON(cmsapi.getArticles, {channelid:channelId,index:index,size:size,times:Math.random()}, function(o) {
			if( o.Ok ){
				count = parseInt(o.Key);
				$("#newspage").empty().append(showPages(count,index,size));
				$("#indexlist tbody").empty()
				//console.log(NewsJson[1].Title)
				$.each(o.Data, function(i,o){
					var pubDate = formatDate(o.Published)
					html += "<tr><td><div class='checkbox'><label><input type='checkbox' name='ids' value='" + o.Id 
						+ "'></label></div></td><td><a href='/admin/article/edit/" + o.Id + "'>" + o.Title + "</a></td><td>" 
						+ pubDate + "</td><td><input type='text' data-id='" + o.Id + "' class='sequence' value='" + o.Sequence + "'></td>"
					if( o.Status ){
						html += "<td><input type='checkbox' data-id='" + o.Id + "' class='status' checked='checked' />"
					}else{
						html += "<td><input type='checkbox' data-id='" + o.Id + "' class='status' />"
					}
					html += "<td><div class='btn-group btn-group-sm'><a href='/admin/image/index/" + o.Id + "' class='btn btn-default' title='插入图片'><span class='glyphicon glyphicon-edit'></span>插入图片</a><a href='/admin/article/edit/" + o.Id + "' class='btn btn-default' title='修改'><span class='glyphicon glyphicon-edit'></span>修改</a><a href='/admin/article/delete/" + o.Id + "' class='btn btn-default form-del' title='删除'><span class='glyphicon glyphicon-edit'></span>删除</a></div></td></tr>"
					//console.log(html)
				})
				$("#indexlist tbody").append(html)
			}
		});
	}
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
})
</script>