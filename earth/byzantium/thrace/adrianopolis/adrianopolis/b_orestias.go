package adrianopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥雷斯蒂阿斯OrestiasBarony struct {
	feud.BaseBarony
}

var BOrestias奥雷斯蒂阿斯 feud.Barony = &奥雷斯蒂阿斯OrestiasBarony{}

func init() {
    f := BOrestias奥雷斯蒂阿斯.(*奥雷斯蒂阿斯OrestiasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "orestias",
		TitleName: "奥雷斯蒂阿斯",
		TitleCode: "b_orestias",
	}
}
