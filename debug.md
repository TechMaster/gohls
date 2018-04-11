# curl http://localhost:8080
```<a href="/ui/">Found</a>.```

# curl http://localhost:8080/ui
```<a href="/ui/">Moved Permanently</a>```

# curl http://localhost:8080/ui/index.html
```
<!doctype html>
<html lang="en">
<head>
	<meta charset="utf-8">
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	<meta name="viewport" content="width=device-width, initial-scale=1">

	<title>Videos</title>

	<link rel="stylesheet" href="/ui/assets/css/vendor.css"></script>
	<link rel="stylesheet" href="/ui/assets/css/app.css">
</head>
<body>
	<div id="app">

	</div>
	<script src="/ui/assets/js/vendor.js"></script>
	<script src="/ui/assets/js/app.js"></script>
</body>
</html>
```
# curl http://localhost:8080/list/
```json
{"error":null,"folders":[{"name":"bunny","path":"bunny"}],"videos":[{"name":"El Condor Pasa.mp4","path":"El Condor Pasa.mp4","info":{"duration":194.838633,"lastModified":"2018-04-08T22:49:42+07:00"}},{"name":"Jason Statham.mp4","path":"Jason Statham.mp4","info":{"duration":959.587844,"lastModified":"2018-04-05T00:17:14+07:00"}},{"name":"NeverEnough.mp4","path":"NeverEnough.mp4","info":{"duration":207.307744,"lastModified":"2018-04-05T00:11:49+07:00"}},{"name":"San Andreas.mp4","path":"San Andreas.mp4","info":{"duration":294.846978,"lastModified":"2018-04-05T00:13:58+07:00"}},{"name":"The Last Mohican.mp4","path":"The Last Mohican.mp4","info":{"duration":225.094233,"lastModified":"2018-04-08T22:48:21+07:00"}},{"name":"bunny.avi","path":"bunny.avi","info":{"duration":596.458333,"lastModified":"2018-04-05T19:20:13+07:00"}},{"name":"unpocolo.mp4","path":"unpocolo.mp4","info":{"duration":112.500678,"lastModified":"2018-04-07T18:38:47+07:00"}}]}
```

# curl http://localhost:8080/playlist/720/unpocolo.mp4
curl http://localhost:8080/playlist/720/San%20Andreas.mp4
Trả về danh sách play list mp3u8.
File ts dễ dàng bị ghép

