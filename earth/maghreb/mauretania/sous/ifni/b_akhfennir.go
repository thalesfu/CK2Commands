package ifni

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿克费尼AkhfennirBarony struct {
	feud.BaseBarony
}

var BAkhfennir阿克费尼 feud.Barony = &阿克费尼AkhfennirBarony{}

func init() {
    f := BAkhfennir阿克费尼.(*阿克费尼AkhfennirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "akhfennir",
		TitleName: "阿克费尼",
		TitleCode: "b_akhfennir",
	}
}
