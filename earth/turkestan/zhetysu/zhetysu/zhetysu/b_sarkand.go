package zhetysu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨尔坎德SarkandBarony struct {
	feud.BaseBarony
}

var BSarkand萨尔坎德 feud.Barony = &萨尔坎德SarkandBarony{}

func init() {
    f := BSarkand萨尔坎德.(*萨尔坎德SarkandBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sarkand",
		TitleName: "萨尔坎德",
		TitleCode: "b_sarkand",
	}
}
