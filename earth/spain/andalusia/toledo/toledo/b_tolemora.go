package toledo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫拉TolemoraBarony struct {
	feud.BaseBarony
}

var BTolemora莫拉 feud.Barony = &莫拉TolemoraBarony{}

func init() {
    f := BTolemora莫拉.(*莫拉TolemoraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tolemora",
		TitleName: "莫拉",
		TitleCode: "b_tolemora",
	}
}
