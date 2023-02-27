package lothian

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 爱丁堡EdinburghBarony struct {
	feud.BaseBarony
}

var BEdinburgh爱丁堡 feud.Barony = &爱丁堡EdinburghBarony{}

func init() {
    f := BEdinburgh爱丁堡.(*爱丁堡EdinburghBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "edinburgh",
		TitleName: "爱丁堡",
		TitleCode: "b_edinburgh",
	}
}
