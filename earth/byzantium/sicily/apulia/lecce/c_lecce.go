package lecce

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LecceCounty interface {
    feud.County
    BAndrano安德拉诺() 	feud.Barony
    BBrindisi布林迪西() 	feud.Barony
    BCastro卡斯特罗() 	feud.Barony
    BLecce莱切() 	feud.Barony
    BLeuca莱乌卡() 	feud.Barony
    BLigento利真托() 	feud.Barony
    BOria奥里亚() 	feud.Barony
    BOtranto奥特朗托() 	feud.Barony
}

type 莱切LecceCounty struct {
	feud.BaseCounty
	安德拉诺Andrano 	feud.Barony
	布林迪西Brindisi 	feud.Barony
	卡斯特罗Castro 	feud.Barony
	莱切Lecce 	feud.Barony
	莱乌卡Leuca 	feud.Barony
	利真托Ligento 	feud.Barony
	奥里亚Oria 	feud.Barony
	奥特朗托Otranto 	feud.Barony
}

func (c *莱切LecceCounty) BAndrano安德拉诺() feud.Barony {
	return c.安德拉诺Andrano
}
    
func (c *莱切LecceCounty) BBrindisi布林迪西() feud.Barony {
	return c.布林迪西Brindisi
}
    
func (c *莱切LecceCounty) BCastro卡斯特罗() feud.Barony {
	return c.卡斯特罗Castro
}
    
func (c *莱切LecceCounty) BLecce莱切() feud.Barony {
	return c.莱切Lecce
}
    
func (c *莱切LecceCounty) BLeuca莱乌卡() feud.Barony {
	return c.莱乌卡Leuca
}
    
func (c *莱切LecceCounty) BLigento利真托() feud.Barony {
	return c.利真托Ligento
}
    
func (c *莱切LecceCounty) BOria奥里亚() feud.Barony {
	return c.奥里亚Oria
}
    
func (c *莱切LecceCounty) BOtranto奥特朗托() feud.Barony {
	return c.奥特朗托Otranto
}
    
var CLecce莱切 LecceCounty = &莱切LecceCounty{}

func init() {
	f := CLecce莱切.(*莱切LecceCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "345",
		Title:     "lecce",
		TitleName: "莱切",
		TitleCode: "c_lecce",
		Baronies:  map[string]feud.Barony{},
	}

	f.安德拉诺Andrano = BAndrano安德拉诺
	f.安德拉诺Andrano.SetParent(f)
	
	f.布林迪西Brindisi = BBrindisi布林迪西
	f.布林迪西Brindisi.SetParent(f)
	
	f.卡斯特罗Castro = BCastro卡斯特罗
	f.卡斯特罗Castro.SetParent(f)
	
	f.莱切Lecce = BLecce莱切
	f.莱切Lecce.SetParent(f)
	
	f.莱乌卡Leuca = BLeuca莱乌卡
	f.莱乌卡Leuca.SetParent(f)
	
	f.利真托Ligento = BLigento利真托
	f.利真托Ligento.SetParent(f)
	
	f.奥里亚Oria = BOria奥里亚
	f.奥里亚Oria.SetParent(f)
	
	f.奥特朗托Otranto = BOtranto奥特朗托
	f.奥特朗托Otranto.SetParent(f)
	
}
