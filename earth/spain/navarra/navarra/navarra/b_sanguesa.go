package navarra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 桑圭萨SanguesaBarony struct {
	feud.BaseBarony
}

var BSanguesa桑圭萨 feud.Barony = &桑圭萨SanguesaBarony{}

func init() {
    f := BSanguesa桑圭萨.(*桑圭萨SanguesaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sanguesa",
		TitleName: "桑圭萨",
		TitleCode: "b_sanguesa",
	}
}
