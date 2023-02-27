package irtek

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type IrtekCounty interface {
    feud.County
    BBorili伯里利() 	feud.Barony
    BChirovo奇罗沃() 	feud.Barony
    BIrtek伊尔捷克() 	feud.Barony
    BKhamino哈米诺() 	feud.Barony
    BKirsanovo基尔萨诺沃() 	feud.Barony
    BQanai哈奈() 	feud.Barony
    BTeploye捷普洛耶() 	feud.Barony
}

type 伊尔捷克IrtekCounty struct {
	feud.BaseCounty
	伯里利Borili 	feud.Barony
	奇罗沃Chirovo 	feud.Barony
	伊尔捷克Irtek 	feud.Barony
	哈米诺Khamino 	feud.Barony
	基尔萨诺沃Kirsanovo 	feud.Barony
	哈奈Qanai 	feud.Barony
	捷普洛耶Teploye 	feud.Barony
}

func (c *伊尔捷克IrtekCounty) BBorili伯里利() feud.Barony {
	return c.伯里利Borili
}
    
func (c *伊尔捷克IrtekCounty) BChirovo奇罗沃() feud.Barony {
	return c.奇罗沃Chirovo
}
    
func (c *伊尔捷克IrtekCounty) BIrtek伊尔捷克() feud.Barony {
	return c.伊尔捷克Irtek
}
    
func (c *伊尔捷克IrtekCounty) BKhamino哈米诺() feud.Barony {
	return c.哈米诺Khamino
}
    
func (c *伊尔捷克IrtekCounty) BKirsanovo基尔萨诺沃() feud.Barony {
	return c.基尔萨诺沃Kirsanovo
}
    
func (c *伊尔捷克IrtekCounty) BQanai哈奈() feud.Barony {
	return c.哈奈Qanai
}
    
func (c *伊尔捷克IrtekCounty) BTeploye捷普洛耶() feud.Barony {
	return c.捷普洛耶Teploye
}
    
var CIrtek伊尔捷克 IrtekCounty = &伊尔捷克IrtekCounty{}

func init() {
	f := CIrtek伊尔捷克.(*伊尔捷克IrtekCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1849",
		Title:     "irtek",
		TitleName: "伊尔捷克",
		TitleCode: "c_irtek",
		Baronies:  map[string]feud.Barony{},
	}

	f.伯里利Borili = BBorili伯里利
	f.伯里利Borili.SetParent(f)
	
	f.奇罗沃Chirovo = BChirovo奇罗沃
	f.奇罗沃Chirovo.SetParent(f)
	
	f.伊尔捷克Irtek = BIrtek伊尔捷克
	f.伊尔捷克Irtek.SetParent(f)
	
	f.哈米诺Khamino = BKhamino哈米诺
	f.哈米诺Khamino.SetParent(f)
	
	f.基尔萨诺沃Kirsanovo = BKirsanovo基尔萨诺沃
	f.基尔萨诺沃Kirsanovo.SetParent(f)
	
	f.哈奈Qanai = BQanai哈奈
	f.哈奈Qanai.SetParent(f)
	
	f.捷普洛耶Teploye = BTeploye捷普洛耶
	f.捷普洛耶Teploye.SetParent(f)
	
}
