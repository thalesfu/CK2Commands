package jersika

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥季内AutineBarony struct {
	feud.BaseBarony
}

var BAutine奥季内 feud.Barony = &奥季内AutineBarony{}

func init() {
    f := BAutine奥季内.(*奥季内AutineBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "autine",
		TitleName: "奥季内",
		TitleCode: "b_autine",
	}
}
