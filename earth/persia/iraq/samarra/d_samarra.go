package samarra

import (
	"github.com/thalesfu/CK2Commands/earth/persia/iraq/samarra/deir"
	"github.com/thalesfu/CK2Commands/earth/persia/iraq/samarra/euphrates"
	"github.com/thalesfu/CK2Commands/earth/persia/iraq/samarra/kirkuk"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SamarraDuke interface {
    feud.Duke
    CDeir代尔() 	deir.DeirCounty
    CEuphrates萨迈拉() 	euphrates.EuphratesCounty
    CKirkuk基尔库克() 	kirkuk.KirkukCounty
}

type 萨迈拉SamarraDuke struct {
	feud.BaseDuke
	代尔Deir 	deir.DeirCounty
	萨迈拉Euphrates 	euphrates.EuphratesCounty
	基尔库克Kirkuk 	kirkuk.KirkukCounty
}

func (d *萨迈拉SamarraDuke) CDeir代尔() deir.DeirCounty {
	return d.代尔Deir
}
    
func (d *萨迈拉SamarraDuke) CEuphrates萨迈拉() euphrates.EuphratesCounty {
	return d.萨迈拉Euphrates
}
    
func (d *萨迈拉SamarraDuke) CKirkuk基尔库克() kirkuk.KirkukCounty {
	return d.基尔库克Kirkuk
}
    
var DSamarra萨迈拉 SamarraDuke = &萨迈拉SamarraDuke{}

func init() {
	f := DSamarra萨迈拉.(*萨迈拉SamarraDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "samarra",
		TitleName: "萨迈拉",
		TitleCode: "d_samarra",
		Counties:  map[string]feud.County{},
	}

	f.代尔Deir = deir.CDeir代尔
	f.代尔Deir.SetParent(f)
	
	f.萨迈拉Euphrates = euphrates.CEuphrates萨迈拉
	f.萨迈拉Euphrates.SetParent(f)
	
	f.基尔库克Kirkuk = kirkuk.CKirkuk基尔库克
	f.基尔库克Kirkuk.SetParent(f)
	
}
