package rennes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RennesCounty interface {
    feud.County
    BDinan迪南() 	feud.Barony
    BDol多勒() 	feud.Barony
    BFougeres富热尔() 	feud.Barony
    BPorhoet波尔厄() 	feud.Barony
    BRennes雷恩() 	feud.Barony
    BSt_malo圣马洛() 	feud.Barony
    BSt_meen圣梅昂() 	feud.Barony
    BSt_michel圣米歇尔() 	feud.Barony
}

type 雷恩RennesCounty struct {
	feud.BaseCounty
	迪南Dinan 	feud.Barony
	多勒Dol 	feud.Barony
	富热尔Fougeres 	feud.Barony
	波尔厄Porhoet 	feud.Barony
	雷恩Rennes 	feud.Barony
	圣马洛St_malo 	feud.Barony
	圣梅昂St_meen 	feud.Barony
	圣米歇尔St_michel 	feud.Barony
}

func (c *雷恩RennesCounty) BDinan迪南() feud.Barony {
	return c.迪南Dinan
}
    
func (c *雷恩RennesCounty) BDol多勒() feud.Barony {
	return c.多勒Dol
}
    
func (c *雷恩RennesCounty) BFougeres富热尔() feud.Barony {
	return c.富热尔Fougeres
}
    
func (c *雷恩RennesCounty) BPorhoet波尔厄() feud.Barony {
	return c.波尔厄Porhoet
}
    
func (c *雷恩RennesCounty) BRennes雷恩() feud.Barony {
	return c.雷恩Rennes
}
    
func (c *雷恩RennesCounty) BSt_malo圣马洛() feud.Barony {
	return c.圣马洛St_malo
}
    
func (c *雷恩RennesCounty) BSt_meen圣梅昂() feud.Barony {
	return c.圣梅昂St_meen
}
    
func (c *雷恩RennesCounty) BSt_michel圣米歇尔() feud.Barony {
	return c.圣米歇尔St_michel
}
    
var CRennes雷恩 RennesCounty = &雷恩RennesCounty{}

func init() {
	f := CRennes雷恩.(*雷恩RennesCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "101",
		Title:     "rennes",
		TitleName: "雷恩",
		TitleCode: "c_rennes",
		Baronies:  map[string]feud.Barony{},
	}

	f.迪南Dinan = BDinan迪南
	f.迪南Dinan.SetParent(f)
	
	f.多勒Dol = BDol多勒
	f.多勒Dol.SetParent(f)
	
	f.富热尔Fougeres = BFougeres富热尔
	f.富热尔Fougeres.SetParent(f)
	
	f.波尔厄Porhoet = BPorhoet波尔厄
	f.波尔厄Porhoet.SetParent(f)
	
	f.雷恩Rennes = BRennes雷恩
	f.雷恩Rennes.SetParent(f)
	
	f.圣马洛St_malo = BSt_malo圣马洛
	f.圣马洛St_malo.SetParent(f)
	
	f.圣梅昂St_meen = BSt_meen圣梅昂
	f.圣梅昂St_meen.SetParent(f)
	
	f.圣米歇尔St_michel = BSt_michel圣米歇尔
	f.圣米歇尔St_michel.SetParent(f)
	
}
