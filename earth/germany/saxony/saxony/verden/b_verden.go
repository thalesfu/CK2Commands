package verden

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 费尔登VerdenBarony struct {
	feud.BaseBarony
}

var BVerden费尔登 feud.Barony = &费尔登VerdenBarony{}

func init() {
    f := BVerden费尔登.(*费尔登VerdenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "verden",
		TitleName: "费尔登",
		TitleCode: "b_verden",
	}
}
