package rothas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吠多罗BetlaBarony struct {
	feud.BaseBarony
}

var BBetla吠多罗 feud.Barony = &吠多罗BetlaBarony{}

func init() {
	f := BBetla吠多罗.(*吠多罗BetlaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "betla",
		TitleName: "吠多罗",
		TitleCode: "b_betla",
	}
}
