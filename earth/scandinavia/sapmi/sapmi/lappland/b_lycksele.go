package lappland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吕克瑟勒LyckseleBarony struct {
	feud.BaseBarony
}

var BLycksele吕克瑟勒 feud.Barony = &吕克瑟勒LyckseleBarony{}

func init() {
	f := BLycksele吕克瑟勒.(*吕克瑟勒LyckseleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lycksele",
		TitleName: "吕克瑟勒",
		TitleCode: "b_lycksele",
	}
}
