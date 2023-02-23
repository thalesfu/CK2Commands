package amalfi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿马尔菲AmalfiBarony struct {
	feud.BaseBarony
}

var BAmalfi阿马尔菲 feud.Barony = &阿马尔菲AmalfiBarony{}

func init() {
	f := BAmalfi阿马尔菲.(*阿马尔菲AmalfiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amalfi",
		TitleName: "阿马尔菲",
		TitleCode: "b_amalfi",
	}
}
