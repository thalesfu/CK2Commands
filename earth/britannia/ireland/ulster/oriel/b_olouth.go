package oriel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 劳斯OlouthBarony struct {
	feud.BaseBarony
}

var BOlouth劳斯 feud.Barony = &劳斯OlouthBarony{}

func init() {
	f := BOlouth劳斯.(*劳斯OlouthBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "olouth",
		TitleName: "劳斯",
		TitleCode: "b_olouth",
	}
}
