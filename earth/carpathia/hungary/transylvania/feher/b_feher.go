package feher

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 费黑尔FeherBarony struct {
	feud.BaseBarony
}

var BFeher费黑尔 feud.Barony = &费黑尔FeherBarony{}

func init() {
	f := BFeher费黑尔.(*费黑尔FeherBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "feher",
		TitleName: "费黑尔",
		TitleCode: "b_feher",
	}
}
