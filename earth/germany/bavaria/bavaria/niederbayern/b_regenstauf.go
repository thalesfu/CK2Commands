package niederbayern

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷根施陶夫RegenstaufBarony struct {
	feud.BaseBarony
}

var BRegenstauf雷根施陶夫 feud.Barony = &雷根施陶夫RegenstaufBarony{}

func init() {
	f := BRegenstauf雷根施陶夫.(*雷根施陶夫RegenstaufBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "regenstauf",
		TitleName: "雷根施陶夫",
		TitleCode: "b_regenstauf",
	}
}
