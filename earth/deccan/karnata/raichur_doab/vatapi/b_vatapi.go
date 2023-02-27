package vatapi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伐达比VatapiBarony struct {
	feud.BaseBarony
}

var BVatapi伐达比 feud.Barony = &伐达比VatapiBarony{}

func init() {
    f := BVatapi伐达比.(*伐达比VatapiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vatapi",
		TitleName: "伐达比",
		TitleCode: "b_vatapi",
	}
}
