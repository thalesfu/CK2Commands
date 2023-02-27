package tell_bashir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图比瑟尔TurbesselBarony struct {
	feud.BaseBarony
}

var BTurbessel图比瑟尔 feud.Barony = &图比瑟尔TurbesselBarony{}

func init() {
    f := BTurbessel图比瑟尔.(*图比瑟尔TurbesselBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "turbessel",
		TitleName: "图比瑟尔",
		TitleCode: "b_turbessel",
	}
}
