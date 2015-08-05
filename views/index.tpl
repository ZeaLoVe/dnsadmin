<!DOCTYPE html>

<html>
<head>
  <title>{{.Website}}</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
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
      <div class="author">
      <p>&nbsp;</p>
      <form name="input" action="/insert" onsubmit="return validate_form(this)" method="post">
域&nbsp;名:
<input type="domain" name="domain" />
</br>
IP地址:
<input type="content" name="content" />
</br>
时&nbsp;效:
<input type="ttl" name="ttl" />
</br>
添加者:
<input type="auth" name="auth" />
</br>
<input type="submit" value="添加记录" />
      </form>
	  <p>&nbsp;</p>
      <form name="form2" method="post" action="/search">
        <label>
          <input type="text" name="searchDNS" id="searchDNS">
        </label>
        <label>
          <input type="submit" name="search" id="search" value="查找记录">
        </label>
      </form>
	<p>&nbsp;</p>      
    </div>
  <footer>
	 {{.LayoutContent}}
  </footer>
</body>
</html>
