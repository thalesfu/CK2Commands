package orkhon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鄂尔浑图勒OrkhontuulBarony struct {
	feud.BaseBarony
}

var BOrkhontuul鄂尔浑图勒 feud.Barony = &鄂尔浑图勒OrkhontuulBarony{}

func init() {
    f := BOrkhontuul鄂尔浑图勒.(*鄂尔浑图勒OrkhontuulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "orkhontuul",
		TitleName: "鄂尔浑图勒",
		TitleCode: "b_orkhontuul",
	}
}
