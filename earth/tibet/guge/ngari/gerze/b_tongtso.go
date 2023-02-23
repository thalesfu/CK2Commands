package gerze

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洞措TongtsoBarony struct {
	feud.BaseBarony
}

var BTongtso洞措 feud.Barony = &洞措TongtsoBarony{}

func init() {
	f := BTongtso洞措.(*洞措TongtsoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tongtso",
		TitleName: "洞措",
		TitleCode: "b_tongtso",
	}
}
