package baidutranslatefree

const userAgent = "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Mobile Safari/537.36"
const baseUrl = "https://fanyi.baidu.com/basetrans"

var LangList = map[string]string{
	"zh":    "中文",
	"jp":    "日语",
	"jpka":  "日语假名",
	"th":    "泰语",
	"fra":   "法语",
	"en":    "英语",
	"spa":   "西班牙语",
	"kor":   "韩语",
	"tr":    "土耳其语",
	"vie":   "越南语",
	"ms":    "马来语",
	"de":    "德语",
	"ru":    "俄语",
	"ir":    "伊朗语",
	"ara":   "阿拉伯语",
	"est":   "爱沙尼亚语",
	"be":    "白俄罗斯语",
	"bul":   "保加利亚语",
	"hi":    "印地语",
	"is":    "冰岛语",
	"pl":    "波兰语",
	"fa":    "波斯语",
	"dan":   "丹麦语",
	"tl":    "菲律宾语",
	"fin":   "芬兰语",
	"nl":    "荷兰语",
	"ca":    "加泰罗尼亚语",
	"cs":    "捷克语",
	"hr":    "克罗地亚语",
	"lv":    "拉脱维亚语",
	"lt":    "立陶宛语",
	"rom":   "罗马尼亚语",
	"af":    "南非语",
	"no":    "挪威语",
	"pt_BR": "巴西语",
	"pt":    "葡萄牙语",
	"swe":   "瑞典语",
	"sr":    "塞尔维亚语",
	"eo":    "世界语",
	"sk":    "斯洛伐克语",
	"slo":   "斯洛文尼亚语",
	"sw":    "斯瓦希里语",
	"uk":    "乌克兰语",
	"iw":    "希伯来语",
	"el":    "希腊语",
	"hu":    "匈牙利语",
	"hy":    "亚美尼亚语",
	"it":    "意大利语",
	"id":    "印尼语",
	"sq":    "阿尔巴尼亚语",
	"am":    "阿姆哈拉语",
	"as":    "阿萨姆语",
	"az":    "阿塞拜疆语",
	"eu":    "巴斯克语",
	"bn":    "孟加拉语",
	"bs":    "波斯尼亚语",
	"gl":    "加利西亚语",
	"ka":    "格鲁吉亚语",
	"gu":    "古吉拉特语",
	"ha":    "豪萨语",
	"ig":    "伊博语",
	"iu":    "因纽特语",
	"ga":    "爱尔兰语",
	"zu":    "祖鲁语",
	"kn":    "卡纳达语",
	"kk":    "哈萨克语",
	"ky":    "吉尔吉斯语",
	"lb":    "卢森堡语",
	"mk":    "马其顿语",
	"mt":    "马耳他语",
	"mi":    "毛利语",
	"mr":    "马拉提语",
	"ne":    "尼泊尔语",
	"or":    "奥利亚语",
	"pa":    "旁遮普语",
	"qu":    "凯楚亚语",
	"tn":    "塞茨瓦纳语",
	"si":    "僧加罗语",
	"ta":    "泰米尔语",
	"tt":    "塔塔尔语",
	"te":    "泰卢固语",
	"ur":    "乌尔都语",
	"uz":    "乌兹别克语",
	"cy":    "威尔士语",
	"yo":    "约鲁巴语",
	"yue":   "粤语",
	"wyw":   "文言文",
	"cht":   "中文繁体",
}

const JS = `
var _gtk = "320305.131321201"
function a(r, o) {
    for (var t = 0; t < o.length - 2; t += 3) {
        var a = o.charAt(t + 2);
        a = a >= "a" ? a.charCodeAt(0) - 87 : Number(a),
        a = "+" === o.charAt(t + 1) ? r >>> a: r << a,
        r = "+" === o.charAt(t) ? r + a & 4294967295 : r ^ a
    }
    return r
}
var C = null;
function token(r) {
    var o = r.length;
    o > 30 && (r = "" + r.substr(0, 10) + r.substr(Math.floor(o / 2) - 5, 10) + r.substring(r.length, r.length - 10));
    var t = void 0,
    t = null !== C ? C: (C = _gtk || "") || "";
    for (var e = t.split("."), h = Number(e[0]) || 0, i = Number(e[1]) || 0, d = [], f = 0, g = 0; g < r.length; g++) {
        var m = r.charCodeAt(g);
        128 > m ? d[f++] = m: (2048 > m ? d[f++] = m >> 6 | 192 : (55296 === (64512 & m) && g + 1 < r.length && 56320 === (64512 & r.charCodeAt(g + 1)) ? (m = 65536 + ((1023 & m) << 10) + (1023 & r.charCodeAt(++g)), d[f++] = m >> 18 | 240, d[f++] = m >> 12 & 63 | 128) : d[f++] = m >> 12 | 224, d[f++] = m >> 6 & 63 | 128), d[f++] = 63 & m | 128)
    }
    for (var S = h,
    u = "+-a^+6",
    l = "+-3^+b+-f",
    s = 0; s < d.length; s++) S += d[s],
    S = a(S, u);

    return S = a(S, l),
    S ^= i,
    0 > S && (S = (2147483647 & S) + 2147483648),
    S %= 1e6,
    S.toString() + "." + (S ^ h)
}
`
