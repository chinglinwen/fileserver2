# Fileserver2

A http download and upload server

## Features

* Default to append
* Delete support
* Truncate support
* Form value upload
* File upload ( no need extra static index page anymore )
* Any level path specify 
* Web Any path upload support

## Install

```
go get github.com/chinglinwen/fileserver2
```

## Examples

### File upload

```
[~ t1 ] $ cat example.log 
example body
[~ t1 ] $ curl localhost:8000/uploadapi -F file=@example.log
Files uploaded successfully : example.log 13 bytes 
```

Results

```
[~ t ] $ cat example.log 
example body
```

### Form upload

```
[~ t1 ] $ curl localhost:8000/uploadapi -F file=test -F data="test body"
Files uploaded successfully : test 9 bytes 
```

### Path specify

```
[~ t1 ] $ curl localhost:8000/uploadapi -F file="a/b/c" -F data="hello"
Files uploaded successfully : a/b/c 5 bytes 

[~ t1 ] $ curl localhost:8000/uploadapi -F file="a/b/c" -F file=@a.txt 
Files uploaded successfully : a/b/c 3 bytes 

[~ t1 ] $ curl localhost:8000/uploadapi -F file="../a/b/c" -F file=@a.txt
file path should not contain the two dot
```

or

```
[~ t1 ] $ curl localhost:8000/a/b/uploadapi -F file="c" -F data="hello"
Files uploaded successfully : a/b/c 5 bytes 

[~ t1 ] $ curl localhost:8000/a/b/uploadapi -F file="c" -F file=@a.txt 
Files uploaded successfully : a/b/c 3 bytes 

[~ t1 ] $ curl localhost:8000/../a/b/uploadapi -F file="c" -F file=@a.txt   # won't work
```

### Delete

```
[~ t1 ] $ curl localhost:8000/uploadapi -F file=example.log -F delete=yes 
file: example.log deleted

[~ t1 ] $ curl localhost:8000/uploadapi -F file="a/b/c" -F delete=yes  
file: a/b/c deleted

[~ t1 ] $ curl localhost:8000/uploadapi -F file="a" -F delete=yes      
file: a deleted
```

### Truncate

if file not exist, it will create the file

if file exist, by default it will append to the file

use truncate to overwrite the file

```
curl localhost:8000/uploadapi -F file=@example.log -F truncate=yes 
```

or

```
curl localhost:8000/uploadapi -F file=test -F data="test body" -F truncate=yes 
```

### Usage

```
$ fileserver2 -h
Usage of fileserver2:
  -author
        Show author.
  -path string
        File server path. (default ".")
  -port string
        Port number. (default "8000")
  -v    Show version.
```

Notes: **path** specify where the file will be stored


## Web examples

### Choose files

![choose_files](doc/fileserver2-web1.png)

### Specify any path in url to upload

![webupload1](doc/webupload1.png)

![webupload2](doc/webupload2.png)

### Info

![choose_files](doc/fileserver2-web2.png)

### Results

![choose_files](doc/fileserver2-web3.png)

## Download example

### Web

![download](doc/fileserver2-web4.png)

### CLI

```
[~ t1 ] $ wget "localhost:8000/Track 1.wav"
--2016-06-23 11:43:18--  http://localhost:8000/Track%201.wav
Resolving localhost... ::1, 127.0.0.1
Connecting to localhost|::1|:8000... connected.
HTTP request sent, awaiting response... 200 OK
Length: 26869360 (26M) [audio/x-wav]
Saving to: `Track 1.wav'

100%[=============================================================================>] 26,869,360  --.-K/s   in 0.1s    

2016-06-23 11:43:18 (240 MB/s) - `Track 1.wav' saved [26869360/26869360]

[~ t1 ] $
```

## Service register and deregister

register

```
./reg.sh
```

deregister

```
./dereg.sh
```

![register](doc/reg.png)

end
