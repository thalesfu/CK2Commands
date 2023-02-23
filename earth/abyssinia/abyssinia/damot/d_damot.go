package damot

import (
	"github.com/thalesfu/CK2Commands/earth/abyssinia/abyssinia/damot/asosa"
	"github.com/thalesfu/CK2Commands/earth/abyssinia/abyssinia/damot/damot"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DamotDuke interface {
	feud.Duke
	CAsosa阿索萨() asosa.AsosaCounty
	CDamot达莫特() damot.DamotCounty
}

type 达莫特DamotDuke struct {
	feud.BaseDuke
	阿索萨Asosa asosa.AsosaCounty
	达莫特Damot damot.DamotCounty
}

func (d *达莫特DamotDuke) CAsosa阿索萨() asosa.AsosaCounty {
	return d.阿索萨Asosa
}

func (d *达莫特DamotDuke) CDamot达莫特() damot.DamotCounty {
	return d.达莫特Damot
}

var DDamot达莫特 DamotDuke = &达莫特DamotDuke{}

func init() {
	f := DDamot达莫特.(*达莫特DamotDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "damot",
		TitleName: "达莫特",
		TitleCode: "d_damot",
		Counties:  map[string]feud.County{},
	}

	f.阿索萨Asosa = asosa.CAsosa阿索萨
	f.阿索萨Asosa.SetParent(f)

	f.达莫特Damot = damot.CDamot达莫特
	f.达莫特Damot.SetParent(f)

}
