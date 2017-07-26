<!DOCTYPE html>

<html>
<head>
  <title>{{.Website}}</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<link rel="stylesheet" type="text/css" href="../static/easyui/themes/default/easyui.css">
    <link rel="stylesheet" type="text/css" href="../static/easyui/themes/icon.css">
    <link rel="stylesheet" type="text/css" href="../static/easyui/themes/color.css">
    <link rel="stylesheet" type="text/css" href="../static/easyui/demo/demo.css">
    <script type="text/javascript" src="../static/js/jquery-1.6.min.js"></script>
    <script type="text/javascript" src="../static/easyui/jquery.easyui.min.js"></script>
	<script type="text/javascript" src="../static/js/dnsadmin.js"></script>
</head>

<body>
  <header>
    <h1 class="user_name" id="{{.UserName}}" >Welcome to {{.Website}} , {{.UserName}} !</h1>
  </header>
	
<div id="toolbar" class="easyui-panel" title="域名管理页面">
	<table class="altrowstable" id="domain_filters" cellpadding="5">
	<tr>
	<td>通过域名筛选:</td>
	<td><input id="filter_name" type="name" name="name" onblur="fleshUrl()"/></td>
	<td>通过IP筛选:</td>
	<td><input id="filter_content" type="content" name="content" onblur="fleshUrl()"/></td>
	<td><span>&nbsp&nbsp只看自己的</span></td>
    <td><input type="checkbox" name="mine" id="mine" onblur="fleshUrl()"></td>     
	</tr>
	
	</table>
    <a href="#" class="easyui-linkbutton" id="toolbar_new" iconCls="icon-add" plain="true" onclick="newDomain()">添加</a>
    <a href="#" class="easyui-linkbutton" id="toolbar_edit" iconCls="icon-edit" plain="true" onclick="editDomain()">编辑</a>
    <a href="#" class="easyui-linkbutton" id="toolbar_remove" iconCls="icon-remove" plain="true" onclick="destroyDomain()">删除</a>
	<a href="#" class="easyui-linkbutton" id="toolbar_enable" iconCls="icon-redo" plain="true" onclick="enableDomain()">启用</a>
	<a href="#" class="easyui-linkbutton" id="toolbar_disable" iconCls="icon-cancel" plain="true" onclick="disableDomain()">停用</a>
</div>
<table id="dg"></table>

<div id="dlg" class="easyui-dialog" style="width:350px;height:200px;padding:10px 20px"
		closed="true" buttons="#dlg-buttons">
	<div class="ftitle"></div>
	<form id="fm" method="post">
		<div class="fitem">
			<label>域名:&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp</label>
			<input id="domain_name" name="name" class="easyui-validatebox" required="true" style="width:200px">
		</div>
		<div class="fitem">
			<label>IP:&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp</label>
			<input name="content" class="easyui-validatebox" required="true" style="width:200px">
		</div>
		<div class="fitem">
			<label>有效时间: </label>
			<input name="ttl" style="width:200px">
		</div>
		<div class="fitem">
			<input id="new_domain_auth" name="auth" style="width:200px" type="hidden">
		</div>
		<div class="fitem">
			<input id="new_domain_disabled" name="auth" style="width:200px" type="hidden">
		</div>
	</form>
</div>
<div id="dlg-buttons">
	<a href="#" class="easyui-linkbutton" iconCls="icon-ok" onclick="saveDomain()">Save</a>
	<a href="#" class="easyui-linkbutton" iconCls="icon-cancel" onclick="javascript:$('#dlg').dialog('close')">Cancel</a>
</div>

</tr>     
	</div>
	</div>
</body>
</html>
