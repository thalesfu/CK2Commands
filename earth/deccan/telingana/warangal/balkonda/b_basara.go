package balkonda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴萨拉BasaraBarony struct {
	feud.BaseBarony
}

var BBasara巴萨拉 feud.Barony = &巴萨拉BasaraBarony{}

func init() {
	f := BBasara巴萨拉.(*巴萨拉BasaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "basara",
		TitleName: "巴萨拉",
		TitleCode: "b_basara",
	}
}
