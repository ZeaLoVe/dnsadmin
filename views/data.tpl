<table border="1">
<tr>
<td>域名</td>
<td>Content</td>
<td>TTL</td>
<td>添加者</td>
<td>修改者IP</td>
<td>修改日期</td>
<td>是否停用</td>
<td>操作</td>
</tr>
{{range .s}}
<tr>
<td>{{.Name}}</td>
<td>{{.Content}}</td>
<td>{{.Ttl}}</td>
<td>{{.Auth}}</td>
<td>{{.Modifier_ip}}</td>
<td>{{.Change_date}}</td>
<td>{{.Disabled}}</td>
{{if .Disabled}}<td><a href="/enable/{{.Name}}">启用</a></td>
{{else}}<td><a href="/disable/{{.Name}}">停用</a>
{{end}}
</tr>
{{end}}
</table>