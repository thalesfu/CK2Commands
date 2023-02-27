package phiti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曼那罗MannaraBarony struct {
	feud.BaseBarony
}

var BMannara曼那罗 feud.Barony = &曼那罗MannaraBarony{}

func init() {
    f := BMannara曼那罗.(*曼那罗MannaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mannara",
		TitleName: "曼那罗",
		TitleCode: "b_mannara",
	}
}
