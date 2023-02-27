package yuni

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type YuniCounty interface {
    feud.County
    BBeilanzhuang北兰庄() 	feud.Barony
    BHuachuantsun花川村() 	feud.Barony
    BPailouxia牌楼下() 	feud.Barony
    BRuoqiang婼羌() 	feud.Barony
    BShikeng石坑() 	feud.Barony
    BXiafang狭方() 	feud.Barony
    BYuni扜泥() 	feud.Barony
}

type 扜泥YuniCounty struct {
	feud.BaseCounty
	北兰庄Beilanzhuang 	feud.Barony
	花川村Huachuantsun 	feud.Barony
	牌楼下Pailouxia 	feud.Barony
	婼羌Ruoqiang 	feud.Barony
	石坑Shikeng 	feud.Barony
	狭方Xiafang 	feud.Barony
	扜泥Yuni 	feud.Barony
}

func (c *扜泥YuniCounty) BBeilanzhuang北兰庄() feud.Barony {
	return c.北兰庄Beilanzhuang
}
    
func (c *扜泥YuniCounty) BHuachuantsun花川村() feud.Barony {
	return c.花川村Huachuantsun
}
    
func (c *扜泥YuniCounty) BPailouxia牌楼下() feud.Barony {
	return c.牌楼下Pailouxia
}
    
func (c *扜泥YuniCounty) BRuoqiang婼羌() feud.Barony {
	return c.婼羌Ruoqiang
}
    
func (c *扜泥YuniCounty) BShikeng石坑() feud.Barony {
	return c.石坑Shikeng
}
    
func (c *扜泥YuniCounty) BXiafang狭方() feud.Barony {
	return c.狭方Xiafang
}
    
func (c *扜泥YuniCounty) BYuni扜泥() feud.Barony {
	return c.扜泥Yuni
}
    
var CYuni扜泥 YuniCounty = &扜泥YuniCounty{}

func init() {
	f := CYuni扜泥.(*扜泥YuniCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1520",
		Title:     "yuni",
		TitleName: "扜泥",
		TitleCode: "c_yuni",
		Baronies:  map[string]feud.Barony{},
	}

	f.北兰庄Beilanzhuang = BBeilanzhuang北兰庄
	f.北兰庄Beilanzhuang.SetParent(f)
	
	f.花川村Huachuantsun = BHuachuantsun花川村
	f.花川村Huachuantsun.SetParent(f)
	
	f.牌楼下Pailouxia = BPailouxia牌楼下
	f.牌楼下Pailouxia.SetParent(f)
	
	f.婼羌Ruoqiang = BRuoqiang婼羌
	f.婼羌Ruoqiang.SetParent(f)
	
	f.石坑Shikeng = BShikeng石坑
	f.石坑Shikeng.SetParent(f)
	
	f.狭方Xiafang = BXiafang狭方
	f.狭方Xiafang.SetParent(f)
	
	f.扜泥Yuni = BYuni扜泥
	f.扜泥Yuni.SetParent(f)
	
}
