package sambhal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多檀陀补罗TattandapuraBarony struct {
	feud.BaseBarony
}

var BTattandapura多檀陀补罗 feud.Barony = &多檀陀补罗TattandapuraBarony{}

func init() {
    f := BTattandapura多檀陀补罗.(*多檀陀补罗TattandapuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tattandapura",
		TitleName: "多檀陀补罗",
		TitleCode: "b_tattandapura",
	}
}
