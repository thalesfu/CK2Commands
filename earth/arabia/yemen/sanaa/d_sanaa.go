package sanaa

import (
	"github.com/thalesfu/CK2Commands/earth/arabia/yemen/sanaa/najran"
	"github.com/thalesfu/CK2Commands/earth/arabia/yemen/sanaa/sanaa"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SanaaDuke interface {
	feud.Duke
	CNajran纳季兰() najran.NajranCounty
	CSanaa萨那() sanaa.SanaaCounty
}

type 萨那SanaaDuke struct {
	feud.BaseDuke
	纳季兰Najran najran.NajranCounty
	萨那Sanaa   sanaa.SanaaCounty
}

func (d *萨那SanaaDuke) CNajran纳季兰() najran.NajranCounty {
	return d.纳季兰Najran
}

func (d *萨那SanaaDuke) CSanaa萨那() sanaa.SanaaCounty {
	return d.萨那Sanaa
}

var DSanaa萨那 SanaaDuke = &萨那SanaaDuke{}

func init() {
	f := DSanaa萨那.(*萨那SanaaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "sanaa",
		TitleName: "萨那",
		TitleCode: "d_sanaa",
		Counties:  map[string]feud.County{},
	}

	f.纳季兰Najran = najran.CNajran纳季兰
	f.纳季兰Najran.SetParent(f)

	f.萨那Sanaa = sanaa.CSanaa萨那
	f.萨那Sanaa.SetParent(f)

}
