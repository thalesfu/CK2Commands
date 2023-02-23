package ferghana

import (
	"github.com/thalesfu/CK2Commands/earth/turkestan/khiva/ferghana/fergana"
	"github.com/thalesfu/CK2Commands/earth/turkestan/khiva/ferghana/khaylam"
	"github.com/thalesfu/CK2Commands/earth/turkestan/khiva/ferghana/khojand"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FerghanaDuke interface {
	feud.Duke
	CFergana费尔干纳() fergana.FerganaCounty
	CKhaylam海拉姆() khaylam.KhaylamCounty
	CKhojand苦盏() khojand.KhojandCounty
}

type 费尔干纳FerghanaDuke struct {
	feud.BaseDuke
	费尔干纳Fergana fergana.FerganaCounty
	海拉姆Khaylam  khaylam.KhaylamCounty
	苦盏Khojand   khojand.KhojandCounty
}

func (d *费尔干纳FerghanaDuke) CFergana费尔干纳() fergana.FerganaCounty {
	return d.费尔干纳Fergana
}

func (d *费尔干纳FerghanaDuke) CKhaylam海拉姆() khaylam.KhaylamCounty {
	return d.海拉姆Khaylam
}

func (d *费尔干纳FerghanaDuke) CKhojand苦盏() khojand.KhojandCounty {
	return d.苦盏Khojand
}

var DFerghana费尔干纳 FerghanaDuke = &费尔干纳FerghanaDuke{}

func init() {
	f := DFerghana费尔干纳.(*费尔干纳FerghanaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "ferghana",
		TitleName: "费尔干纳",
		TitleCode: "d_ferghana",
		Counties:  map[string]feud.County{},
	}

	f.费尔干纳Fergana = fergana.CFergana费尔干纳
	f.费尔干纳Fergana.SetParent(f)

	f.海拉姆Khaylam = khaylam.CKhaylam海拉姆
	f.海拉姆Khaylam.SetParent(f)

	f.苦盏Khojand = khojand.CKhojand苦盏
	f.苦盏Khojand.SetParent(f)

}
