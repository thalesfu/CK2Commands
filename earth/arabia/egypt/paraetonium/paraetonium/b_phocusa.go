package paraetonium

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 福库萨PhocusaBarony struct {
	feud.BaseBarony
}

var BPhocusa福库萨 feud.Barony = &福库萨PhocusaBarony{}

func init() {
    f := BPhocusa福库萨.(*福库萨PhocusaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "phocusa",
		TitleName: "福库萨",
		TitleCode: "b_phocusa",
	}
}
