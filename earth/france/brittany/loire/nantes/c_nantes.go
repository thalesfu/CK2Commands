package nantes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NantesCounty interface {
    feud.County
    BChateaubriant沙托布里昂() 	feud.Barony
    BDonges栋日() 	feud.Barony
    BGuerande盖朗德() 	feud.Barony
    BLa_roche_bernard拉罗什_贝尔纳() 	feud.Barony
    BNantes南特() 	feud.Barony
    BRedon勒东() 	feud.Barony
    BSt_nazaire圣纳泽尔() 	feud.Barony
}

type 南特NantesCounty struct {
	feud.BaseCounty
	沙托布里昂Chateaubriant 	feud.Barony
	栋日Donges 	feud.Barony
	盖朗德Guerande 	feud.Barony
	拉罗什_贝尔纳La_roche_bernard 	feud.Barony
	南特Nantes 	feud.Barony
	勒东Redon 	feud.Barony
	圣纳泽尔St_nazaire 	feud.Barony
}

func (c *南特NantesCounty) BChateaubriant沙托布里昂() feud.Barony {
	return c.沙托布里昂Chateaubriant
}
    
func (c *南特NantesCounty) BDonges栋日() feud.Barony {
	return c.栋日Donges
}
    
func (c *南特NantesCounty) BGuerande盖朗德() feud.Barony {
	return c.盖朗德Guerande
}
    
func (c *南特NantesCounty) BLa_roche_bernard拉罗什_贝尔纳() feud.Barony {
	return c.拉罗什_贝尔纳La_roche_bernard
}
    
func (c *南特NantesCounty) BNantes南特() feud.Barony {
	return c.南特Nantes
}
    
func (c *南特NantesCounty) BRedon勒东() feud.Barony {
	return c.勒东Redon
}
    
func (c *南特NantesCounty) BSt_nazaire圣纳泽尔() feud.Barony {
	return c.圣纳泽尔St_nazaire
}
    
var CNantes南特 NantesCounty = &南特NantesCounty{}

func init() {
	f := CNantes南特.(*南特NantesCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "106",
		Title:     "nantes",
		TitleName: "南特",
		TitleCode: "c_nantes",
		Baronies:  map[string]feud.Barony{},
	}

	f.沙托布里昂Chateaubriant = BChateaubriant沙托布里昂
	f.沙托布里昂Chateaubriant.SetParent(f)
	
	f.栋日Donges = BDonges栋日
	f.栋日Donges.SetParent(f)
	
	f.盖朗德Guerande = BGuerande盖朗德
	f.盖朗德Guerande.SetParent(f)
	
	f.拉罗什_贝尔纳La_roche_bernard = BLa_roche_bernard拉罗什_贝尔纳
	f.拉罗什_贝尔纳La_roche_bernard.SetParent(f)
	
	f.南特Nantes = BNantes南特
	f.南特Nantes.SetParent(f)
	
	f.勒东Redon = BRedon勒东
	f.勒东Redon.SetParent(f)
	
	f.圣纳泽尔St_nazaire = BSt_nazaire圣纳泽尔
	f.圣纳泽尔St_nazaire.SetParent(f)
	
}
