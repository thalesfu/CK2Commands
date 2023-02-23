package verden

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥胡森OuhusenBarony struct {
	feud.BaseBarony
}

var BOuhusen奥胡森 feud.Barony = &奥胡森OuhusenBarony{}

func init() {
	f := BOuhusen奥胡森.(*奥胡森OuhusenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ouhusen",
		TitleName: "奥胡森",
		TitleCode: "b_ouhusen",
	}
}
