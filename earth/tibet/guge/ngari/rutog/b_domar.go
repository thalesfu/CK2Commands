package rutog

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多玛DomarBarony struct {
	feud.BaseBarony
}

var BDomar多玛 feud.Barony = &多玛DomarBarony{}

func init() {
    f := BDomar多玛.(*多玛DomarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "domar",
		TitleName: "多玛",
		TitleCode: "b_domar",
	}
}
