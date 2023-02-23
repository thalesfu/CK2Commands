package racakonda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 娑婆补罗SobhapurBarony struct {
	feud.BaseBarony
}

var BSobhapur娑婆补罗 feud.Barony = &娑婆补罗SobhapurBarony{}

func init() {
	f := BSobhapur娑婆补罗.(*娑婆补罗SobhapurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sobhapur",
		TitleName: "娑婆补罗",
		TitleCode: "b_sobhapur",
	}
}
