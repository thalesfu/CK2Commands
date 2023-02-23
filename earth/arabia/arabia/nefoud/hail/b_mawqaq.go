package hail

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毛盖格MawqaqBarony struct {
	feud.BaseBarony
}

var BMawqaq毛盖格 feud.Barony = &毛盖格MawqaqBarony{}

func init() {
	f := BMawqaq毛盖格.(*毛盖格MawqaqBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mawqaq",
		TitleName: "毛盖格",
		TitleCode: "b_mawqaq",
	}
}
