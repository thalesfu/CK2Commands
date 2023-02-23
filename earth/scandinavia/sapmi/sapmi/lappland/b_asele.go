package lappland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥瑟勒AseleBarony struct {
	feud.BaseBarony
}

var BAsele奥瑟勒 feud.Barony = &奥瑟勒AseleBarony{}

func init() {
	f := BAsele奥瑟勒.(*奥瑟勒AseleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "asele",
		TitleName: "奥瑟勒",
		TitleCode: "b_asele",
	}
}
