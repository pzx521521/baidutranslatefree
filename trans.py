# !/usr/bin/python3
# -*- coding: utf-8 -*-

"""
info:
author:CriseLYJ
github:https://github.com/CriseLYJ/
"""

"""
请求url分析 :https://fanyi.baidu.com/basetrans
请求方式分析 :POST
请求参数分析 : {
query: hello
from: en
to: zh
token: 6f5c83b84d69ad3633abdf18abcb030d
sign: 54706.276099}
请求头分析
"""

# 代码实现流程
# 1. 实现面对对象构建爬虫对象
# 2. 爬虫流程四步骤
# 2.1 获取URl
# 2.2 发送请求获取响应
# 2.3 从响应中提取数据
# 2.4 保存数据

import requests
import js2py

context = js2py.EvalJs()

# 翻译模式
# 0:英译中 1:中译英
translating_mode = 0

class BaiDuTranslater(object):
    """
    百度翻译爬虫
    """

    def __init__(self, query):
        # 初始化
        self.url = "https://fanyi.baidu.com/basetrans"
        self.query = query
        self.headers = {
            "User-Agent": "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1",
            "Referer": "https://fanyi.baidu.com/",
            "Cookie": "BAIDUID=B61F7D6CE4B16113D11D2324CEDF91EF:SL=0:NR=10:FG=1;"
        }

    def make_sign(self):
        # js逆向获取sign的值
        with open("translate.js", "r", encoding="utf-8") as f:
            context.execute(f.read())

        # 调用js中的函数生成sign
        sign = context.a(self.query)
        # 将sign加入到data中
        return sign

    def make_data(self, sign):
        # 判断翻译模式,选取对应的 from 和 to 值.
        if translating_mode == 0:
            from_str = "en"
            to_str = "zh"
        else:
            from_str = "zh"
            to_str = "en"
        data = {
            "query": self.query,
            "from": from_str,
            "to": to_str,
            "token": "f1a8c5324e2a5667fe9501845cb83b69",
            "sign": sign
        }
        return data

    def get_content(self, data):
        # 发送请求获取响应
        response = requests.post(
            url=self.url,
            headers=self.headers,
            data=data
        )
        return response.json()["trans"][0]["dst"]

    def run(self):
        """运行程序"""
        # 获取sign的值
        sign = self.make_sign()
        # 构建参数
        data = self.make_data(sign)
        # 获取翻译内容
        content = self.get_content(data)
        print(content)


if __name__ == '__main__':
    translater = BaiDuTranslater("Hello,World!")
    translater.run()