package delgerkhangai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赛罕SaikhanBarony struct {
	feud.BaseBarony
}

var BSaikhan赛罕 feud.Barony = &赛罕SaikhanBarony{}

func init() {
    f := BSaikhan赛罕.(*赛罕SaikhanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saikhan",
		TitleName: "赛罕",
		TitleCode: "b_saikhan",
	}
}
