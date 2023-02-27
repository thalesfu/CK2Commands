package lower_don

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马特韦耶夫_库尔干MatveevkurganBarony struct {
	feud.BaseBarony
}

var BMatveevkurgan马特韦耶夫_库尔干 feud.Barony = &马特韦耶夫_库尔干MatveevkurganBarony{}

func init() {
    f := BMatveevkurgan马特韦耶夫_库尔干.(*马特韦耶夫_库尔干MatveevkurganBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "matveevkurgan",
		TitleName: "马特韦耶夫_库尔干",
		TitleCode: "b_matveevkurgan",
	}
}
