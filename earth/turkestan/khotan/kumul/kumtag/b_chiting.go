package kumtag

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赤亭ChitingBarony struct {
	feud.BaseBarony
}

var BChiting赤亭 feud.Barony = &赤亭ChitingBarony{}

func init() {
    f := BChiting赤亭.(*赤亭ChitingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chiting",
		TitleName: "赤亭",
		TitleCode: "b_chiting",
	}
}
