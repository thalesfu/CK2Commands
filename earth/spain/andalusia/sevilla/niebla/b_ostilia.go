package niebla

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥斯蒂利亚OstiliaBarony struct {
	feud.BaseBarony
}

var BOstilia奥斯蒂利亚 feud.Barony = &奥斯蒂利亚OstiliaBarony{}

func init() {
	f := BOstilia奥斯蒂利亚.(*奥斯蒂利亚OstiliaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ostilia",
		TitleName: "奥斯蒂利亚",
		TitleCode: "b_ostilia",
	}
}
