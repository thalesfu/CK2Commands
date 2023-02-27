package vaud

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗曼莫捷RomainmotierBarony struct {
	feud.BaseBarony
}

var BRomainmotier罗曼莫捷 feud.Barony = &罗曼莫捷RomainmotierBarony{}

func init() {
    f := BRomainmotier罗曼莫捷.(*罗曼莫捷RomainmotierBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "romainmotier",
		TitleName: "罗曼莫捷",
		TitleCode: "b_romainmotier",
	}
}
