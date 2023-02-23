package medog

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格当GedangBarony struct {
	feud.BaseBarony
}

var BGedang格当 feud.Barony = &格当GedangBarony{}

func init() {
	f := BGedang格当.(*格当GedangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gedang",
		TitleName: "格当",
		TitleCode: "b_gedang",
	}
}
