package gotland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维斯比VisbyBarony struct {
	feud.BaseBarony
}

var BVisby维斯比 feud.Barony = &维斯比VisbyBarony{}

func init() {
    f := BVisby维斯比.(*维斯比VisbyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "visby",
		TitleName: "维斯比",
		TitleCode: "b_visby",
	}
}
