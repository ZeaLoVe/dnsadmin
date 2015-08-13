<!-- Javascript goes in the document HEAD -->
<script type="text/javascript">
function altRows(id){
	if(document.getElementsByTagName){  
		
		var table = document.getElementById(id);  
		var rows = table.getElementsByTagName("tr"); 
		 
		for(i = 0; i < rows.length; i++){          
			if(i % 2 == 0){
				rows[i].className = "evenrowcolor";
			}else{
				rows[i].className = "oddrowcolor";
			}      
		}
	}
}
window.onload=function(){
	altRows('alternatecolor');
}
</script>

<style type="text/css">
table.altrowstable {
	font-family: verdana,arial,sans-serif;
	font-size:11px;
	color:#333333;
	border-width: 1px;
	border-color: #a9c6c9;
	border-collapse: collapse;
}
table.altrowstable th {
	border-width: 1px;
	padding: 8px;
	border-style: solid;
	border-color: #a9c6c9;
}
table.altrowstable td {
	border-width: 1px;
	padding: 8px;
	border-style: solid;
	border-color: #a9c6c9;
}
.oddrowcolor{
	background-color:#d4e3e5;
}
.evenrowcolor{
	background-color:#c3dde0;
}
</style>
<table class="altrowstable">
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