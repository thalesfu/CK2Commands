package hamburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洛克施泰特LokstedtBarony struct {
	feud.BaseBarony
}

var BLokstedt洛克施泰特 feud.Barony = &洛克施泰特LokstedtBarony{}

func init() {
    f := BLokstedt洛克施泰特.(*洛克施泰特LokstedtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lokstedt",
		TitleName: "洛克施泰特",
		TitleCode: "b_lokstedt",
	}
}
