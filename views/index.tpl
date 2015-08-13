<!DOCTYPE html>

<html>
<head>
  <title>{{.Website}}</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<link rel="stylesheet" type="text/css" href="http://www.jeasyui.com/easyui/themes/default/easyui.css">
    <link rel="stylesheet" type="text/css" href="http://www.jeasyui.com/easyui/themes/icon.css">
    <link rel="stylesheet" type="text/css" href="http://www.jeasyui.com/easyui/themes/color.css">
    <link rel="stylesheet" type="text/css" href="http://www.jeasyui.com/easyui/demo/demo.css">
    <script type="text/javascript" src="http://code.jquery.com/jquery-1.6.min.js"></script>
    <script type="text/javascript" src="http://www.jeasyui.com/easyui/jquery.easyui.min.js"></script>
	<script type="text/javascript">
function validate_form(thisform){
with (thisform){
	if(domain.value ==""||domain==null){	
		alert('domain is null!');
		domain.focus();
 		return false;
	}
	if(content.value=="" ||content==null){	
	    alert('content is a ip or cname !');
		content.focus();
 		return false;
	}
  }
}
</script>
</head>

<body>
  <header>
    <h1 class="logo">Welcome to {{.Website}}</h1>
  </header>
      <p>ps:修改DNS记录就是添加一个相同域名的记录</p>
	<div class="easyui-panel" title="添加记录">
	<div style="padding:10px 30px 20px 30px">
      <form name="input" action="/insert" method="post" onsubmit="return validate_form(this)">
	<table cellpadding="5">
<tr>
<td>域名:</td>
<td><input class="easyui-textbox" type="domain" name="domain" data-options="required:true"/></td>
</tr>
<tr>
<td>IP地址:</td>
<td><input class="easyui-textbox" type="content" name="content" data-options="required:true"/></td>
</tr>
<tr>
<td>时效:</td>
<td><input class="easyui-textbox" type="ttl" name="ttl" /></td>
</tr>
<tr>
<td>添加者:</td>
<td><input class="easyui-textbox" type="auth" name="auth" data-options="required:true"/></td>
</tr>
<tr>
<input class="easyui-linkbutton" type="submit" value="添加记录" /><tr></tr>
</tr>
</table>
</form>
</div>
</div>
     <div class="easyui-panel" title="查询记录">
	<div style="padding:10px 30px 20px 30px">
		
      <form name="form2" method="post" action="/search">
        <table cellpadding="5">
		<tr>
		<td><label>
          <input type="text" name="searchDNS" id="searchDNS">
        </label></td>
		<td><label>
          <input class="easyui-linkbutton" type="submit" name="search" id="search" value="查询记录">
        </label></td>
		</tr>
		</table>       
      </form>
	{{.LayoutContent}}
</tr>     
	</div>
	</div>
</body>
</html>
