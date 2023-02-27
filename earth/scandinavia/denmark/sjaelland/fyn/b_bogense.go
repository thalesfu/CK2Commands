package fyn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博恩瑟BogenseBarony struct {
	feud.BaseBarony
}

var BBogense博恩瑟 feud.Barony = &博恩瑟BogenseBarony{}

func init() {
    f := BBogense博恩瑟.(*博恩瑟BogenseBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bogense",
		TitleName: "博恩瑟",
		TitleCode: "b_bogense",
	}
}
