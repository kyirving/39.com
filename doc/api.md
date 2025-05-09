## API接口文档

### 1. 接口域名
``` 127.0.0.1:8080 ```

### 2. 请求方式
``` POST ```

### 3. 请求头
```Content-Type: application/json```

### 4. 编码方式
``` utf-8 ```



### 3. 公共参数
| 名称 | 类型 | 必须 | 描述 |
| :----: | :----: | :----: | :----: |
| access_token | string | 是 | 登陆授权后获得 |
| udid | string | 是 | 用户设备号 |
| timestamp | string | 是 | 时间戳 |
| version | string | 是 | 版本号 |
| sign_type | string | 是 | 签名类型 |
| request_id | string | 是 | 请求ID，由调用方自行生成，需保证全局唯一性 |
| sign | string | 是 | 签名 |

### 4. 签名算法
* 将所有请求字段(包含公共参数、私有参数，但sign字段除外)按照字段名的字母顺序排列，字段名和其字段值之间用=号相连，字段间用&符相连，得到待签字符串
   > akey=aval&bkey=bval& …… &zkey=zval
* 将待签字符进行md5运算，最终得到一个32位的字符串散列值（转换为小写），此值即为sign参数的值