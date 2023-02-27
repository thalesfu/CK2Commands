package franche_comte

import (
	"github.com/thalesfu/CK2Commands/earth/france/burgundy/franche_comte/amous"
	"github.com/thalesfu/CK2Commands/earth/france/burgundy/franche_comte/besancon"
	"github.com/thalesfu/CK2Commands/earth/france/burgundy/franche_comte/escuens"
	"github.com/thalesfu/CK2Commands/earth/france/burgundy/franche_comte/montbeliard"
	"github.com/thalesfu/CK2Commands/earth/france/burgundy/franche_comte/vesoul"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Franche_comteDuke interface {
    feud.Duke
    CAmous阿穆() 	amous.AmousCounty
    CBesancon贝桑松() 	besancon.BesanconCounty
    CEscuens埃斯库昂() 	escuens.EscuensCounty
    CMontbeliard蒙贝利亚尔() 	montbeliard.MontbeliardCounty
    CVesoul波图瓦() 	vesoul.VesoulCounty
}

type 弗朗什孔泰Franche_comteDuke struct {
	feud.BaseDuke
	阿穆Amous 	amous.AmousCounty
	贝桑松Besancon 	besancon.BesanconCounty
	埃斯库昂Escuens 	escuens.EscuensCounty
	蒙贝利亚尔Montbeliard 	montbeliard.MontbeliardCounty
	波图瓦Vesoul 	vesoul.VesoulCounty
}

func (d *弗朗什孔泰Franche_comteDuke) CAmous阿穆() amous.AmousCounty {
	return d.阿穆Amous
}
    
func (d *弗朗什孔泰Franche_comteDuke) CBesancon贝桑松() besancon.BesanconCounty {
	return d.贝桑松Besancon
}
    
func (d *弗朗什孔泰Franche_comteDuke) CEscuens埃斯库昂() escuens.EscuensCounty {
	return d.埃斯库昂Escuens
}
    
func (d *弗朗什孔泰Franche_comteDuke) CMontbeliard蒙贝利亚尔() montbeliard.MontbeliardCounty {
	return d.蒙贝利亚尔Montbeliard
}
    
func (d *弗朗什孔泰Franche_comteDuke) CVesoul波图瓦() vesoul.VesoulCounty {
	return d.波图瓦Vesoul
}
    
var DFranche_comte弗朗什孔泰 Franche_comteDuke = &弗朗什孔泰Franche_comteDuke{}

func init() {
	f := DFranche_comte弗朗什孔泰.(*弗朗什孔泰Franche_comteDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "franche_comte",
		TitleName: "弗朗什孔泰",
		TitleCode: "d_franche_comte",
		Counties:  map[string]feud.County{},
	}

	f.阿穆Amous = amous.CAmous阿穆
	f.阿穆Amous.SetParent(f)
	
	f.贝桑松Besancon = besancon.CBesancon贝桑松
	f.贝桑松Besancon.SetParent(f)
	
	f.埃斯库昂Escuens = escuens.CEscuens埃斯库昂
	f.埃斯库昂Escuens.SetParent(f)
	
	f.蒙贝利亚尔Montbeliard = montbeliard.CMontbeliard蒙贝利亚尔
	f.蒙贝利亚尔Montbeliard.SetParent(f)
	
	f.波图瓦Vesoul = vesoul.CVesoul波图瓦
	f.波图瓦Vesoul.SetParent(f)
	
}
