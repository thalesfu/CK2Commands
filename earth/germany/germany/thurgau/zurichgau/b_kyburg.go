package zurichgau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基堡KyburgBarony struct {
	feud.BaseBarony
}

var BKyburg基堡 feud.Barony = &基堡KyburgBarony{}

func init() {
    f := BKyburg基堡.(*基堡KyburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kyburg",
		TitleName: "基堡",
		TitleCode: "b_kyburg",
	}
}
