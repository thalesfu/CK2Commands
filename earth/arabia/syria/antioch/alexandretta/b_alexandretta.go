package alexandretta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚历山大勒塔AlexandrettaBarony struct {
	feud.BaseBarony
}

var BAlexandretta亚历山大勒塔 feud.Barony = &亚历山大勒塔AlexandrettaBarony{}

func init() {
    f := BAlexandretta亚历山大勒塔.(*亚历山大勒塔AlexandrettaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alexandretta",
		TitleName: "亚历山大勒塔",
		TitleCode: "b_alexandretta",
	}
}
