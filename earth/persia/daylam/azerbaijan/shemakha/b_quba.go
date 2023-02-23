package shemakha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库巴QubaBarony struct {
	feud.BaseBarony
}

var BQuba库巴 feud.Barony = &库巴QubaBarony{}

func init() {
	f := BQuba库巴.(*库巴QubaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "quba",
		TitleName: "库巴",
		TitleCode: "b_quba",
	}
}
