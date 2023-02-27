package rayy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲁代亨RoudehenBarony struct {
	feud.BaseBarony
}

var BRoudehen鲁代亨 feud.Barony = &鲁代亨RoudehenBarony{}

func init() {
    f := BRoudehen鲁代亨.(*鲁代亨RoudehenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "roudehen",
		TitleName: "鲁代亨",
		TitleCode: "b_roudehen",
	}
}
