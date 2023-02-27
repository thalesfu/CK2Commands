package parma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕尔马ParmaBarony struct {
	feud.BaseBarony
}

var BParma帕尔马 feud.Barony = &帕尔马ParmaBarony{}

func init() {
    f := BParma帕尔马.(*帕尔马ParmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "parma",
		TitleName: "帕尔马",
		TitleCode: "b_parma",
	}
}
