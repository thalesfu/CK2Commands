package jersika

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 内格斯特NegesteBarony struct {
	feud.BaseBarony
}

var BNegeste内格斯特 feud.Barony = &内格斯特NegesteBarony{}

func init() {
    f := BNegeste内格斯特.(*内格斯特NegesteBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "negeste",
		TitleName: "内格斯特",
		TitleCode: "b_negeste",
	}
}
