<ul>
	{{$action := .action}}
	{{range $index,$nav := .navs}}
	<li {{if eq $index 0}} class="lbor" {{end}}>
		<a {{if eq $action .Enname}}class="hover"{{end}} href="/home/cn/{{.Enname}}">{{.Name}}</a>
	</li>
	{{end}}
</ul>