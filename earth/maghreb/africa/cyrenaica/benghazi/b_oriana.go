package benghazi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥利安OrianaBarony struct {
	feud.BaseBarony
}

var BOriana奥利安 feud.Barony = &奥利安OrianaBarony{}

func init() {
    f := BOriana奥利安.(*奥利安OrianaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oriana",
		TitleName: "奥利安",
		TitleCode: "b_oriana",
	}
}
