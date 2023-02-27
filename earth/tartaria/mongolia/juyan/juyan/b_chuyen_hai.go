package juyan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 居延海Chuyen_haiBarony struct {
	feud.BaseBarony
}

var BChuyen_hai居延海 feud.Barony = &居延海Chuyen_haiBarony{}

func init() {
    f := BChuyen_hai居延海.(*居延海Chuyen_haiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chuyen_hai",
		TitleName: "居延海",
		TitleCode: "b_chuyen_hai",
	}
}
