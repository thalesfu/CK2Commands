package charsianon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科马纳Comana_charsianonBarony struct {
	feud.BaseBarony
}

var BComana_charsianon科马纳 feud.Barony = &科马纳Comana_charsianonBarony{}

func init() {
    f := BComana_charsianon科马纳.(*科马纳Comana_charsianonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "comana_charsianon",
		TitleName: "科马纳",
		TitleCode: "b_comana_charsianon",
	}
}
