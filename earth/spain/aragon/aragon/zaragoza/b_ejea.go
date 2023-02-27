package zaragoza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃赫阿EjeaBarony struct {
	feud.BaseBarony
}

var BEjea埃赫阿 feud.Barony = &埃赫阿EjeaBarony{}

func init() {
    f := BEjea埃赫阿.(*埃赫阿EjeaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ejea",
		TitleName: "埃赫阿",
		TitleCode: "b_ejea",
	}
}
