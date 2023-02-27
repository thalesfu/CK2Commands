package saqsin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 撒哈辛SaqsinBarony struct {
	feud.BaseBarony
}

var BSaqsin撒哈辛 feud.Barony = &撒哈辛SaqsinBarony{}

func init() {
    f := BSaqsin撒哈辛.(*撒哈辛SaqsinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saqsin",
		TitleName: "撒哈辛",
		TitleCode: "b_saqsin",
	}
}
