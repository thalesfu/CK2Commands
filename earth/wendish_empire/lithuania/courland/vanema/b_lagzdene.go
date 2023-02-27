package vanema

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉格兹德内LagzdeneBarony struct {
	feud.BaseBarony
}

var BLagzdene拉格兹德内 feud.Barony = &拉格兹德内LagzdeneBarony{}

func init() {
    f := BLagzdene拉格兹德内.(*拉格兹德内LagzdeneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lagzdene",
		TitleName: "拉格兹德内",
		TitleCode: "b_lagzdene",
	}
}
