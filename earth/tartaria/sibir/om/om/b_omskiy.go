package om

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鄂木斯基OmskiyBarony struct {
	feud.BaseBarony
}

var BOmskiy鄂木斯基 feud.Barony = &鄂木斯基OmskiyBarony{}

func init() {
	f := BOmskiy鄂木斯基.(*鄂木斯基OmskiyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "omskiy",
		TitleName: "鄂木斯基",
		TitleCode: "b_omskiy",
	}
}
