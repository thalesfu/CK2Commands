package mansura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曼苏拉MansuraBarony struct {
	feud.BaseBarony
}

var BMansura曼苏拉 feud.Barony = &曼苏拉MansuraBarony{}

func init() {
	f := BMansura曼苏拉.(*曼苏拉MansuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mansura",
		TitleName: "曼苏拉",
		TitleCode: "b_mansura",
	}
}
