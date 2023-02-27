package delgerkhangai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 昌特TsantBarony struct {
	feud.BaseBarony
}

var BTsant昌特 feud.Barony = &昌特TsantBarony{}

func init() {
    f := BTsant昌特.(*昌特TsantBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tsant",
		TitleName: "昌特",
		TitleCode: "b_tsant",
	}
}
