package gemer

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 大勒采NagyroceBarony struct {
	feud.BaseBarony
}

var BNagyroce大勒采 feud.Barony = &大勒采NagyroceBarony{}

func init() {
    f := BNagyroce大勒采.(*大勒采NagyroceBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nagyroce",
		TitleName: "大勒采",
		TitleCode: "b_nagyroce",
	}
}
