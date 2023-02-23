package mainling

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卧龙WolongBarony struct {
	feud.BaseBarony
}

var BWolong卧龙 feud.Barony = &卧龙WolongBarony{}

func init() {
	f := BWolong卧龙.(*卧龙WolongBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wolong",
		TitleName: "卧龙",
		TitleCode: "b_wolong",
	}
}
