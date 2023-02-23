package samtho

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 错仁TsoringBarony struct {
	feud.BaseBarony
}

var BTsoring错仁 feud.Barony = &错仁TsoringBarony{}

func init() {
	f := BTsoring错仁.(*错仁TsoringBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tsoring",
		TitleName: "错仁",
		TitleCode: "b_tsoring",
	}
}
