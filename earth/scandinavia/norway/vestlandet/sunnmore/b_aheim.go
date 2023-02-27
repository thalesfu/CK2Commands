package sunnmore

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥海姆AheimBarony struct {
	feud.BaseBarony
}

var BAheim奥海姆 feud.Barony = &奥海姆AheimBarony{}

func init() {
    f := BAheim奥海姆.(*奥海姆AheimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aheim",
		TitleName: "奥海姆",
		TitleCode: "b_aheim",
	}
}
