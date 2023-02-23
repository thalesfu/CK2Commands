package bremen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝沃施泰特BeverstedtBarony struct {
	feud.BaseBarony
}

var BBeverstedt贝沃施泰特 feud.Barony = &贝沃施泰特BeverstedtBarony{}

func init() {
	f := BBeverstedt贝沃施泰特.(*贝沃施泰特BeverstedtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "beverstedt",
		TitleName: "贝沃施泰特",
		TitleCode: "b_beverstedt",
	}
}
