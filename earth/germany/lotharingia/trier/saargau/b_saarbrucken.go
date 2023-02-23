package saargau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨尔布吕肯SaarbruckenBarony struct {
	feud.BaseBarony
}

var BSaarbrucken萨尔布吕肯 feud.Barony = &萨尔布吕肯SaarbruckenBarony{}

func init() {
	f := BSaarbrucken萨尔布吕肯.(*萨尔布吕肯SaarbruckenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saarbrucken",
		TitleName: "萨尔布吕肯",
		TitleCode: "b_saarbrucken",
	}
}
