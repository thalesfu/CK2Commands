package somerset

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 约维尔YeovilBarony struct {
	feud.BaseBarony
}

var BYeovil约维尔 feud.Barony = &约维尔YeovilBarony{}

func init() {
    f := BYeovil约维尔.(*约维尔YeovilBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yeovil",
		TitleName: "约维尔",
		TitleCode: "b_yeovil",
	}
}
