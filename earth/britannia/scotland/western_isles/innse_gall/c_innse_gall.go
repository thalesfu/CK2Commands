package innse_gall

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Innse_gallCounty interface {
    feud.County
    BDunvegan邓韦根() 	feud.Barony
    BKiltaraglen基尔塔拉格伦() 	feud.Barony
    BKisimul基西穆尔() 	feud.Barony
    BSleat斯利特() 	feud.Barony
    BSnizort斯奈泽特() 	feud.Barony
    BStornoway斯托诺韦() 	feud.Barony
    BUig乌伊格() 	feud.Barony
}

type 赫布里底群岛Innse_gallCounty struct {
	feud.BaseCounty
	邓韦根Dunvegan 	feud.Barony
	基尔塔拉格伦Kiltaraglen 	feud.Barony
	基西穆尔Kisimul 	feud.Barony
	斯利特Sleat 	feud.Barony
	斯奈泽特Snizort 	feud.Barony
	斯托诺韦Stornoway 	feud.Barony
	乌伊格Uig 	feud.Barony
}

func (c *赫布里底群岛Innse_gallCounty) BDunvegan邓韦根() feud.Barony {
	return c.邓韦根Dunvegan
}
    
func (c *赫布里底群岛Innse_gallCounty) BKiltaraglen基尔塔拉格伦() feud.Barony {
	return c.基尔塔拉格伦Kiltaraglen
}
    
func (c *赫布里底群岛Innse_gallCounty) BKisimul基西穆尔() feud.Barony {
	return c.基西穆尔Kisimul
}
    
func (c *赫布里底群岛Innse_gallCounty) BSleat斯利特() feud.Barony {
	return c.斯利特Sleat
}
    
func (c *赫布里底群岛Innse_gallCounty) BSnizort斯奈泽特() feud.Barony {
	return c.斯奈泽特Snizort
}
    
func (c *赫布里底群岛Innse_gallCounty) BStornoway斯托诺韦() feud.Barony {
	return c.斯托诺韦Stornoway
}
    
func (c *赫布里底群岛Innse_gallCounty) BUig乌伊格() feud.Barony {
	return c.乌伊格Uig
}
    
var CInnse_gall赫布里底群岛 Innse_gallCounty = &赫布里底群岛Innse_gallCounty{}

func init() {
	f := CInnse_gall赫布里底群岛.(*赫布里底群岛Innse_gallCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "35",
		Title:     "innse_gall",
		TitleName: "赫布里底群岛",
		TitleCode: "c_innse_gall",
		Baronies:  map[string]feud.Barony{},
	}

	f.邓韦根Dunvegan = BDunvegan邓韦根
	f.邓韦根Dunvegan.SetParent(f)
	
	f.基尔塔拉格伦Kiltaraglen = BKiltaraglen基尔塔拉格伦
	f.基尔塔拉格伦Kiltaraglen.SetParent(f)
	
	f.基西穆尔Kisimul = BKisimul基西穆尔
	f.基西穆尔Kisimul.SetParent(f)
	
	f.斯利特Sleat = BSleat斯利特
	f.斯利特Sleat.SetParent(f)
	
	f.斯奈泽特Snizort = BSnizort斯奈泽特
	f.斯奈泽特Snizort.SetParent(f)
	
	f.斯托诺韦Stornoway = BStornoway斯托诺韦
	f.斯托诺韦Stornoway.SetParent(f)
	
	f.乌伊格Uig = BUig乌伊格
	f.乌伊格Uig.SetParent(f)
	
}
