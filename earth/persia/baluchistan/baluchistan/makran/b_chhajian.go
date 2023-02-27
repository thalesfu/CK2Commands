package makran

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 查绛ChhajianBarony struct {
	feud.BaseBarony
}

var BChhajian查绛 feud.Barony = &查绛ChhajianBarony{}

func init() {
    f := BChhajian查绛.(*查绛ChhajianBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chhajian",
		TitleName: "查绛",
		TitleCode: "b_chhajian",
	}
}
