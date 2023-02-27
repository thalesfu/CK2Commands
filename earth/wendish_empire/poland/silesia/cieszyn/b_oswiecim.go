package cieszyn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥斯威辛OswiecimBarony struct {
	feud.BaseBarony
}

var BOswiecim奥斯威辛 feud.Barony = &奥斯威辛OswiecimBarony{}

func init() {
    f := BOswiecim奥斯威辛.(*奥斯威辛OswiecimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oswiecim",
		TitleName: "奥斯威辛",
		TitleCode: "b_oswiecim",
	}
}
