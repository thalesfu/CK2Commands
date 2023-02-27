package cholamandalam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 甘那罗KannanurBarony struct {
	feud.BaseBarony
}

var BKannanur甘那罗 feud.Barony = &甘那罗KannanurBarony{}

func init() {
    f := BKannanur甘那罗.(*甘那罗KannanurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kannanur",
		TitleName: "甘那罗",
		TitleCode: "b_kannanur",
	}
}
