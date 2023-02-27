package kasmira

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆吉萨尔BaghsarBarony struct {
	feud.BaseBarony
}

var BBaghsar婆吉萨尔 feud.Barony = &婆吉萨尔BaghsarBarony{}

func init() {
    f := BBaghsar婆吉萨尔.(*婆吉萨尔BaghsarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baghsar",
		TitleName: "婆吉萨尔",
		TitleCode: "b_baghsar",
	}
}
