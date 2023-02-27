package monferrato

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡萨莱CasaleBarony struct {
	feud.BaseBarony
}

var BCasale卡萨莱 feud.Barony = &卡萨莱CasaleBarony{}

func init() {
    f := BCasale卡萨莱.(*卡萨莱CasaleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "casale",
		TitleName: "卡萨莱",
		TitleCode: "b_casale",
	}
}
