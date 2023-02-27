package massat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法达拉FadalaBarony struct {
	feud.BaseBarony
}

var BFadala法达拉 feud.Barony = &法达拉FadalaBarony{}

func init() {
    f := BFadala法达拉.(*法达拉FadalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fadala",
		TitleName: "法达拉",
		TitleCode: "b_fadala",
	}
}
