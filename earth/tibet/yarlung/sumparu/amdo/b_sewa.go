package amdo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 色务SewaBarony struct {
	feud.BaseBarony
}

var BSewa色务 feud.Barony = &色务SewaBarony{}

func init() {
	f := BSewa色务.(*色务SewaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sewa",
		TitleName: "色务",
		TitleCode: "b_sewa",
	}
}
