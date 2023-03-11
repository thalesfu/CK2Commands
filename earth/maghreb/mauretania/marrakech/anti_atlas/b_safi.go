package anti_atlas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨非SafiBarony struct {
	feud.BaseBarony
}

var BSafi萨非 feud.Barony = &萨非SafiBarony{}

func init() {
    f := BSafi萨非.(*萨非SafiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "safi",
		TitleName: "萨非",
		TitleCode: "b_safi",
	}
}
