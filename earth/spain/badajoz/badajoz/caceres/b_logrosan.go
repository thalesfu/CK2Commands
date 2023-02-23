package caceres

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洛格罗桑LogrosanBarony struct {
	feud.BaseBarony
}

var BLogrosan洛格罗桑 feud.Barony = &洛格罗桑LogrosanBarony{}

func init() {
	f := BLogrosan洛格罗桑.(*洛格罗桑LogrosanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "logrosan",
		TitleName: "洛格罗桑",
		TitleCode: "b_logrosan",
	}
}
