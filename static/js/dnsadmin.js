var g_url = "/domains";
var isModify;
var domain_patt = new RegExp("([0-9a-z][0-9a-z-]+)([.][0-9a-z][0-9a-z-]+){3,10}");

function validDomain(domain){
	return domain_patt.test(domain);
}

function fleshUrl(){
	var name = $('#filter_name').val();
	var content = $('#filter_content').val();
	var auth = "";
	if ( $('#mine').is(':checked') == true ){
		auth = $('.user_name').attr("id") 
	}
	g_url = "/domains?name=" + name +"&content=" + content + "&auth=" +auth
		
	$('#dg').datagrid({
   	 	url:g_url,
		pageNumber:1,
		pageSize:20,
		method:"GET",
		rownumbers:"true",
		fitColumns:"true",
		singleSelect:"true",
		pagination:"true",
		toolbar:"#toolbar",
    	columns:[[
        	{field:'Name',title:'域名',width:50},
        	{field:'Content',title:'IP或者别名',width:50},
			{field:'Ttl',title:'有效时间TTL',width:30},
        	{field:'Disabled',title:'是否停用(0为在用 1为停用)',width:20}
    	]]
	});
	$()
	$('#dg').datagrid('reload'); 
}
	
function newDomain(){
	$('#dlg').dialog('open').dialog('setTitle','添加域名');
	isModify = false;
	$('#domain_name').attr("disabled",false);
	$('#fm').form('clear');
}

function editDomain(){
	var row = $('#dg').datagrid('getSelected');	
	if (row){
		row.name = row.Name
		row.content = row.Content
		row.ttl = row.Ttl
		row.auth = row.Auth
		$('#dlg').dialog('open').dialog('setTitle','修改域名');
		isModify =true
		$('#domain_name').attr("disabled",true);
		$('#fm').form('load',row);
	}
}

function saveDomain(){
	var name = $("#fm input[name='name']").val();
	var auth = $('.user_name').attr("id")
	$('#new_domain_auth').attr("value",auth)
	var url = "/domains/" + name
	$('#fm').form('submit',{
		url: url,
		onSubmit: function(){
			if( isModify  == false ){
				if( validDomain(name) ==false ){
					alert("输入的域名不符合规范！请参考域名规范\n实例名.环境标识.实例标识.机房标识.服务标识.sdp");
					return false;
				}
			}else{
				return true;//add validate function here
			}
		},
		success: function(result){
			var result = eval('('+result+')');
			if (result.code == "ok"){
				$('#dlg').dialog('close');		// close the dialog
				$('#dg').datagrid('reload');	// reload the user data		
			} else {
				$.messager.show({
					title: 'Error',
					msg: result.message
				});
			}
		}
	});
}
function destroyDomain(){
	var row = $('#dg').datagrid('getSelected');
	if (row){
		$.messager.confirm('Confirm','确认删除域名?',function(r){
			if (r){
				var url = "/domains/" + row.Name
				$.ajax({
					url:url,
					type:'DELETE',
					success:function(result){
						if (result.code == "ok"){
							$('#dg').datagrid('reload');
						} else {
							$.messager.show({
								title: 'Error',
								msg: result.message
							});
						}
					}
				});
			}
		});
	}
}
function enableDomain(){
	var row = $('#dg').datagrid('getSelected');
	if (row){
		$.messager.confirm('Confirm','确认启用域名?',function(r){
			if (r){
				var url = "/enable/" + row.Name
				$.ajax({
					url:url,
					type:'POST',
					success:function(result){
						if (result.code == "ok"){
							$('#dg').datagrid('reload');
						} else {
							$.messager.show({
								title: 'Error',
								msg: result.message
							});
						}
					}
				});
			}
		});
	}
}
function disableDomain(){
	var row = $('#dg').datagrid('getSelected');
	if (row){
		$.messager.confirm('Confirm','确认停用域名?',function(r){
			if (r){
				var url = "/disable/" + row.Name
				$.ajax({
					url:url,
					type:'POST',
					success:function(result){
						if (result.code == "ok"){
							$('#dg').datagrid('reload');
						} else {
							$.messager.show({
								title: 'Error',
								msg: result.message
							});
						}
					}
				});
			}
		});
	}
}