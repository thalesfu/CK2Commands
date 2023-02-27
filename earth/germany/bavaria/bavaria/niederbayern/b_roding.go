package niederbayern

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗丁RodingBarony struct {
	feud.BaseBarony
}

var BRoding罗丁 feud.Barony = &罗丁RodingBarony{}

func init() {
    f := BRoding罗丁.(*罗丁RodingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "roding",
		TitleName: "罗丁",
		TitleCode: "b_roding",
	}
}
