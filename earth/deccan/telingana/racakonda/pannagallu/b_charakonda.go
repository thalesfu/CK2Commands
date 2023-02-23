package pannagallu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 车罗军荼CharakondaBarony struct {
	feud.BaseBarony
}

var BCharakonda车罗军荼 feud.Barony = &车罗军荼CharakondaBarony{}

func init() {
	f := BCharakonda车罗军荼.(*车罗军荼CharakondaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "charakonda",
		TitleName: "车罗军荼",
		TitleCode: "b_charakonda",
	}
}
