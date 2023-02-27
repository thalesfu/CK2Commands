package bulgar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰维勒TawilleBarony struct {
	feud.BaseBarony
}

var BTawille泰维勒 feud.Barony = &泰维勒TawilleBarony{}

func init() {
    f := BTawille泰维勒.(*泰维勒TawilleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tawille",
		TitleName: "泰维勒",
		TitleCode: "b_tawille",
	}
}
