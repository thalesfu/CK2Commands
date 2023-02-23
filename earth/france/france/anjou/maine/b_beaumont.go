package maine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博蒙BeaumontBarony struct {
	feud.BaseBarony
}

var BBeaumont博蒙 feud.Barony = &博蒙BeaumontBarony{}

func init() {
	f := BBeaumont博蒙.(*博蒙BeaumontBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "beaumont",
		TitleName: "博蒙",
		TitleCode: "b_beaumont",
	}
}
