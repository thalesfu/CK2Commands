package durham

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈特尔浦HartlepoolBarony struct {
	feud.BaseBarony
}

var BHartlepool哈特尔浦 feud.Barony = &哈特尔浦HartlepoolBarony{}

func init() {
    f := BHartlepool哈特尔浦.(*哈特尔浦HartlepoolBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hartlepool",
		TitleName: "哈特尔浦",
		TitleCode: "b_hartlepool",
	}
}
