package ivrea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比耶拉BiellaBarony struct {
	feud.BaseBarony
}

var BBiella比耶拉 feud.Barony = &比耶拉BiellaBarony{}

func init() {
	f := BBiella比耶拉.(*比耶拉BiellaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "biella",
		TitleName: "比耶拉",
		TitleCode: "b_biella",
	}
}
