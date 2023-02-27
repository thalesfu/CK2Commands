package lut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿玛尼什AmaniehBarony struct {
	feud.BaseBarony
}

var BAmanieh阿玛尼什 feud.Barony = &阿玛尼什AmaniehBarony{}

func init() {
    f := BAmanieh阿玛尼什.(*阿玛尼什AmaniehBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amanieh",
		TitleName: "阿玛尼什",
		TitleCode: "b_amanieh",
	}
}
