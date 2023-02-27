package baghdad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希拉HillahBarony struct {
	feud.BaseBarony
}

var BHillah希拉 feud.Barony = &希拉HillahBarony{}

func init() {
    f := BHillah希拉.(*希拉HillahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hillah",
		TitleName: "希拉",
		TitleCode: "b_hillah",
	}
}
