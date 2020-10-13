# webp server go 是go 版本的图片压缩显示服务

## 打包 go build

## 配置config.json

```json
{
	"HOST": "127.0.0.1",
	"PORT": "3333",
	"QUALITY": "80",
	"IMG_PATH": "/path/to/pics",
	"EXHAUST_PATH": "/path/to/exhaust",
	"ALLOWED_TYPES": ["jpg","png","jpeg"]
}
```
host:地址
port:端口
QUALITY:质量 1-99
IMG_PATH：图片存放文件夹
ALLOWED_TYPES：支持转化的图片类型

## 执行(windows) webp-server.exe (后面添加可选参数，比如-config xxx.json) 
 
## 访问图片(IMG_PATH 文件夹下的图片。)：比如 http://localhost:3333/no.gif


<p align="center">
	<img src="./pics/webp_server.png"/>
</p>
<img src="https://api.travis-ci.org/webp-sh/webp_server_go.svg?branch=master"/>

[Documentation](https://docs.webp.sh/) | [Website](https://webp.sh/)

This is a Server based on Golang, which allows you to serve WebP images on the fly. 
It will convert `jpg,jpeg,png` files by default, this can be customized by editing the `config.json`.. 
* currently supported  image format: JPEG, PNG, BMP, GIF(static image for now)


> e.g When you visit `https://your.website/pics/tsuki.jpg`，it will serve as `image/webp` format without changing the URL.
>
> For Safari and Opera users, the original image will be used.


## Simple Usage Steps

### 1. Download or build the binary
Download the `webp-server` from [release](https://github.com/webp-sh/webp_server_go/releases) page.

### 2. Dump config file

```
./webp-server -dump-config > config.json
```

The default `config.json` may look like this.
```json
{
	"HOST": "127.0.0.1",
	"PORT": "3333",
	"QUALITY": "80",
	"IMG_PATH": "/path/to/pics",
	"EXHAUST_PATH": "/path/to/exhaust",
	"ALLOWED_TYPES": ["jpg","png","jpeg"]
}
```

#### Config Example

In the following example, the image path and website URL.

| Image Path                            | Website Path                         |
| ------------------------------------- | ------------------------------------ |
| `/var/www/img.webp.sh/path/tsuki.jpg` | `https://img.webp.sh/path/tsuki.jpg` |

The `config.json` should be like:

| IMG_PATH               |
| ---------------------- |
| `/var/www/img.webp.sh` |


`EXHAUST_PATH` is cache folder for output `webp` images, with `EXHAUST_PATH` set to `/var/cache/webp` 
in the example above, your `webp` image will be saved at `/var/cache/webp/pics/tsuki.jpg.1582558990.webp`.

### 3. Run

```
./webp-server --config=/path/to/config.json
```

### 4. Nginx proxy_pass
Let Nginx to `proxy_pass http://localhost:3333/;`, and your webp-server is on-the-fly.

## Advanced Usage

For supervisor, Docker sections, please read our documentation at [https://docs.webp.sh/](https://docs.webp.sh/)

## License

WebP Server is under the GPLv3. See the [LICENSE](./LICENSE) file for details.

