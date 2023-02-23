package strathearn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 肯莫尔KenmoreBarony struct {
	feud.BaseBarony
}

var BKenmore肯莫尔 feud.Barony = &肯莫尔KenmoreBarony{}

func init() {
	f := BKenmore肯莫尔.(*肯莫尔KenmoreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kenmore",
		TitleName: "肯莫尔",
		TitleCode: "b_kenmore",
	}
}
