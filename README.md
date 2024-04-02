## 描述

一个简介的多设备扫码登录器（纯服务端）。

服务端生成二维码，其他设备扫描该二维码然后服务端验证。

## 开发

复制`env.example`到`.env`然后修改`QR_CALLBACK_SERVICE_URL`的值

```env
...

QR_CALLBACK_SERVICE_URL = http://{服务端局域网IP}:8090

...
```

请确保你的多设备位于同一局域网内，并且上述的局域网IP是服务端所在的IP。

此时你就可以在你的电脑上访问

```http
GET /generateQR
```

就会生成一个二维码，接着使用其他设备（如手机，平板）访问，这样在服务端就会生成一条访问记录啦。