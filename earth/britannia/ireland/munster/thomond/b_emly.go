package thomond

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 紫杉境EmlyBarony struct {
	feud.BaseBarony
}

var BEmly紫杉境 feud.Barony = &紫杉境EmlyBarony{}

func init() {
	f := BEmly紫杉境.(*紫杉境EmlyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "emly",
		TitleName: "紫杉境",
		TitleCode: "b_emly",
	}
}
