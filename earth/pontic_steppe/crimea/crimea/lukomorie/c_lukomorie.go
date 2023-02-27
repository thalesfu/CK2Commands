package lukomorie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LukomorieCounty interface {
    feud.County
    BChernihivka切尔尼戈夫卡() 	feud.Barony
    BHuliaipole古利艾波列() 	feud.Barony
    BKalchik卡利奇克() 	feud.Barony
    BKalmius卡利米乌斯() 	feud.Barony
    BKushunum库舒古姆() 	feud.Barony
    BOnkhiv呕基夫() 	feud.Barony
    BPolohy波洛吉() 	feud.Barony
    BZaporzhye扎波罗热() 	feud.Barony
}

type 卢科莫里耶LukomorieCounty struct {
	feud.BaseCounty
	切尔尼戈夫卡Chernihivka 	feud.Barony
	古利艾波列Huliaipole 	feud.Barony
	卡利奇克Kalchik 	feud.Barony
	卡利米乌斯Kalmius 	feud.Barony
	库舒古姆Kushunum 	feud.Barony
	呕基夫Onkhiv 	feud.Barony
	波洛吉Polohy 	feud.Barony
	扎波罗热Zaporzhye 	feud.Barony
}

func (c *卢科莫里耶LukomorieCounty) BChernihivka切尔尼戈夫卡() feud.Barony {
	return c.切尔尼戈夫卡Chernihivka
}
    
func (c *卢科莫里耶LukomorieCounty) BHuliaipole古利艾波列() feud.Barony {
	return c.古利艾波列Huliaipole
}
    
func (c *卢科莫里耶LukomorieCounty) BKalchik卡利奇克() feud.Barony {
	return c.卡利奇克Kalchik
}
    
func (c *卢科莫里耶LukomorieCounty) BKalmius卡利米乌斯() feud.Barony {
	return c.卡利米乌斯Kalmius
}
    
func (c *卢科莫里耶LukomorieCounty) BKushunum库舒古姆() feud.Barony {
	return c.库舒古姆Kushunum
}
    
func (c *卢科莫里耶LukomorieCounty) BOnkhiv呕基夫() feud.Barony {
	return c.呕基夫Onkhiv
}
    
func (c *卢科莫里耶LukomorieCounty) BPolohy波洛吉() feud.Barony {
	return c.波洛吉Polohy
}
    
func (c *卢科莫里耶LukomorieCounty) BZaporzhye扎波罗热() feud.Barony {
	return c.扎波罗热Zaporzhye
}
    
var CLukomorie卢科莫里耶 LukomorieCounty = &卢科莫里耶LukomorieCounty{}

func init() {
	f := CLukomorie卢科莫里耶.(*卢科莫里耶LukomorieCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "557",
		Title:     "lukomorie",
		TitleName: "卢科莫里耶",
		TitleCode: "c_lukomorie",
		Baronies:  map[string]feud.Barony{},
	}

	f.切尔尼戈夫卡Chernihivka = BChernihivka切尔尼戈夫卡
	f.切尔尼戈夫卡Chernihivka.SetParent(f)
	
	f.古利艾波列Huliaipole = BHuliaipole古利艾波列
	f.古利艾波列Huliaipole.SetParent(f)
	
	f.卡利奇克Kalchik = BKalchik卡利奇克
	f.卡利奇克Kalchik.SetParent(f)
	
	f.卡利米乌斯Kalmius = BKalmius卡利米乌斯
	f.卡利米乌斯Kalmius.SetParent(f)
	
	f.库舒古姆Kushunum = BKushunum库舒古姆
	f.库舒古姆Kushunum.SetParent(f)
	
	f.呕基夫Onkhiv = BOnkhiv呕基夫
	f.呕基夫Onkhiv.SetParent(f)
	
	f.波洛吉Polohy = BPolohy波洛吉
	f.波洛吉Polohy.SetParent(f)
	
	f.扎波罗热Zaporzhye = BZaporzhye扎波罗热
	f.扎波罗热Zaporzhye.SetParent(f)
	
}
