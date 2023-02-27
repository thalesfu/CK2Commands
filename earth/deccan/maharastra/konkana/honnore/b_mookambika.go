package honnore

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩甘比加MookambikaBarony struct {
	feud.BaseBarony
}

var BMookambika摩甘比加 feud.Barony = &摩甘比加MookambikaBarony{}

func init() {
    f := BMookambika摩甘比加.(*摩甘比加MookambikaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mookambika",
		TitleName: "摩甘比加",
		TitleCode: "b_mookambika",
	}
}
