package viscaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托洛萨TolosaBarony struct {
	feud.BaseBarony
}

var BTolosa托洛萨 feud.Barony = &托洛萨TolosaBarony{}

func init() {
	f := BTolosa托洛萨.(*托洛萨TolosaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tolosa",
		TitleName: "托洛萨",
		TitleCode: "b_tolosa",
	}
}
