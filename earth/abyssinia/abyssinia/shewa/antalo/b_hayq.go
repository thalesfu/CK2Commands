package antalo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 海格HayqBarony struct {
	feud.BaseBarony
}

var BHayq海格 feud.Barony = &海格HayqBarony{}

func init() {
	f := BHayq海格.(*海格HayqBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hayq",
		TitleName: "海格",
		TitleCode: "b_hayq",
	}
}
