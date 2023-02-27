package oshrusana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OshrusanaCounty interface {
    feud.County
    BBanjikat班吉卡特() 	feud.Barony
    BBekobod贝科博德() 	feud.Barony
    BIsfana伊斯法纳() 	feud.Barony
    BKhujand苦盏() 	feud.Barony
    BKurkat古尔卡特() 	feud.Barony
    BNau那乌() 	feud.Barony
    BOshrusana奥什鲁萨纳() 	feud.Barony
}

type 奥什鲁萨纳OshrusanaCounty struct {
	feud.BaseCounty
	班吉卡特Banjikat 	feud.Barony
	贝科博德Bekobod 	feud.Barony
	伊斯法纳Isfana 	feud.Barony
	苦盏Khujand 	feud.Barony
	古尔卡特Kurkat 	feud.Barony
	那乌Nau 	feud.Barony
	奥什鲁萨纳Oshrusana 	feud.Barony
}

func (c *奥什鲁萨纳OshrusanaCounty) BBanjikat班吉卡特() feud.Barony {
	return c.班吉卡特Banjikat
}
    
func (c *奥什鲁萨纳OshrusanaCounty) BBekobod贝科博德() feud.Barony {
	return c.贝科博德Bekobod
}
    
func (c *奥什鲁萨纳OshrusanaCounty) BIsfana伊斯法纳() feud.Barony {
	return c.伊斯法纳Isfana
}
    
func (c *奥什鲁萨纳OshrusanaCounty) BKhujand苦盏() feud.Barony {
	return c.苦盏Khujand
}
    
func (c *奥什鲁萨纳OshrusanaCounty) BKurkat古尔卡特() feud.Barony {
	return c.古尔卡特Kurkat
}
    
func (c *奥什鲁萨纳OshrusanaCounty) BNau那乌() feud.Barony {
	return c.那乌Nau
}
    
func (c *奥什鲁萨纳OshrusanaCounty) BOshrusana奥什鲁萨纳() feud.Barony {
	return c.奥什鲁萨纳Oshrusana
}
    
var COshrusana奥什鲁萨纳 OshrusanaCounty = &奥什鲁萨纳OshrusanaCounty{}

func init() {
	f := COshrusana奥什鲁萨纳.(*奥什鲁萨纳OshrusanaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1181",
		Title:     "oshrusana",
		TitleName: "奥什鲁萨纳",
		TitleCode: "c_oshrusana",
		Baronies:  map[string]feud.Barony{},
	}

	f.班吉卡特Banjikat = BBanjikat班吉卡特
	f.班吉卡特Banjikat.SetParent(f)
	
	f.贝科博德Bekobod = BBekobod贝科博德
	f.贝科博德Bekobod.SetParent(f)
	
	f.伊斯法纳Isfana = BIsfana伊斯法纳
	f.伊斯法纳Isfana.SetParent(f)
	
	f.苦盏Khujand = BKhujand苦盏
	f.苦盏Khujand.SetParent(f)
	
	f.古尔卡特Kurkat = BKurkat古尔卡特
	f.古尔卡特Kurkat.SetParent(f)
	
	f.那乌Nau = BNau那乌
	f.那乌Nau.SetParent(f)
	
	f.奥什鲁萨纳Oshrusana = BOshrusana奥什鲁萨纳
	f.奥什鲁萨纳Oshrusana.SetParent(f)
	
}
