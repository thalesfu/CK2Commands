package gyesar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 桑木拉SemolaBarony struct {
	feud.BaseBarony
}

var BSemola桑木拉 feud.Barony = &桑木拉SemolaBarony{}

func init() {
	f := BSemola桑木拉.(*桑木拉SemolaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "semola",
		TitleName: "桑木拉",
		TitleCode: "b_semola",
	}
}
