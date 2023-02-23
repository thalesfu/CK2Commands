package orleans

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥尔良OrleansBarony struct {
	feud.BaseBarony
}

var BOrleans奥尔良 feud.Barony = &奥尔良OrleansBarony{}

func init() {
	f := BOrleans奥尔良.(*奥尔良OrleansBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "orleans",
		TitleName: "奥尔良",
		TitleCode: "b_orleans",
	}
}
