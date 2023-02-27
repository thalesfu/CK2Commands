package bambuk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 桑卡朗SankaranBarony struct {
	feud.BaseBarony
}

var BSankaran桑卡朗 feud.Barony = &桑卡朗SankaranBarony{}

func init() {
    f := BSankaran桑卡朗.(*桑卡朗SankaranBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sankaran",
		TitleName: "桑卡朗",
		TitleCode: "b_sankaran",
	}
}
