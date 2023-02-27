package varmland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 约瑟JosseBarony struct {
	feud.BaseBarony
}

var BJosse约瑟 feud.Barony = &约瑟JosseBarony{}

func init() {
    f := BJosse约瑟.(*约瑟JosseBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "josse",
		TitleName: "约瑟",
		TitleCode: "b_josse",
	}
}
