package mainz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多恩堡DornburgBarony struct {
	feud.BaseBarony
}

var BDornburg多恩堡 feud.Barony = &多恩堡DornburgBarony{}

func init() {
    f := BDornburg多恩堡.(*多恩堡DornburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dornburg",
		TitleName: "多恩堡",
		TitleCode: "b_dornburg",
	}
}
