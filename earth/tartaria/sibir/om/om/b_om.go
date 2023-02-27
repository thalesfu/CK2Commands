package om

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鄂木OmBarony struct {
	feud.BaseBarony
}

var BOm鄂木 feud.Barony = &鄂木OmBarony{}

func init() {
    f := BOm鄂木.(*鄂木OmBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "om",
		TitleName: "鄂木",
		TitleCode: "b_om",
	}
}
