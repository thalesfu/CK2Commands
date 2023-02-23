package muscat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利兹奇LizqBarony struct {
	feud.BaseBarony
}

var BLizq利兹奇 feud.Barony = &利兹奇LizqBarony{}

func init() {
	f := BLizq利兹奇.(*利兹奇LizqBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lizq",
		TitleName: "利兹奇",
		TitleCode: "b_lizq",
	}
}
