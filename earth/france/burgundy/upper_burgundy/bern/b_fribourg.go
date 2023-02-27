package bern

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗里堡FribourgBarony struct {
	feud.BaseBarony
}

var BFribourg弗里堡 feud.Barony = &弗里堡FribourgBarony{}

func init() {
    f := BFribourg弗里堡.(*弗里堡FribourgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fribourg",
		TitleName: "弗里堡",
		TitleCode: "b_fribourg",
	}
}
