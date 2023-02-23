package pannagallu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佛罗BodlaBarony struct {
	feud.BaseBarony
}

var BBodla佛罗 feud.Barony = &佛罗BodlaBarony{}

func init() {
	f := BBodla佛罗.(*佛罗BodlaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bodla",
		TitleName: "佛罗",
		TitleCode: "b_bodla",
	}
}
