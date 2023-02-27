package tobruk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TobrukCounty interface {
    feud.County
    BAcroma阿雷马() 	feud.Barony
    BAlqardabah艾卡拉达拜() 	feud.Barony
    BBardia拜尔迪耶() 	feud.Barony
    BKambut坎布特() 	feud.Barony
    BRukbah鲁克班() 	feud.Barony
    BSollum塞卢姆() 	feud.Barony
    BTobruk图卜鲁格() 	feud.Barony
    BZanzur赞祖尔() 	feud.Barony
}

type 图卜鲁格TobrukCounty struct {
	feud.BaseCounty
	阿雷马Acroma 	feud.Barony
	艾卡拉达拜Alqardabah 	feud.Barony
	拜尔迪耶Bardia 	feud.Barony
	坎布特Kambut 	feud.Barony
	鲁克班Rukbah 	feud.Barony
	塞卢姆Sollum 	feud.Barony
	图卜鲁格Tobruk 	feud.Barony
	赞祖尔Zanzur 	feud.Barony
}

func (c *图卜鲁格TobrukCounty) BAcroma阿雷马() feud.Barony {
	return c.阿雷马Acroma
}
    
func (c *图卜鲁格TobrukCounty) BAlqardabah艾卡拉达拜() feud.Barony {
	return c.艾卡拉达拜Alqardabah
}
    
func (c *图卜鲁格TobrukCounty) BBardia拜尔迪耶() feud.Barony {
	return c.拜尔迪耶Bardia
}
    
func (c *图卜鲁格TobrukCounty) BKambut坎布特() feud.Barony {
	return c.坎布特Kambut
}
    
func (c *图卜鲁格TobrukCounty) BRukbah鲁克班() feud.Barony {
	return c.鲁克班Rukbah
}
    
func (c *图卜鲁格TobrukCounty) BSollum塞卢姆() feud.Barony {
	return c.塞卢姆Sollum
}
    
func (c *图卜鲁格TobrukCounty) BTobruk图卜鲁格() feud.Barony {
	return c.图卜鲁格Tobruk
}
    
func (c *图卜鲁格TobrukCounty) BZanzur赞祖尔() feud.Barony {
	return c.赞祖尔Zanzur
}
    
var CTobruk图卜鲁格 TobrukCounty = &图卜鲁格TobrukCounty{}

func init() {
	f := CTobruk图卜鲁格.(*图卜鲁格TobrukCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "805",
		Title:     "tobruk",
		TitleName: "图卜鲁格",
		TitleCode: "c_tobruk",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿雷马Acroma = BAcroma阿雷马
	f.阿雷马Acroma.SetParent(f)
	
	f.艾卡拉达拜Alqardabah = BAlqardabah艾卡拉达拜
	f.艾卡拉达拜Alqardabah.SetParent(f)
	
	f.拜尔迪耶Bardia = BBardia拜尔迪耶
	f.拜尔迪耶Bardia.SetParent(f)
	
	f.坎布特Kambut = BKambut坎布特
	f.坎布特Kambut.SetParent(f)
	
	f.鲁克班Rukbah = BRukbah鲁克班
	f.鲁克班Rukbah.SetParent(f)
	
	f.塞卢姆Sollum = BSollum塞卢姆
	f.塞卢姆Sollum.SetParent(f)
	
	f.图卜鲁格Tobruk = BTobruk图卜鲁格
	f.图卜鲁格Tobruk.SetParent(f)
	
	f.赞祖尔Zanzur = BZanzur赞祖尔
	f.赞祖尔Zanzur.SetParent(f)
	
}
