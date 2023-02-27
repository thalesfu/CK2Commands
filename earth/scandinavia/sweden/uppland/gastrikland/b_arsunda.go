package gastrikland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥松达ArsundaBarony struct {
	feud.BaseBarony
}

var BArsunda奥松达 feud.Barony = &奥松达ArsundaBarony{}

func init() {
    f := BArsunda奥松达.(*奥松达ArsundaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arsunda",
		TitleName: "奥松达",
		TitleCode: "b_arsunda",
	}
}
