package usora

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type UsoraCounty interface {
    feud.County
    BBanjaluka巴尼亚卢卡() 	feud.Barony
    BBihac比哈奇() 	feud.Barony
    BBocac博查茨() 	feud.Barony
    BGlamoc格拉莫奇() 	feud.Barony
    BJajce亚伊采() 	feud.Barony
    BKljuc克柳奇() 	feud.Barony
    BKotor_donjikraji科托尔() 	feud.Barony
    BPrijedor普里耶多尔() 	feud.Barony
}

type 下波斯尼亚UsoraCounty struct {
	feud.BaseCounty
	巴尼亚卢卡Banjaluka 	feud.Barony
	比哈奇Bihac 	feud.Barony
	博查茨Bocac 	feud.Barony
	格拉莫奇Glamoc 	feud.Barony
	亚伊采Jajce 	feud.Barony
	克柳奇Kljuc 	feud.Barony
	科托尔Kotor_donjikraji 	feud.Barony
	普里耶多尔Prijedor 	feud.Barony
}

func (c *下波斯尼亚UsoraCounty) BBanjaluka巴尼亚卢卡() feud.Barony {
	return c.巴尼亚卢卡Banjaluka
}
    
func (c *下波斯尼亚UsoraCounty) BBihac比哈奇() feud.Barony {
	return c.比哈奇Bihac
}
    
func (c *下波斯尼亚UsoraCounty) BBocac博查茨() feud.Barony {
	return c.博查茨Bocac
}
    
func (c *下波斯尼亚UsoraCounty) BGlamoc格拉莫奇() feud.Barony {
	return c.格拉莫奇Glamoc
}
    
func (c *下波斯尼亚UsoraCounty) BJajce亚伊采() feud.Barony {
	return c.亚伊采Jajce
}
    
func (c *下波斯尼亚UsoraCounty) BKljuc克柳奇() feud.Barony {
	return c.克柳奇Kljuc
}
    
func (c *下波斯尼亚UsoraCounty) BKotor_donjikraji科托尔() feud.Barony {
	return c.科托尔Kotor_donjikraji
}
    
func (c *下波斯尼亚UsoraCounty) BPrijedor普里耶多尔() feud.Barony {
	return c.普里耶多尔Prijedor
}
    
var CUsora下波斯尼亚 UsoraCounty = &下波斯尼亚UsoraCounty{}

func init() {
	f := CUsora下波斯尼亚.(*下波斯尼亚UsoraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "463",
		Title:     "usora",
		TitleName: "下波斯尼亚",
		TitleCode: "c_usora",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴尼亚卢卡Banjaluka = BBanjaluka巴尼亚卢卡
	f.巴尼亚卢卡Banjaluka.SetParent(f)
	
	f.比哈奇Bihac = BBihac比哈奇
	f.比哈奇Bihac.SetParent(f)
	
	f.博查茨Bocac = BBocac博查茨
	f.博查茨Bocac.SetParent(f)
	
	f.格拉莫奇Glamoc = BGlamoc格拉莫奇
	f.格拉莫奇Glamoc.SetParent(f)
	
	f.亚伊采Jajce = BJajce亚伊采
	f.亚伊采Jajce.SetParent(f)
	
	f.克柳奇Kljuc = BKljuc克柳奇
	f.克柳奇Kljuc.SetParent(f)
	
	f.科托尔Kotor_donjikraji = BKotor_donjikraji科托尔
	f.科托尔Kotor_donjikraji.SetParent(f)
	
	f.普里耶多尔Prijedor = BPrijedor普里耶多尔
	f.普里耶多尔Prijedor.SetParent(f)
	
}
