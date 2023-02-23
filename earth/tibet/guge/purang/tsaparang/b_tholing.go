package tsaparang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托林TholingBarony struct {
	feud.BaseBarony
}

var BTholing托林 feud.Barony = &托林TholingBarony{}

func init() {
	f := BTholing托林.(*托林TholingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tholing",
		TitleName: "托林",
		TitleCode: "b_tholing",
	}
}
