package urgell

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏埃尔特SuertBarony struct {
	feud.BaseBarony
}

var BSuert苏埃尔特 feud.Barony = &苏埃尔特SuertBarony{}

func init() {
    f := BSuert苏埃尔特.(*苏埃尔特SuertBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "suert",
		TitleName: "苏埃尔特",
		TitleCode: "b_suert",
	}
}
