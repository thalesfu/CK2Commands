package beja

import (
	"github.com/thalesfu/CK2Commands/earth/spain/badajoz/beja/alcacer_do_sal"
	"github.com/thalesfu/CK2Commands/earth/spain/badajoz/beja/egitanea"
	"github.com/thalesfu/CK2Commands/earth/spain/badajoz/beja/evora"
	"github.com/thalesfu/CK2Commands/earth/spain/badajoz/beja/mertola"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BejaDuke interface {
    feud.Duke
    CAlcacer_do_sal萨尔堡() 	alcacer_do_sal.Alcacer_do_salCounty
    CEgitaneac_egitanea() 	egitanea.EgitaneaCounty
    CEvora埃武拉() 	evora.EvoraCounty
    CMertola梅尔图拉() 	mertola.MertolaCounty
}

type 贝雅BejaDuke struct {
	feud.BaseDuke
	萨尔堡Alcacer_do_sal 	alcacer_do_sal.Alcacer_do_salCounty
	c_egitaneaEgitanea 	egitanea.EgitaneaCounty
	埃武拉Evora 	evora.EvoraCounty
	梅尔图拉Mertola 	mertola.MertolaCounty
}

func (d *贝雅BejaDuke) CAlcacer_do_sal萨尔堡() alcacer_do_sal.Alcacer_do_salCounty {
	return d.萨尔堡Alcacer_do_sal
}
    
func (d *贝雅BejaDuke) CEgitaneac_egitanea() egitanea.EgitaneaCounty {
	return d.c_egitaneaEgitanea
}
    
func (d *贝雅BejaDuke) CEvora埃武拉() evora.EvoraCounty {
	return d.埃武拉Evora
}
    
func (d *贝雅BejaDuke) CMertola梅尔图拉() mertola.MertolaCounty {
	return d.梅尔图拉Mertola
}
    
var DBeja贝雅 BejaDuke = &贝雅BejaDuke{}

func init() {
	f := DBeja贝雅.(*贝雅BejaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "beja",
		TitleName: "贝雅",
		TitleCode: "d_beja",
		Counties:  map[string]feud.County{},
	}

	f.萨尔堡Alcacer_do_sal = alcacer_do_sal.CAlcacer_do_sal萨尔堡
	f.萨尔堡Alcacer_do_sal.SetParent(f)
	
	f.c_egitaneaEgitanea = egitanea.CEgitaneac_egitanea
	f.c_egitaneaEgitanea.SetParent(f)
	
	f.埃武拉Evora = evora.CEvora埃武拉
	f.埃武拉Evora.SetParent(f)
	
	f.梅尔图拉Mertola = mertola.CMertola梅尔图拉
	f.梅尔图拉Mertola.SetParent(f)
	
}
