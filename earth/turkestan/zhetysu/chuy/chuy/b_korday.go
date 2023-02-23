package chuy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科尔代KordayBarony struct {
	feud.BaseBarony
}

var BKorday科尔代 feud.Barony = &科尔代KordayBarony{}

func init() {
	f := BKorday科尔代.(*科尔代KordayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "korday",
		TitleName: "科尔代",
		TitleCode: "b_korday",
	}
}
