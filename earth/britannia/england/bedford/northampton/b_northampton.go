package northampton

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 北安普顿NorthamptonBarony struct {
	feud.BaseBarony
}

var BNorthampton北安普顿 feud.Barony = &北安普顿NorthamptonBarony{}

func init() {
	f := BNorthampton北安普顿.(*北安普顿NorthamptonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "northampton",
		TitleName: "北安普顿",
		TitleCode: "b_northampton",
	}
}
