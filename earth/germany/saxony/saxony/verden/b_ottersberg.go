package verden

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥特斯贝格OttersbergBarony struct {
	feud.BaseBarony
}

var BOttersberg奥特斯贝格 feud.Barony = &奥特斯贝格OttersbergBarony{}

func init() {
	f := BOttersberg奥特斯贝格.(*奥特斯贝格OttersbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ottersberg",
		TitleName: "奥特斯贝格",
		TitleCode: "b_ottersberg",
	}
}
