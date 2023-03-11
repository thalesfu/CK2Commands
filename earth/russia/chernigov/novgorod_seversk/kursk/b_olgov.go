package kursk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥尔戈夫OlgovBarony struct {
	feud.BaseBarony
}

var BOlgov奥尔戈夫 feud.Barony = &奥尔戈夫OlgovBarony{}

func init() {
    f := BOlgov奥尔戈夫.(*奥尔戈夫OlgovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "olgov",
		TitleName: "奥尔戈夫",
		TitleCode: "b_olgov",
	}
}
