package tahoua

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿戈阿德瓦勒Ago_adewaleBarony struct {
	feud.BaseBarony
}

var BAgo_adewale阿戈阿德瓦勒 feud.Barony = &阿戈阿德瓦勒Ago_adewaleBarony{}

func init() {
    f := BAgo_adewale阿戈阿德瓦勒.(*阿戈阿德瓦勒Ago_adewaleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ago_adewale",
		TitleName: "阿戈阿德瓦勒",
		TitleCode: "b_ago_adewale",
	}
}
