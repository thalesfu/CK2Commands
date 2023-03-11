package beresty

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 别尔斯克BielskBarony struct {
	feud.BaseBarony
}

var BBielsk别尔斯克 feud.Barony = &别尔斯克BielskBarony{}

func init() {
    f := BBielsk别尔斯克.(*别尔斯克BielskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bielsk",
		TitleName: "别尔斯克",
		TitleCode: "b_bielsk",
	}
}
