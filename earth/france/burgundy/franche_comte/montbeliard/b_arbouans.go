package montbeliard

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔布昂ArbouansBarony struct {
	feud.BaseBarony
}

var BArbouans阿尔布昂 feud.Barony = &阿尔布昂ArbouansBarony{}

func init() {
    f := BArbouans阿尔布昂.(*阿尔布昂ArbouansBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arbouans",
		TitleName: "阿尔布昂",
		TitleCode: "b_arbouans",
	}
}
