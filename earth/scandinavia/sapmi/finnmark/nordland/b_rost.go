package nordland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 勒斯特RostBarony struct {
	feud.BaseBarony
}

var BRost勒斯特 feud.Barony = &勒斯特RostBarony{}

func init() {
    f := BRost勒斯特.(*勒斯特RostBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rost",
		TitleName: "勒斯特",
		TitleCode: "b_rost",
	}
}
