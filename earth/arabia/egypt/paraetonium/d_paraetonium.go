package paraetonium

import (
	"github.com/thalesfu/CK2Commands/earth/arabia/egypt/paraetonium/al_alamayn"
	"github.com/thalesfu/CK2Commands/earth/arabia/egypt/paraetonium/paraetonium"
	"github.com/thalesfu/CK2Commands/earth/arabia/egypt/paraetonium/quattara"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ParaetoniumDuke interface {
    feud.Duke
    CAl_alamayn卒格拉() 	al_alamayn.Al_alamaynCounty
    CParaetonium帕莱托尼翁() 	paraetonium.ParaetoniumCounty
    CQuattara盖塔拉() 	quattara.QuattaraCounty
}

type 帕莱托尼翁ParaetoniumDuke struct {
	feud.BaseDuke
	卒格拉Al_alamayn 	al_alamayn.Al_alamaynCounty
	帕莱托尼翁Paraetonium 	paraetonium.ParaetoniumCounty
	盖塔拉Quattara 	quattara.QuattaraCounty
}

func (d *帕莱托尼翁ParaetoniumDuke) CAl_alamayn卒格拉() al_alamayn.Al_alamaynCounty {
	return d.卒格拉Al_alamayn
}
    
func (d *帕莱托尼翁ParaetoniumDuke) CParaetonium帕莱托尼翁() paraetonium.ParaetoniumCounty {
	return d.帕莱托尼翁Paraetonium
}
    
func (d *帕莱托尼翁ParaetoniumDuke) CQuattara盖塔拉() quattara.QuattaraCounty {
	return d.盖塔拉Quattara
}
    
var DParaetonium帕莱托尼翁 ParaetoniumDuke = &帕莱托尼翁ParaetoniumDuke{}

func init() {
	f := DParaetonium帕莱托尼翁.(*帕莱托尼翁ParaetoniumDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "paraetonium",
		TitleName: "帕莱托尼翁",
		TitleCode: "d_paraetonium",
		Counties:  map[string]feud.County{},
	}

	f.卒格拉Al_alamayn = al_alamayn.CAl_alamayn卒格拉
	f.卒格拉Al_alamayn.SetParent(f)
	
	f.帕莱托尼翁Paraetonium = paraetonium.CParaetonium帕莱托尼翁
	f.帕莱托尼翁Paraetonium.SetParent(f)
	
	f.盖塔拉Quattara = quattara.CQuattara盖塔拉
	f.盖塔拉Quattara.SetParent(f)
	
}
