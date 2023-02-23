package amman

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰拉什JerashBarony struct {
	feud.BaseBarony
}

var BJerash杰拉什 feud.Barony = &杰拉什JerashBarony{}

func init() {
	f := BJerash杰拉什.(*杰拉什JerashBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jerash",
		TitleName: "杰拉什",
		TitleCode: "b_jerash",
	}
}
