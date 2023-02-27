package kola

import (
	"github.com/thalesfu/CK2Commands/earth/scandinavia/sapmi/kola/kandalax"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/sapmi/kola/kemi"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/sapmi/kola/kola"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/sapmi/kola/rovaniemi"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KolaDuke interface {
    feud.Duke
    CKandalax坎达拉克斯() 	kandalax.KandalaxCounty
    CKemi凯米() 	kemi.KemiCounty
    CKola科拉() 	kola.KolaCounty
    CRovaniemi罗瓦涅米() 	rovaniemi.RovaniemiCounty
}

type 科拉KolaDuke struct {
	feud.BaseDuke
	坎达拉克斯Kandalax 	kandalax.KandalaxCounty
	凯米Kemi 	kemi.KemiCounty
	科拉Kola 	kola.KolaCounty
	罗瓦涅米Rovaniemi 	rovaniemi.RovaniemiCounty
}

func (d *科拉KolaDuke) CKandalax坎达拉克斯() kandalax.KandalaxCounty {
	return d.坎达拉克斯Kandalax
}
    
func (d *科拉KolaDuke) CKemi凯米() kemi.KemiCounty {
	return d.凯米Kemi
}
    
func (d *科拉KolaDuke) CKola科拉() kola.KolaCounty {
	return d.科拉Kola
}
    
func (d *科拉KolaDuke) CRovaniemi罗瓦涅米() rovaniemi.RovaniemiCounty {
	return d.罗瓦涅米Rovaniemi
}
    
var DKola科拉 KolaDuke = &科拉KolaDuke{}

func init() {
	f := DKola科拉.(*科拉KolaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "kola",
		TitleName: "科拉",
		TitleCode: "d_kola",
		Counties:  map[string]feud.County{},
	}

	f.坎达拉克斯Kandalax = kandalax.CKandalax坎达拉克斯
	f.坎达拉克斯Kandalax.SetParent(f)
	
	f.凯米Kemi = kemi.CKemi凯米
	f.凯米Kemi.SetParent(f)
	
	f.科拉Kola = kola.CKola科拉
	f.科拉Kola.SetParent(f)
	
	f.罗瓦涅米Rovaniemi = rovaniemi.CRovaniemi罗瓦涅米
	f.罗瓦涅米Rovaniemi.SetParent(f)
	
}
