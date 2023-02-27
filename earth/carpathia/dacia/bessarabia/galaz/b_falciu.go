package galaz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗尔丘FalciuBarony struct {
	feud.BaseBarony
}

var BFalciu弗尔丘 feud.Barony = &弗尔丘FalciuBarony{}

func init() {
    f := BFalciu弗尔丘.(*弗尔丘FalciuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "falciu",
		TitleName: "弗尔丘",
		TitleCode: "b_falciu",
	}
}
