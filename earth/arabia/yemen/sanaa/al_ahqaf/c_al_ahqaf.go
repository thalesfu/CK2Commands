package al_ahqaf

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Al_ahqafCounty interface {
    feud.County
    BAl_abr阿卜尔堡() 	feud.Barony
    BAl_ahqaf艾哈嘎弗() 	feud.Barony
    BAl_jayat杰阿特() 	feud.Barony
    BAlan阿兰() 	feud.Barony
    BBujan不阿詹() 	feud.Barony
    BKatabiya卡塔比亚() 	feud.Barony
    BSabwa舍卜沃() 	feud.Barony
}

type 艾哈嘎弗Al_ahqafCounty struct {
	feud.BaseCounty
	阿卜尔堡Al_abr 	feud.Barony
	艾哈嘎弗Al_ahqaf 	feud.Barony
	杰阿特Al_jayat 	feud.Barony
	阿兰Alan 	feud.Barony
	不阿詹Bujan 	feud.Barony
	卡塔比亚Katabiya 	feud.Barony
	舍卜沃Sabwa 	feud.Barony
}

func (c *艾哈嘎弗Al_ahqafCounty) BAl_abr阿卜尔堡() feud.Barony {
	return c.阿卜尔堡Al_abr
}
    
func (c *艾哈嘎弗Al_ahqafCounty) BAl_ahqaf艾哈嘎弗() feud.Barony {
	return c.艾哈嘎弗Al_ahqaf
}
    
func (c *艾哈嘎弗Al_ahqafCounty) BAl_jayat杰阿特() feud.Barony {
	return c.杰阿特Al_jayat
}
    
func (c *艾哈嘎弗Al_ahqafCounty) BAlan阿兰() feud.Barony {
	return c.阿兰Alan
}
    
func (c *艾哈嘎弗Al_ahqafCounty) BBujan不阿詹() feud.Barony {
	return c.不阿詹Bujan
}
    
func (c *艾哈嘎弗Al_ahqafCounty) BKatabiya卡塔比亚() feud.Barony {
	return c.卡塔比亚Katabiya
}
    
func (c *艾哈嘎弗Al_ahqafCounty) BSabwa舍卜沃() feud.Barony {
	return c.舍卜沃Sabwa
}
    
var CAl_ahqaf艾哈嘎弗 Al_ahqafCounty = &艾哈嘎弗Al_ahqafCounty{}

func init() {
	f := CAl_ahqaf艾哈嘎弗.(*艾哈嘎弗Al_ahqafCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1534",
		Title:     "al_ahqaf",
		TitleName: "艾哈嘎弗",
		TitleCode: "c_al_ahqaf",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿卜尔堡Al_abr = BAl_abr阿卜尔堡
	f.阿卜尔堡Al_abr.SetParent(f)
	
	f.艾哈嘎弗Al_ahqaf = BAl_ahqaf艾哈嘎弗
	f.艾哈嘎弗Al_ahqaf.SetParent(f)
	
	f.杰阿特Al_jayat = BAl_jayat杰阿特
	f.杰阿特Al_jayat.SetParent(f)
	
	f.阿兰Alan = BAlan阿兰
	f.阿兰Alan.SetParent(f)
	
	f.不阿詹Bujan = BBujan不阿詹
	f.不阿詹Bujan.SetParent(f)
	
	f.卡塔比亚Katabiya = BKatabiya卡塔比亚
	f.卡塔比亚Katabiya.SetParent(f)
	
	f.舍卜沃Sabwa = BSabwa舍卜沃
	f.舍卜沃Sabwa.SetParent(f)
	
}
