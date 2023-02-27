package telemark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菲勒斯达尔FyresdalBarony struct {
	feud.BaseBarony
}

var BFyresdal菲勒斯达尔 feud.Barony = &菲勒斯达尔FyresdalBarony{}

func init() {
    f := BFyresdal菲勒斯达尔.(*菲勒斯达尔FyresdalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fyresdal",
		TitleName: "菲勒斯达尔",
		TitleCode: "b_fyresdal",
	}
}
