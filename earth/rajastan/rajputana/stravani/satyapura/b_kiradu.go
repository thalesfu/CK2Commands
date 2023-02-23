package satyapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉罗菟KiraduBarony struct {
	feud.BaseBarony
}

var BKiradu吉罗菟 feud.Barony = &吉罗菟KiraduBarony{}

func init() {
	f := BKiradu吉罗菟.(*吉罗菟KiraduBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kiradu",
		TitleName: "吉罗菟",
		TitleCode: "b_kiradu",
	}
}
