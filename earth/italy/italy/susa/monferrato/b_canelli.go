package monferrato

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡内利CanelliBarony struct {
	feud.BaseBarony
}

var BCanelli卡内利 feud.Barony = &卡内利CanelliBarony{}

func init() {
	f := BCanelli卡内利.(*卡内利CanelliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "canelli",
		TitleName: "卡内利",
		TitleCode: "b_canelli",
	}
}
