package kutch

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普杰BhujBarony struct {
	feud.BaseBarony
}

var BBhuj普杰 feud.Barony = &普杰BhujBarony{}

func init() {
    f := BBhuj普杰.(*普杰BhujBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhuj",
		TitleName: "普杰",
		TitleCode: "b_bhuj",
	}
}
