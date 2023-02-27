package neumark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 桑托克SantokBarony struct {
	feud.BaseBarony
}

var BSantok桑托克 feud.Barony = &桑托克SantokBarony{}

func init() {
    f := BSantok桑托克.(*桑托克SantokBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "santok",
		TitleName: "桑托克",
		TitleCode: "b_santok",
	}
}
