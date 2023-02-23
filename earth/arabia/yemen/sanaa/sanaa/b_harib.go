package sanaa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈里卜HaribBarony struct {
	feud.BaseBarony
}

var BHarib哈里卜 feud.Barony = &哈里卜HaribBarony{}

func init() {
	f := BHarib哈里卜.(*哈里卜HaribBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "harib",
		TitleName: "哈里卜",
		TitleCode: "b_harib",
	}
}
