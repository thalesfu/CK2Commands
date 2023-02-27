package hastinapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诃补罗HapurBarony struct {
	feud.BaseBarony
}

var BHapur诃补罗 feud.Barony = &诃补罗HapurBarony{}

func init() {
    f := BHapur诃补罗.(*诃补罗HapurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hapur",
		TitleName: "诃补罗",
		TitleCode: "b_hapur",
	}
}
