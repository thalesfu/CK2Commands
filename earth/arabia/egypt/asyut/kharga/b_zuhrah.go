package kharga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 祖赫拉ZuhrahBarony struct {
	feud.BaseBarony
}

var BZuhrah祖赫拉 feud.Barony = &祖赫拉ZuhrahBarony{}

func init() {
    f := BZuhrah祖赫拉.(*祖赫拉ZuhrahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zuhrah",
		TitleName: "祖赫拉",
		TitleCode: "b_zuhrah",
	}
}
