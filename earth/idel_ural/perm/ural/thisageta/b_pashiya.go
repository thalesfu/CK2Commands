package thisageta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕希亚PashiyaBarony struct {
	feud.BaseBarony
}

var BPashiya帕希亚 feud.Barony = &帕希亚PashiyaBarony{}

func init() {
    f := BPashiya帕希亚.(*帕希亚PashiyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pashiya",
		TitleName: "帕希亚",
		TitleCode: "b_pashiya",
	}
}
