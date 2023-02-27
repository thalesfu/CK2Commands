package tuul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 独乐TuulBarony struct {
	feud.BaseBarony
}

var BTuul独乐 feud.Barony = &独乐TuulBarony{}

func init() {
    f := BTuul独乐.(*独乐TuulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tuul",
		TitleName: "独乐",
		TitleCode: "b_tuul",
	}
}
