package vairata

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曷利沙那他HarshnathBarony struct {
	feud.BaseBarony
}

var BHarshnath曷利沙那他 feud.Barony = &曷利沙那他HarshnathBarony{}

func init() {
    f := BHarshnath曷利沙那他.(*曷利沙那他HarshnathBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "harshnath",
		TitleName: "曷利沙那他",
		TitleCode: "b_harshnath",
	}
}
