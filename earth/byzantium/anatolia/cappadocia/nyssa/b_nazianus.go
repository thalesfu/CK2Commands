package nyssa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳齐安NazianusBarony struct {
	feud.BaseBarony
}

var BNazianus纳齐安 feud.Barony = &纳齐安NazianusBarony{}

func init() {
	f := BNazianus纳齐安.(*纳齐安NazianusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nazianus",
		TitleName: "纳齐安",
		TitleCode: "b_nazianus",
	}
}
