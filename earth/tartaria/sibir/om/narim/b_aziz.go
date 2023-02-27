package narim

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿齐兹AzizBarony struct {
	feud.BaseBarony
}

var BAziz阿齐兹 feud.Barony = &阿齐兹AzizBarony{}

func init() {
    f := BAziz阿齐兹.(*阿齐兹AzizBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aziz",
		TitleName: "阿齐兹",
		TitleCode: "b_aziz",
	}
}
