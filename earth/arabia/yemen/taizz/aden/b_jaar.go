package aden

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰阿尔JaarBarony struct {
	feud.BaseBarony
}

var BJaar杰阿尔 feud.Barony = &杰阿尔JaarBarony{}

func init() {
	f := BJaar杰阿尔.(*杰阿尔JaarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jaar",
		TitleName: "杰阿尔",
		TitleCode: "b_jaar",
	}
}
