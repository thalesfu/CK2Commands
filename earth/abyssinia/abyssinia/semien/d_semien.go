package semien

import (
	"github.com/thalesfu/CK2Commands/earth/abyssinia/abyssinia/semien/semien"
	"github.com/thalesfu/CK2Commands/earth/abyssinia/abyssinia/semien/tigrinya"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SemienDuke interface {
	feud.Duke
	CSemien塞米恩() semien.SemienCounty
	CTigrinya提格利尼亚() tigrinya.TigrinyaCounty
}

type 塞米恩SemienDuke struct {
	feud.BaseDuke
	塞米恩Semien     semien.SemienCounty
	提格利尼亚Tigrinya tigrinya.TigrinyaCounty
}

func (d *塞米恩SemienDuke) CSemien塞米恩() semien.SemienCounty {
	return d.塞米恩Semien
}

func (d *塞米恩SemienDuke) CTigrinya提格利尼亚() tigrinya.TigrinyaCounty {
	return d.提格利尼亚Tigrinya
}

var DSemien塞米恩 SemienDuke = &塞米恩SemienDuke{}

func init() {
	f := DSemien塞米恩.(*塞米恩SemienDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "semien",
		TitleName: "塞米恩",
		TitleCode: "d_semien",
		Counties:  map[string]feud.County{},
	}

	f.塞米恩Semien = semien.CSemien塞米恩
	f.塞米恩Semien.SetParent(f)

	f.提格利尼亚Tigrinya = tigrinya.CTigrinya提格利尼亚
	f.提格利尼亚Tigrinya.SetParent(f)

}
