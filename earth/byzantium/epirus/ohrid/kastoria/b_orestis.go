package kastoria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥雷斯蒂斯OrestisBarony struct {
	feud.BaseBarony
}

var BOrestis奥雷斯蒂斯 feud.Barony = &奥雷斯蒂斯OrestisBarony{}

func init() {
	f := BOrestis奥雷斯蒂斯.(*奥雷斯蒂斯OrestisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "orestis",
		TitleName: "奥雷斯蒂斯",
		TitleCode: "b_orestis",
	}
}
