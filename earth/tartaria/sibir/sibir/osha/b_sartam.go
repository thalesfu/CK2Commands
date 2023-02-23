package osha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨尔塔姆SartamBarony struct {
	feud.BaseBarony
}

var BSartam萨尔塔姆 feud.Barony = &萨尔塔姆SartamBarony{}

func init() {
	f := BSartam萨尔塔姆.(*萨尔塔姆SartamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sartam",
		TitleName: "萨尔塔姆",
		TitleCode: "b_sartam",
	}
}
