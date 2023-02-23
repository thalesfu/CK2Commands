package vijayawada

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩苏拉MasulaBarony struct {
	feud.BaseBarony
}

var BMasula摩苏拉 feud.Barony = &摩苏拉MasulaBarony{}

func init() {
	f := BMasula摩苏拉.(*摩苏拉MasulaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "masula",
		TitleName: "摩苏拉",
		TitleCode: "b_masula",
	}
}
