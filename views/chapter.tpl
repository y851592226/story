{{define "head"}}
<meta charset="utf8" />
<title>{{ .Title }}</title>
<meta name="MobileOptimized" content="240"/>
<meta name="applicable-device" content="mobile"/>
<meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0" />
<meta http-equiv="Content-Type" content="application/vnd.wap.xhtml+xml;charset=gbk"/>
<link rel="shortcut icon" href="/favicon.ico" />
<link href="/file/css/style.css" rel="stylesheet">
<link href="/file/css/layer.css" rel="stylesheet">
<link rel="stylesheet" href="/file/css/jquery-ui.css">
<script src="/file/js/jquery-1.9.1.js"></script>
<script src="/file/js/jquery-ui.js"></script>
<script src="/file/js/style.js" type="text/javascript"></script>
<script src="/file/js/layer.js?" type="text/javascript"></script>
{{end}}
{{define "line"}}
{{.}}<br>
{{end}}
{{define "mainchapterbody"}}
<body id="BookBody">
<div id="main">
  <section id="container">
  <div class="bookset clrfix"><script>bookset();</script></div>
    <article>
      <h1 id="BookTitle">{{ .Chapter_name }}</h1>
      <p class="link bx"><a href={{ .Href_before }}>上一章</a><a href={{ .Href_mulu }}>回目录</a><a href={{ .Href_after }}>下一章</a></p>
      <div id="content">
        {{ range .Content }}
        {{template "line" .}}
        {{end}}
      </div>
      <p class="link"><a href={{ .Href_before }}>上一章</a><a href={{ .Href_mulu }}>回目录</a><a href={{ .Href_after }}>下一章</a></p>
    </article>
  </section>
</div>
<script>getset()</script>
<script type="text/javascript">var preview_page={{ .Href_before }};var next_page={{ .Href_after }};var index_page={{ .Href_mulu }};</script>
</body>
{{end}}
<!DOCTYPE html>
<html>
<head>
{{template "head" .}}
</head>
</html>