package baygal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝加尔BaygalBarony struct {
	feud.BaseBarony
}

var BBaygal贝加尔 feud.Barony = &贝加尔BaygalBarony{}

func init() {
	f := BBaygal贝加尔.(*贝加尔BaygalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baygal",
		TitleName: "贝加尔",
		TitleCode: "b_baygal",
	}
}
