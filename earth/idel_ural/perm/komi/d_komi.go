package komi

import (
	"github.com/thalesfu/CK2Commands/earth/idel_ural/perm/komi/izhma"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/perm/komi/kozhva"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/perm/komi/lyzha"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/perm/komi/severnaya"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/perm/komi/vel"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KomiDuke interface {
    feud.Duke
    CIzhma伊日马() 	izhma.IzhmaCounty
    CKozhva科日瓦() 	kozhva.KozhvaCounty
    CLyzha雷扎() 	lyzha.LyzhaCounty
    CSevernaya谢韦尔纳亚() 	severnaya.SevernayaCounty
    CVel韦尔() 	vel.VelCounty
}

type 科米KomiDuke struct {
	feud.BaseDuke
	伊日马Izhma 	izhma.IzhmaCounty
	科日瓦Kozhva 	kozhva.KozhvaCounty
	雷扎Lyzha 	lyzha.LyzhaCounty
	谢韦尔纳亚Severnaya 	severnaya.SevernayaCounty
	韦尔Vel 	vel.VelCounty
}

func (d *科米KomiDuke) CIzhma伊日马() izhma.IzhmaCounty {
	return d.伊日马Izhma
}
    
func (d *科米KomiDuke) CKozhva科日瓦() kozhva.KozhvaCounty {
	return d.科日瓦Kozhva
}
    
func (d *科米KomiDuke) CLyzha雷扎() lyzha.LyzhaCounty {
	return d.雷扎Lyzha
}
    
func (d *科米KomiDuke) CSevernaya谢韦尔纳亚() severnaya.SevernayaCounty {
	return d.谢韦尔纳亚Severnaya
}
    
func (d *科米KomiDuke) CVel韦尔() vel.VelCounty {
	return d.韦尔Vel
}
    
var DKomi科米 KomiDuke = &科米KomiDuke{}

func init() {
	f := DKomi科米.(*科米KomiDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "komi",
		TitleName: "科米",
		TitleCode: "d_komi",
		Counties:  map[string]feud.County{},
	}

	f.伊日马Izhma = izhma.CIzhma伊日马
	f.伊日马Izhma.SetParent(f)
	
	f.科日瓦Kozhva = kozhva.CKozhva科日瓦
	f.科日瓦Kozhva.SetParent(f)
	
	f.雷扎Lyzha = lyzha.CLyzha雷扎
	f.雷扎Lyzha.SetParent(f)
	
	f.谢韦尔纳亚Severnaya = severnaya.CSevernaya谢韦尔纳亚
	f.谢韦尔纳亚Severnaya.SetParent(f)
	
	f.韦尔Vel = vel.CVel韦尔
	f.韦尔Vel.SetParent(f)
	
}
