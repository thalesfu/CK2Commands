package sarqihya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿鲁达AgrudaBarony struct {
	feud.BaseBarony
}

var BAgruda阿鲁达 feud.Barony = &阿鲁达AgrudaBarony{}

func init() {
	f := BAgruda阿鲁达.(*阿鲁达AgrudaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "agruda",
		TitleName: "阿鲁达",
		TitleCode: "b_agruda",
	}
}
