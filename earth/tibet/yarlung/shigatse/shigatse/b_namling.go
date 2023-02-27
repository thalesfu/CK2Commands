package shigatse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 南木林NamlingBarony struct {
	feud.BaseBarony
}

var BNamling南木林 feud.Barony = &南木林NamlingBarony{}

func init() {
    f := BNamling南木林.(*南木林NamlingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "namling",
		TitleName: "南木林",
		TitleCode: "b_namling",
	}
}
