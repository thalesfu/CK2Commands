package mahoba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 旬曼补JhammanpurBarony struct {
	feud.BaseBarony
}

var BJhammanpur旬曼补 feud.Barony = &旬曼补JhammanpurBarony{}

func init() {
    f := BJhammanpur旬曼补.(*旬曼补JhammanpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jhammanpur",
		TitleName: "旬曼补",
		TitleCode: "b_jhammanpur",
	}
}
