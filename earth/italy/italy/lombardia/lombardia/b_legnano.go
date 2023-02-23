package lombardia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱尼亚诺LegnanoBarony struct {
	feud.BaseBarony
}

var BLegnano莱尼亚诺 feud.Barony = &莱尼亚诺LegnanoBarony{}

func init() {
	f := BLegnano莱尼亚诺.(*莱尼亚诺LegnanoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "legnano",
		TitleName: "莱尼亚诺",
		TitleCode: "b_legnano",
	}
}
