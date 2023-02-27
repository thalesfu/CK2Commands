package hijaz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰马TaymaBarony struct {
	feud.BaseBarony
}

var BTayma泰马 feud.Barony = &泰马TaymaBarony{}

func init() {
    f := BTayma泰马.(*泰马TaymaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tayma",
		TitleName: "泰马",
		TitleCode: "b_tayma",
	}
}
