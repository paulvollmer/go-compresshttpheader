package compresshttpheader

const dictionaryHeader = `Access-Control-Allow-Credentials` +
	`Access-Control-Allow-Headers` +
	`Access-Control-Allow-Methods` +
	`Access-Control-Allow-Origin` +
	`Access-Control-Expose-Headers` +
	`Access-Control-Max-Age` +
	`Accept-Charset` +
	`Accept-Encoding` +
	`Accept-Ranges` +
	`Age` +
	`Allow` +
	`Alternate-Protocol` +
	`Authorization` +
	`Cache-Control` +
	`Cache-Control` +
	`Client-Date` +
	`Client-Peer` +
	`Client-Response-Num` +
	`Connection` +
	`Content-Disposition` +
	`Content-Encoding` +
	`Content-Language` +
	`Content-Length` +
	`Content-Location` +
	`Content-MD5` +
	`Content-Range` +
	`Content-Security-Policy` +
	`Content-Type` +
	`Date` +
	`ETag` +
	`Expires` +
	`HTTP` +
	`Keep-Alive` +
	`Last-Modified` +
	`Link` +
	`Location` +
	`P3P` +
	`Pragma` +
	`Proxy-Authenticate` +
	`Proxy-Connection` +
	`Refresh` +
	`Retry-After` +
	`Server` +
	`Set-Cookie` +
	`Status` +
	`Strict-Transport-Security` +
	`Timing-Allow-Origin` +
	`Trailer` +
	`Transfer-Encoding` +
	`Upgrade` +
	`Vary` +
	`Via` +
	`Warning` +
	`WWW-Authenticate` +
	`X-Aspnet-Version` +
	`X-Content-Security-Policy` +
	`X-Content-Type-Options` +
	`X-Frame-Options` +
	`X-Permitted-Cross-Domain-Policies` +
	`X-Pingback` +
	`X-Powered-By` +
	`X-Robots-Tag` +
	`X-UA-Compatible` +
	`X-WebKit-CSP` +
	`X-XSS-Protection`

const dictionaryValues = `true` +
	`Apache` +
	`bytes` +
	`Basic` +
	`chunked` +
	`GET` +
	`HEAD` +
	`PUT` +
	`POST` +
	`PATCH` +
	`DELETE` +
	`OPTIONS` +
	`charset=UTF-8` +
	`no-cache` +
	`deny` +
	`nosniff` +
	`noindex` +
	`nofollow` +
	`must-revalidate` +
	`private` +
	`attachment` +
	`filename=”` +
	`gzip` +
	`en` +
	`de` +
	`index.htm` +
	`index.html` +
	`Unauthorized` +
	`timeout=` +
	`max=` +
	`max-age=` +
	`Max-Forwards` +
	`mode=block` +
	`post-check=` +
	`pre-check=` +
	`preload` +
	`secure` +
	`HttpOnly`

const dictionaryContentTypes = `application/EDI-X12` +
	`application/EDIFACT` +
	`application/javascript` +
	`application/octet-stream` +
	`application/ogg` +
	`application/pdf` +
	`application/xhtml+xml` +
	`application/x-shockwave-flash` +
	`application/json` +
	`application/ld+json` +
	`application/xml` +
	`application/zip` +
	`application/x-www-form-urlencoded` +
	`audio/aac` +
	`audio/mp4` +
	`audio/mpeg` +
	`audio/x-ms-wma` +
	`audio/vnd.rn-realaudio` +
	`audio/x-wav` +
	`font/collection` +
	`font/otf` +
	`font/sfnt` +
	`font/ttf` +
	`font/woff` +
	`font/woff2` +
	`image/bmp` +
	`image/gif` +
	`image/jpeg` +
	`image/png` +
	`image/tiff` +
	`image/wmf` +
	`image/vnd.microsoft.icon` +
	`image/x-icon` +
	`image/vnd.djvu` +
	`image/svg+xml` +
	`multipart/mixed` +
	`multipart/alternative` +
	`multipart/related` +
	`multipart/form-data` +
	`text/css` +
	`text/csv` +
	`text/ecmascript` +
	`text/html` +
	`text/javascript` +
	`text/markdown` +
	`text/plain` +
	`text/xml` +
	`video/mpeg` +
	`video/mp4` +
	`video/quicktime` +
	`video/x-ms-wmv` +
	`video/x-msvideo` +
	`video/x-flv` +
	`video/webm` +
	`application/vnd.oasis.opendocument.text` +
	`application/vnd.oasis.opendocument.spreadsheet` +
	`application/vnd.oasis.opendocument.presentation` +
	`application/vnd.oasis.opendocument.graphics` +
	`application/vnd.ms-excel` +
	`application/vnd.openxmlformats-officedocument.spreadsheetml.sheet` +
	`application/vnd.ms-powerpoint` +
	`application/vnd.openxmlformats-officedocument.presentationml.presentation` +
	`application/msword` +
	`application/vnd.openxmlformats-officedocument.wordprocessingml.document` +
	`application/vnd.mozilla.xul+xml`

var dictionaryDate = `Mon,` +
	`Tue,` +
	`Wed,` +
	`Thu,` +
	`Fri,` +
	`Sat,` +
	`Sun,` +
	` Jan ` +
	` Feb ` +
	` Mar ` +
	` Apr ` +
	` May ` +
	` Jun ` +
	` Jul ` +
	` Aug ` +
	` Sep ` +
	` Oct ` +
	` Nov ` +
	` Dec ` +
	` UT` +
	` GMT` +
	` EDT` +
	` EST` +
	` CDT` +
	` CST` +
	` MDT` +
	` MST` +
	` PDT` +
	` PST`

var dictionary = []byte(dictionaryHeader + dictionaryValues + dictionaryContentTypes + dictionaryDate)
