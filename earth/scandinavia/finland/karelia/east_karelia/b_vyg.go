package east_karelia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维格VygBarony struct {
	feud.BaseBarony
}

var BVyg维格 feud.Barony = &维格VygBarony{}

func init() {
    f := BVyg维格.(*维格VygBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vyg",
		TitleName: "维格",
		TitleCode: "b_vyg",
	}
}
