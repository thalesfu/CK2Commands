package crimea

import (
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/crimea/crimea/crimea"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/crimea/crimea/lower_dniepr"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/crimea/crimea/lower_don"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/crimea/crimea/lukomorie"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CrimeaDuke interface {
    feud.Duke
    CCrimea克里米亚() 	crimea.CrimeaCounty
    CLower_dniepr奥列什基() 	lower_dniepr.Lower_dnieprCounty
    CLower_don下顿河() 	lower_don.Lower_donCounty
    CLukomorie卢科莫里耶() 	lukomorie.LukomorieCounty
}

type 克里米亚CrimeaDuke struct {
	feud.BaseDuke
	克里米亚Crimea 	crimea.CrimeaCounty
	奥列什基Lower_dniepr 	lower_dniepr.Lower_dnieprCounty
	下顿河Lower_don 	lower_don.Lower_donCounty
	卢科莫里耶Lukomorie 	lukomorie.LukomorieCounty
}

func (d *克里米亚CrimeaDuke) CCrimea克里米亚() crimea.CrimeaCounty {
	return d.克里米亚Crimea
}
    
func (d *克里米亚CrimeaDuke) CLower_dniepr奥列什基() lower_dniepr.Lower_dnieprCounty {
	return d.奥列什基Lower_dniepr
}
    
func (d *克里米亚CrimeaDuke) CLower_don下顿河() lower_don.Lower_donCounty {
	return d.下顿河Lower_don
}
    
func (d *克里米亚CrimeaDuke) CLukomorie卢科莫里耶() lukomorie.LukomorieCounty {
	return d.卢科莫里耶Lukomorie
}
    
var DCrimea克里米亚 CrimeaDuke = &克里米亚CrimeaDuke{}

func init() {
	f := DCrimea克里米亚.(*克里米亚CrimeaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "crimea",
		TitleName: "克里米亚",
		TitleCode: "d_crimea",
		Counties:  map[string]feud.County{},
	}

	f.克里米亚Crimea = crimea.CCrimea克里米亚
	f.克里米亚Crimea.SetParent(f)
	
	f.奥列什基Lower_dniepr = lower_dniepr.CLower_dniepr奥列什基
	f.奥列什基Lower_dniepr.SetParent(f)
	
	f.下顿河Lower_don = lower_don.CLower_don下顿河
	f.下顿河Lower_don.SetParent(f)
	
	f.卢科莫里耶Lukomorie = lukomorie.CLukomorie卢科莫里耶
	f.卢科莫里耶Lukomorie.SetParent(f)
	
}
