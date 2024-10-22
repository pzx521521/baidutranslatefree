# 百度翻译的免费接口
百度翻译的免费接口,起因是想在 Go 语言中实现翻译功能，但是没有找到合适的翻译 API，于是就自己写了一个。
## 如何使用

### Install:
```
go get github.com/pzx521521/baidutranslatefree
```


Example usage:
```golang
package main

import (
    "fmt"
    gt "github.com/pzx521521/baidutranslatefree"
)

func main(){
	input := "Hello,World!"
	translater, _ := NewBaiduTranslater()
	// you can set it or not, it will be set by default("en", "zh")
	translater.SetFromTo("en", "zh")
	text, _ := translater.TransPort(input)
	fmt.Printf("%v\n", text)
    // Output: "你好，世界！"
}
```
All language codes can be found here:
[const.go](https://github.com/pzx521521/googletranslatefree/blob/main/const.go)


## 破解过程分享
### 前因说明
+ 目前大部分google接口都是用的`https://translate.google.com/translate_a/single`
    但他总是弹验证(google的reCAPTCHA),
    因为国内的原因,ip经常改变,所以经常弹验证.导致用不了
+ 发现  `https://translate.google.com/m` 是可以用的
    python版本[deep-translator](https://github.com/nidhaloff/deep-translator)

+ 看了一下,并没有golang版的(因为想二进制在arm上跑),所以就写了一个
    [googletranslatefree](https://github.com/pzx521521/googletranslatefree)

+ 写完了,想到很多人没法正常访问google, 忽然想看下buidu的
    电脑版的有点麻烦，使用的是**手机版**的接口分析
+ 为什么要写golang版本
  想在arm芯片/手机 上以二进制包的方式运行

### CURL
原网页curl
```curl
curl -v 'https://fanyi.baidu.com/basetrans' \
  -H 'Cookie: BIDUPSID=6C787B3A6E1FFD2F18A23A5B5102C6CE; BDUSS=loa3hzU0JRT3RTSVlFYk0zNFNtZzIwbTNidWI2c0Raa2xPelB3NG1TZTdsU05uSVFBQUFBJCQAAAAAAAAAAAEAAADrTaMHNDYwMDY3OTYwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAALsI~Ga7CPxmSz; BDUSS_BFESS=loa3hzU0JRT3RTSVlFYk0zNFNtZzIwbTNidWI2c0Raa2xPelB3NG1TZTdsU05uSVFBQUFBJCQAAAAAAAAAAAEAAADrTaMHNDYwMDY3OTYwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAALsI~Ga7CPxmSz; BAIDUID=B61F7D6CE4B16113D11D2324CEDF91EF:SL=0:NR=10:FG=1; H_PS_PSSID=60450_60815_60620_60886_60875; BAIDUID_BFESS=B61F7D6CE4B16113D11D2324CEDF91EF:SL=0:NR=10:FG=1; delPer=0; BDRCVFR[feWj1Vr5u3D]=mk3SLVN4HKm; BA_HECTOR=25ak80a0ak0d2h0485242l8hkvn57h1jhcma61v; ZFY=QP9uUFZMCwFF3L:BXPFu5zDpYVWlQ15NOJnK29YKERpU:C; rsv_i=5106rSWG0Js3Emz4RuzSqb/EJYbeRuxQYP6WEHU4wULex+al+HEXfax397XgfGX3VosTb9Ykex7UMF6EOCUr6Kz8OehxUYE; BDORZ=AE84CDB3A529C0F8A2B9DCDD1D18B695; SE_LAUNCH=5%3A28825299_0%3A28825299_16%3A28825299_15%3A28825299; BDPASSGATE=IlPT2AEptyoA_yiU4SOA3lIN8eDEUsCD34OtViBi3ECGh67BmhH84NxhSl8LSSurGULMdI3JmcldjijsQmFuirMenBIUiixSb6Ta7tiVx_TwSa62-rss_2zgGFkUwuDPbRhL-3MEF3VEVFpbevHGj3wigBSnsQlsecbNs6LWmMr81TyG3mHQx6iIK8RkLGPYKr_C94rSjE6KLl_WZOmtT6jebCJHO7As70aOb2YiAvfdyEoXCurSRvEA0nXQ37J2_4K-1PW79cC8AE9SzZR8ST1tdEiV5tD6MSgsAM4UlK6DSKzePcCQUiqG_swka0PbSwdWKQ3zmtkGDUYX6XYlJHVj_NnFFTHrDjUNYMGGihjZCXwVqlOEIw4xqptfE0LpW4-0LBpDYit-hS7fy3CmpSr2HwPgi0FvMtoMInOmzJl4bn63Gm4K6Hqzc7Jv-oahWVaG1SbdArnlLWNKuuKyJaXaVKr82VEhqS_XgECavkjiZh4aNMR7kzgCwnV7TWiB9e7WVpXjUwTtyNlfwrXqvszDu9SUqzaTw6uCLdyyNAAKVYI0qBJ4CmjJzGy_eXx_O4G6v7z9IUjZj5RMwtyIjyZnlPA01NJQOs4rBc1LwcSIkWwV0yYto_Sz1wZD2-rL9AZT_dWkY98LAobVVmVRgrtA2V4vIJTzvt4wE52TBz4p1HXh0KhpxUF8h71WOQ4AZ5buavsEUVuCbQYjL2qip7diFOEjEi25v8zY3J8NCRn7y4REqUEut0or0oat104XwIk7oagqUyCuyq0ZIai7hZgZy8yJgQZBV0q3X7vk8FKyIBEq-nnzEFzbxK;' \
  --data-raw 'query=hello%2Cword&from=en&to=zh&token=f1a8c5324e2a5667fe9501845cb83b69&sign=645327.867326'
```

抓包结果,可以看到需要填入cookie, token, sign
### 总体流程
+ GET[主页](fanyi.baidu.com),从Set-Cookie获取BAIDUID
+ 通过GET[主页](fanyi.baidu.com) 携带BAIDUID获取token
+ 每次插叙的时候,根据查询的内容计算sign
+ 访问 `https://fanyi.baidu.com/basetrans` 携带BAIDUID,token,sign 返回结果

### BAIDUID 
+ BAIDUID 直接GET主页,从Set-Cookie获取BAIDUID
+ GET主页的时候如果不携带cookie,会返回Set-Cookie,如果携带则不会返回Set-Cookie
+ GET的时候注意一定要加**手机的UA**,手机和电脑端的不一样,估计是在路由层(如:nginx)做了处理
经测试(逐条删除cookies数据看返回结果) cookie仅需保留 BAIDUID 和 UA

###  token 获取
  + 注意要请求两遍,第一遍是获取BAIDUID，第二遍是通过BAIDUID获取token
  + token位于html中, 具体代码如下:
  `fanyi.baidu.com/ line 21`
  ```
  <script>
  page.common = {
      token: '3079487ebfc9fc012ef8dbcc18894a7d',
      langList: {
  ```

+ 如果不携带BAIDUID,会返回
  `fanyi.baidu.com/ line 21`
  ```js
  <script>
  page.common = {
      token: '',
      langList: {
  ```
  然后从js代码中重新加载页面
  `fanyi.baidu.com/ line 927`
  ```js
  if (!page.common.token) {
      location.reload();
  }
  ```
  
### Sign 计算
+ 比如翻译`hello`,Sign就是对`hello`的计算
+ 参考代码如下:
  `https://github.com/Kr1s77/awesome-python-login-model`
  文件中的`translate.js`就是从他那里拿的,
  `trans.py` 修改最新的BAIDUID和token之后还是能用的
+ golang中使用的是[OTTO](github.com/robertkrimen/otto)解析json,本来想让chatGPT直接把js转成golang的,
  但是他转的有问题(calsign_chatgpt_err.go),懒得一点一点跟.放弃,直接使用原生js
+ 使用otto解析的时候会有问题,因为js中`translate.js`中会有一些逻辑永远用不到,在这些逻辑中会有未定义的函数,
  otto解析的时候就会报错(python-js2py/浏览器解析都没有问题),
  本想pr给otto的,发现他测试文档中大篇幅测试了函数未定义的错误,pr之后也不会采纳,就放弃了
  golang中通过修改js代码来解决这个问题

### todo
目前没有写python版本的想法,因为[作者](https://github.com/Kr1s77/awesome-python-login-model)并不是一个单独的项目,所以没办法pr