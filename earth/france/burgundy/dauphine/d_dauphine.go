package dauphine

import (
	"github.com/thalesfu/CK2Commands/earth/france/burgundy/dauphine/dauphine_viennois"
	"github.com/thalesfu/CK2Commands/earth/france/burgundy/dauphine/forez"
	"github.com/thalesfu/CK2Commands/earth/france/burgundy/dauphine/lyon"
	"github.com/thalesfu/CK2Commands/earth/france/burgundy/dauphine/viviers"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DauphineDuke interface {
    feud.Duke
    CDauphine_viennois多菲内() 	dauphine_viennois.Dauphine_viennoisCounty
    CForez福雷() 	forez.ForezCounty
    CLyon里昂() 	lyon.LyonCounty
    CViviers维瓦赖() 	viviers.ViviersCounty
}

type 多菲内DauphineDuke struct {
	feud.BaseDuke
	多菲内Dauphine_viennois 	dauphine_viennois.Dauphine_viennoisCounty
	福雷Forez 	forez.ForezCounty
	里昂Lyon 	lyon.LyonCounty
	维瓦赖Viviers 	viviers.ViviersCounty
}

func (d *多菲内DauphineDuke) CDauphine_viennois多菲内() dauphine_viennois.Dauphine_viennoisCounty {
	return d.多菲内Dauphine_viennois
}
    
func (d *多菲内DauphineDuke) CForez福雷() forez.ForezCounty {
	return d.福雷Forez
}
    
func (d *多菲内DauphineDuke) CLyon里昂() lyon.LyonCounty {
	return d.里昂Lyon
}
    
func (d *多菲内DauphineDuke) CViviers维瓦赖() viviers.ViviersCounty {
	return d.维瓦赖Viviers
}
    
var DDauphine多菲内 DauphineDuke = &多菲内DauphineDuke{}

func init() {
	f := DDauphine多菲内.(*多菲内DauphineDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "dauphine",
		TitleName: "多菲内",
		TitleCode: "d_dauphine",
		Counties:  map[string]feud.County{},
	}

	f.多菲内Dauphine_viennois = dauphine_viennois.CDauphine_viennois多菲内
	f.多菲内Dauphine_viennois.SetParent(f)
	
	f.福雷Forez = forez.CForez福雷
	f.福雷Forez.SetParent(f)
	
	f.里昂Lyon = lyon.CLyon里昂
	f.里昂Lyon.SetParent(f)
	
	f.维瓦赖Viviers = viviers.CViviers维瓦赖
	f.维瓦赖Viviers.SetParent(f)
	
}
