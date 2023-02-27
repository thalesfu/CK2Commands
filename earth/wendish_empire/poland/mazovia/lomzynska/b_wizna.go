package lomzynska

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维兹纳WiznaBarony struct {
	feud.BaseBarony
}

var BWizna维兹纳 feud.Barony = &维兹纳WiznaBarony{}

func init() {
    f := BWizna维兹纳.(*维兹纳WiznaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wizna",
		TitleName: "维兹纳",
		TitleCode: "b_wizna",
	}
}
