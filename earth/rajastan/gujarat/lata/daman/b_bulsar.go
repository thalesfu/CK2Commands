package daman

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 补萨尔BulsarBarony struct {
	feud.BaseBarony
}

var BBulsar补萨尔 feud.Barony = &补萨尔BulsarBarony{}

func init() {
    f := BBulsar补萨尔.(*补萨尔BulsarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bulsar",
		TitleName: "补萨尔",
		TitleCode: "b_bulsar",
	}
}
