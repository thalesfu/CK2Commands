package amiens

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 内勒NesleBarony struct {
	feud.BaseBarony
}

var BNesle内勒 feud.Barony = &内勒NesleBarony{}

func init() {
	f := BNesle内勒.(*内勒NesleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nesle",
		TitleName: "内勒",
		TitleCode: "b_nesle",
	}
}
