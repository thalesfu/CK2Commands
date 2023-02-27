package kongu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 那摩迦尔NamakkalBarony struct {
	feud.BaseBarony
}

var BNamakkal那摩迦尔 feud.Barony = &那摩迦尔NamakkalBarony{}

func init() {
    f := BNamakkal那摩迦尔.(*那摩迦尔NamakkalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "namakkal",
		TitleName: "那摩迦尔",
		TitleCode: "b_namakkal",
	}
}
