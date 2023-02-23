package loulan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鄯善ShanshanBarony struct {
	feud.BaseBarony
}

var BShanshan鄯善 feud.Barony = &鄯善ShanshanBarony{}

func init() {
	f := BShanshan鄯善.(*鄯善ShanshanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shanshan",
		TitleName: "鄯善",
		TitleCode: "b_shanshan",
	}
}
