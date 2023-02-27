package perfeddwlad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 登比DenbighBarony struct {
	feud.BaseBarony
}

var BDenbigh登比 feud.Barony = &登比DenbighBarony{}

func init() {
    f := BDenbigh登比.(*登比DenbighBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "denbigh",
		TitleName: "登比",
		TitleCode: "b_denbigh",
	}
}
