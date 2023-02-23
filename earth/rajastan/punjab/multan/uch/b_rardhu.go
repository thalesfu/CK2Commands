package uch

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗菟RardhuBarony struct {
	feud.BaseBarony
}

var BRardhu罗菟 feud.Barony = &罗菟RardhuBarony{}

func init() {
	f := BRardhu罗菟.(*罗菟RardhuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rardhu",
		TitleName: "罗菟",
		TitleCode: "b_rardhu",
	}
}
