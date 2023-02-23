package malta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 海堡GozoBarony struct {
	feud.BaseBarony
}

var BGozo海堡 feud.Barony = &海堡GozoBarony{}

func init() {
	f := BGozo海堡.(*海堡GozoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gozo",
		TitleName: "海堡",
		TitleCode: "b_gozo",
	}
}
