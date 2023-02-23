package guines

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格拉沃利讷GravelinesBarony struct {
	feud.BaseBarony
}

var BGravelines格拉沃利讷 feud.Barony = &格拉沃利讷GravelinesBarony{}

func init() {
	f := BGravelines格拉沃利讷.(*格拉沃利讷GravelinesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gravelines",
		TitleName: "格拉沃利讷",
		TitleCode: "b_gravelines",
	}
}
