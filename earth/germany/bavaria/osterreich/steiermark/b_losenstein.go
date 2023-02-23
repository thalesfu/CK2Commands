package steiermark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洛森施泰因LosensteinBarony struct {
	feud.BaseBarony
}

var BLosenstein洛森施泰因 feud.Barony = &洛森施泰因LosensteinBarony{}

func init() {
	f := BLosenstein洛森施泰因.(*洛森施泰因LosensteinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "losenstein",
		TitleName: "洛森施泰因",
		TitleCode: "b_losenstein",
	}
}
