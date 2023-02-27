package csanad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 琼格拉德CsongradBarony struct {
	feud.BaseBarony
}

var BCsongrad琼格拉德 feud.Barony = &琼格拉德CsongradBarony{}

func init() {
    f := BCsongrad琼格拉德.(*琼格拉德CsongradBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "csongrad",
		TitleName: "琼格拉德",
		TitleCode: "b_csongrad",
	}
}
