package gastrikland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 海德松达HedesundaBarony struct {
	feud.BaseBarony
}

var BHedesunda海德松达 feud.Barony = &海德松达HedesundaBarony{}

func init() {
    f := BHedesunda海德松达.(*海德松达HedesundaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hedesunda",
		TitleName: "海德松达",
		TitleCode: "b_hedesunda",
	}
}
