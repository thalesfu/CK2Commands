package tana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏林SulinBarony struct {
	feud.BaseBarony
}

var BSulin苏林 feud.Barony = &苏林SulinBarony{}

func init() {
    f := BSulin苏林.(*苏林SulinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sulin",
		TitleName: "苏林",
		TitleCode: "b_sulin",
	}
}
