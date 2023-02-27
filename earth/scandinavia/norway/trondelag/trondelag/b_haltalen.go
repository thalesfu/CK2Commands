package trondelag

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍尔托伦HaltalenBarony struct {
	feud.BaseBarony
}

var BHaltalen霍尔托伦 feud.Barony = &霍尔托伦HaltalenBarony{}

func init() {
    f := BHaltalen霍尔托伦.(*霍尔托伦HaltalenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "haltalen",
		TitleName: "霍尔托伦",
		TitleCode: "b_haltalen",
	}
}
