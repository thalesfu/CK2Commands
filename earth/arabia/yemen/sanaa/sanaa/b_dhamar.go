package sanaa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎马尔DhamarBarony struct {
	feud.BaseBarony
}

var BDhamar扎马尔 feud.Barony = &扎马尔DhamarBarony{}

func init() {
	f := BDhamar扎马尔.(*扎马尔DhamarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dhamar",
		TitleName: "扎马尔",
		TitleCode: "b_dhamar",
	}
}
