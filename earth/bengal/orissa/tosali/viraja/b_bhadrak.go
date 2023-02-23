package viraja

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 跋提城BhadrakBarony struct {
	feud.BaseBarony
}

var BBhadrak跋提城 feud.Barony = &跋提城BhadrakBarony{}

func init() {
	f := BBhadrak跋提城.(*跋提城BhadrakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhadrak",
		TitleName: "跋提城",
		TitleCode: "b_bhadrak",
	}
}
