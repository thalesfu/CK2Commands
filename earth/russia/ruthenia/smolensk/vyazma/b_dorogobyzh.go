package vyazma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多罗戈布日DorogobyzhBarony struct {
	feud.BaseBarony
}

var BDorogobyzh多罗戈布日 feud.Barony = &多罗戈布日DorogobyzhBarony{}

func init() {
    f := BDorogobyzh多罗戈布日.(*多罗戈布日DorogobyzhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dorogobyzh",
		TitleName: "多罗戈布日",
		TitleCode: "b_dorogobyzh",
	}
}
