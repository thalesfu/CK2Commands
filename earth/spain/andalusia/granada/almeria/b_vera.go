package almeria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝拉VeraBarony struct {
	feud.BaseBarony
}

var BVera贝拉 feud.Barony = &贝拉VeraBarony{}

func init() {
	f := BVera贝拉.(*贝拉VeraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vera",
		TitleName: "贝拉",
		TitleCode: "b_vera",
	}
}
