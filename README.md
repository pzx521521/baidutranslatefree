# 百度翻译的免费接口
百度翻译的免费接口
## 如何使用

## Install:
```
go get github.com/pzx521521/baidutranslatefree
```


Example usage:
```
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
## Token baiduID 获取

## CURL
原网页curl
```curl
curl -v 'https://fanyi.baidu.com/basetrans' \
  -H 'Cookie: BIDUPSID=6C787B3A6E1FFD2F18A23A5B5102C6CE; BDUSS=loa3hzU0JRT3RTSVlFYk0zNFNtZzIwbTNidWI2c0Raa2xPelB3NG1TZTdsU05uSVFBQUFBJCQAAAAAAAAAAAEAAADrTaMHNDYwMDY3OTYwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAALsI~Ga7CPxmSz; BDUSS_BFESS=loa3hzU0JRT3RTSVlFYk0zNFNtZzIwbTNidWI2c0Raa2xPelB3NG1TZTdsU05uSVFBQUFBJCQAAAAAAAAAAAEAAADrTaMHNDYwMDY3OTYwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAALsI~Ga7CPxmSz; BAIDUID=B61F7D6CE4B16113D11D2324CEDF91EF:SL=0:NR=10:FG=1; H_PS_PSSID=60450_60815_60620_60886_60875; BAIDUID_BFESS=B61F7D6CE4B16113D11D2324CEDF91EF:SL=0:NR=10:FG=1; delPer=0; BDRCVFR[feWj1Vr5u3D]=mk3SLVN4HKm; BA_HECTOR=25ak80a0ak0d2h0485242l8hkvn57h1jhcma61v; ZFY=QP9uUFZMCwFF3L:BXPFu5zDpYVWlQ15NOJnK29YKERpU:C; rsv_i=5106rSWG0Js3Emz4RuzSqb/EJYbeRuxQYP6WEHU4wULex+al+HEXfax397XgfGX3VosTb9Ykex7UMF6EOCUr6Kz8OehxUYE; BDORZ=AE84CDB3A529C0F8A2B9DCDD1D18B695; SE_LAUNCH=5%3A28825299_0%3A28825299_16%3A28825299_15%3A28825299; BDPASSGATE=IlPT2AEptyoA_yiU4SOA3lIN8eDEUsCD34OtViBi3ECGh67BmhH84NxhSl8LSSurGULMdI3JmcldjijsQmFuirMenBIUiixSb6Ta7tiVx_TwSa62-rss_2zgGFkUwuDPbRhL-3MEF3VEVFpbevHGj3wigBSnsQlsecbNs6LWmMr81TyG3mHQx6iIK8RkLGPYKr_C94rSjE6KLl_WZOmtT6jebCJHO7As70aOb2YiAvfdyEoXCurSRvEA0nXQ37J2_4K-1PW79cC8AE9SzZR8ST1tdEiV5tD6MSgsAM4UlK6DSKzePcCQUiqG_swka0PbSwdWKQ3zmtkGDUYX6XYlJHVj_NnFFTHrDjUNYMGGihjZCXwVqlOEIw4xqptfE0LpW4-0LBpDYit-hS7fy3CmpSr2HwPgi0FvMtoMInOmzJl4bn63Gm4K6Hqzc7Jv-oahWVaG1SbdArnlLWNKuuKyJaXaVKr82VEhqS_XgECavkjiZh4aNMR7kzgCwnV7TWiB9e7WVpXjUwTtyNlfwrXqvszDu9SUqzaTw6uCLdyyNAAKVYI0qBJ4CmjJzGy_eXx_O4G6v7z9IUjZj5RMwtyIjyZnlPA01NJQOs4rBc1LwcSIkWwV0yYto_Sz1wZD2-rL9AZT_dWkY98LAobVVmVRgrtA2V4vIJTzvt4wE52TBz4p1HXh0KhpxUF8h71WOQ4AZ5buavsEUVuCbQYjL2qip7diFOEjEi25v8zY3J8NCRn7y4REqUEut0or0oat104XwIk7oagqUyCuyq0ZIai7hZgZy8yJgQZBV0q3X7vk8FKyIBEq-nnzEFzbxK;' \
  --data-raw 'query=hello%2Cword&from=en&to=zh&token=f1a8c5324e2a5667fe9501845cb83b69&sign=645327.867326'
```
经测试 cookie仅需保留 BAIDUID 和 UA
```curl

```