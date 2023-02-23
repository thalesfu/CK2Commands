package vyazma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨福诺沃SafonovoBarony struct {
	feud.BaseBarony
}

var BSafonovo萨福诺沃 feud.Barony = &萨福诺沃SafonovoBarony{}

func init() {
	f := BSafonovo萨福诺沃.(*萨福诺沃SafonovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "safonovo",
		TitleName: "萨福诺沃",
		TitleCode: "b_safonovo",
	}
}
