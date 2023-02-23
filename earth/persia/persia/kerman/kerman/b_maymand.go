package kerman

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迈曼德MaymandBarony struct {
	feud.BaseBarony
}

var BMaymand迈曼德 feud.Barony = &迈曼德MaymandBarony{}

func init() {
	f := BMaymand迈曼德.(*迈曼德MaymandBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maymand",
		TitleName: "迈曼德",
		TitleCode: "b_maymand",
	}
}
