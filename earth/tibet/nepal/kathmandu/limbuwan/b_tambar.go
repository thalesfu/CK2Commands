package limbuwan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 耽婆罗TambarBarony struct {
	feud.BaseBarony
}

var BTambar耽婆罗 feud.Barony = &耽婆罗TambarBarony{}

func init() {
    f := BTambar耽婆罗.(*耽婆罗TambarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tambar",
		TitleName: "耽婆罗",
		TitleCode: "b_tambar",
	}
}
