package baalbek

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿卡尔AkkarBarony struct {
	feud.BaseBarony
}

var BAkkar阿卡尔 feud.Barony = &阿卡尔AkkarBarony{}

func init() {
    f := BAkkar阿卡尔.(*阿卡尔AkkarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "akkar",
		TitleName: "阿卡尔",
		TitleCode: "b_akkar",
	}
}
