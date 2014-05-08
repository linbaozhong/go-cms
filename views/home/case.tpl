{{$action := .action}}
<div class="wrap">
	<div class="center">
		<div class="title_L">
			<h3>
				<!-- <img src="/static/images/{{$action}}.jpg"/> -->
				{{.channel.Name}}
			</h3>
			<ul>
				{{$channelid := .currentChannelId}}
				{{range $index,$channel := .cases}}
				<li>
					<span style="margin-left:{{.Level}}em;display:inline-block;"></span>
					<a href="/home/cn/{{$action}}/{{.Id}}" 
							{{if eq .Id $channelid}}
								class="hover"
							{{else if and (eq $index 0) (eq $channelid 0)}}
								class="hover"
							{{end}} >{{.Name}}</a>
				</li>
				{{end}}
			</ul>
		</div>
		<div class="title_R" style="width:740px;">
			<div class="title_cen" style="width:auto;position: relative;min-height: 580px;">
				{{if gt (len .images) 0}}
				<div class="silder-wrap">
					<div class="silder">
						<ul>
							{{range .images}}
							<li>
								<a href="{{.Url}}" target="_blank" {{if eq .Url ""}} onclick="return false;" {{end}}>
									<img src="{{.Path}}" width="674" height="506" />
								</a>
								<p class="silder-title">{{.Title}}</p>
							</li>
							{{end}}
						</ul>
					</div>
					<div class="silder-page">
						<a href="javascript:" class="prev">&lt;</a>
						<a href="javascript:" class="next">&gt;</a>
					</div>
					<div class="num"></div>
				</div>
				{{end}}
				<!-- {{.article.Title}} -->
				{{.article.Content | str2html}}
						
				{{if gt .pagination.Size 1}}
				<div class="skipUrl">
					{{if gt .pagination.Prev 0}}
					<a href="/home/cn/{{$action}}/{{$channelid}}/0/{{.pagination.Prev}}">
						<span class="skipL"></span>
					</a>
					{{end}}
					<div class="skipC">
						{{$articleid := .currentArticleId}}
						{{$pagindex := .pagination.Index}}
						{{$pageadd := (Multiply .pagination.Prev .pagination.Size)}}
						
						{{range $index,$article := .articles}}
						<a {{if eq $articleid .Id}}class="hover"{{end}} 
						href="/home/cn/{{$action}}/{{$channelid}}/{{.Id}}/{{$pagindex}}">{{Plus $index (Plus $pageadd 1)}}</a>
						{{end}}
					</div>
					{{if gt .pagination.Next 0}}
					<a href="/home/cn/{{$action}}/{{$channelid}}/0/{{.pagination.Next}}">
						<span class="skipR"></span>
					</a>
					{{end}}
				</div>
				{{end}}
			</div>
		</div>
	</div>
</div>
<script type="text/javascript" src="/static/js/jquery-1.10.2.min.js"></script>
<script>
$(function(){
	var page = 1,
		speed = 500,
		box = $(".silder-wrap"),
		boxW = box.width(),
		wrap = box.find(".silder"),
		numBox = $(".num"),
		num = wrap.find("li").length;
	wrap.width(boxW * num)
	num == 1 ? $(".silder-page,.num").hide() : $(".silder-page,.num").show()
	for(var i = 0; i < num; i++){ numBox.append("<a href='#' class='on'>"+ i +"</a>") }
	box.find(".num a").eq(0).addClass("curren")
	box.on("click",".prev,.next,.num a",function(){
		var $this = $(this),
			thisClass = $this.attr("class")
		if(!wrap.is(":animated")) {
			if( thisClass == "next" ){
				if( page == num ){ 
					wrap.animate({ marginLeft:0 }, speed, "linear")
					page = 1
				}else{
					wrap.animate({ marginLeft:"-=" + boxW }, speed, "linear")
					page++
				}
			}else if( thisClass == "prev" ){
				if( page == 1 ){
					wrap.animate({ marginLeft: -boxW * (num - 1) +'px' }, speed, "linear")
					page = num
				}else{
					wrap.animate({ marginLeft:"+=" + boxW }, speed, "linear")
					page--
				}
			}else if( thisClass == "on" ){
				var $this = $(this),
					index = $this.index() + 1
				$this.addClass("curren").siblings().removeClass("curren");
				wrap.animate({ marginLeft: "-=" + (index - page) * boxW + 'px' }, speed, "linear")
				page = index++
				return false
			}
			box.find(".num a").eq(page - 1).addClass("curren").siblings().removeClass("curren");
		}
		
	})
})
</script>