package druz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞勒海德SalkhardBarony struct {
	feud.BaseBarony
}

var BSalkhard塞勒海德 feud.Barony = &塞勒海德SalkhardBarony{}

func init() {
    f := BSalkhard塞勒海德.(*塞勒海德SalkhardBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "salkhard",
		TitleName: "塞勒海德",
		TitleCode: "b_salkhard",
	}
}
