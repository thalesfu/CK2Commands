package vidin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维丁VidinBarony struct {
	feud.BaseBarony
}

var BVidin维丁 feud.Barony = &维丁VidinBarony{}

func init() {
    f := BVidin维丁.(*维丁VidinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vidin",
		TitleName: "维丁",
		TitleCode: "b_vidin",
	}
}
