import requests

a = "hello,world"

#定义url
url = 'https://fanyi.baidu.com/sug'
#发送一个请求
headers = {
    'user-agent':'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36'
}
data = {'kw':a}
res = requests.post(url,headers=headers,data=data)
code = res.status_code
if code == 200:
    print('请求成功')
#print(res.text)
print(res.json())
resdata = res.json()['data'][0]['v']
print(resdata)