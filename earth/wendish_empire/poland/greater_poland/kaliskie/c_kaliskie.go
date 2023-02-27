package kaliskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KaliskieCounty interface {
    feud.County
    BGniezno格涅兹诺() 	feud.Barony
    BKalisz卡利什() 	feud.Barony
    BKolo科沃() 	feud.Barony
    BKonin科宁() 	feud.Barony
    BKozmin科伊明() 	feud.Barony
    BPyzdry佩兹德雷() 	feud.Barony
    BSroda希罗达() 	feud.Barony
}

type 卡利什KaliskieCounty struct {
	feud.BaseCounty
	格涅兹诺Gniezno 	feud.Barony
	卡利什Kalisz 	feud.Barony
	科沃Kolo 	feud.Barony
	科宁Konin 	feud.Barony
	科伊明Kozmin 	feud.Barony
	佩兹德雷Pyzdry 	feud.Barony
	希罗达Sroda 	feud.Barony
}

func (c *卡利什KaliskieCounty) BGniezno格涅兹诺() feud.Barony {
	return c.格涅兹诺Gniezno
}
    
func (c *卡利什KaliskieCounty) BKalisz卡利什() feud.Barony {
	return c.卡利什Kalisz
}
    
func (c *卡利什KaliskieCounty) BKolo科沃() feud.Barony {
	return c.科沃Kolo
}
    
func (c *卡利什KaliskieCounty) BKonin科宁() feud.Barony {
	return c.科宁Konin
}
    
func (c *卡利什KaliskieCounty) BKozmin科伊明() feud.Barony {
	return c.科伊明Kozmin
}
    
func (c *卡利什KaliskieCounty) BPyzdry佩兹德雷() feud.Barony {
	return c.佩兹德雷Pyzdry
}
    
func (c *卡利什KaliskieCounty) BSroda希罗达() feud.Barony {
	return c.希罗达Sroda
}
    
var CKaliskie卡利什 KaliskieCounty = &卡利什KaliskieCounty{}

func init() {
	f := CKaliskie卡利什.(*卡利什KaliskieCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "432",
		Title:     "kaliskie",
		TitleName: "卡利什",
		TitleCode: "c_kaliskie",
		Baronies:  map[string]feud.Barony{},
	}

	f.格涅兹诺Gniezno = BGniezno格涅兹诺
	f.格涅兹诺Gniezno.SetParent(f)
	
	f.卡利什Kalisz = BKalisz卡利什
	f.卡利什Kalisz.SetParent(f)
	
	f.科沃Kolo = BKolo科沃
	f.科沃Kolo.SetParent(f)
	
	f.科宁Konin = BKonin科宁
	f.科宁Konin.SetParent(f)
	
	f.科伊明Kozmin = BKozmin科伊明
	f.科伊明Kozmin.SetParent(f)
	
	f.佩兹德雷Pyzdry = BPyzdry佩兹德雷
	f.佩兹德雷Pyzdry.SetParent(f)
	
	f.希罗达Sroda = BSroda希罗达
	f.希罗达Sroda.SetParent(f)
	
}
