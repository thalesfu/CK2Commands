package suzdal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维丘加VichugaBarony struct {
	feud.BaseBarony
}

var BVichuga维丘加 feud.Barony = &维丘加VichugaBarony{}

func init() {
    f := BVichuga维丘加.(*维丘加VichugaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vichuga",
		TitleName: "维丘加",
		TitleCode: "b_vichuga",
	}
}
