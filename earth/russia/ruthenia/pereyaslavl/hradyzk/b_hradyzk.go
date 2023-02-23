package hradyzk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格拉季济克HradyzkBarony struct {
	feud.BaseBarony
}

var BHradyzk格拉季济克 feud.Barony = &格拉季济克HradyzkBarony{}

func init() {
	f := BHradyzk格拉季济克.(*格拉季济克HradyzkBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hradyzk",
		TitleName: "格拉季济克",
		TitleCode: "b_hradyzk",
	}
}
