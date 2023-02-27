package magdeburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿恩施泰因ArnsteinBarony struct {
	feud.BaseBarony
}

var BArnstein阿恩施泰因 feud.Barony = &阿恩施泰因ArnsteinBarony{}

func init() {
    f := BArnstein阿恩施泰因.(*阿恩施泰因ArnsteinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arnstein",
		TitleName: "阿恩施泰因",
		TitleCode: "b_arnstein",
	}
}
