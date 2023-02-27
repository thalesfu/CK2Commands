package palencia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣安托林Saint_antolinBarony struct {
	feud.BaseBarony
}

var BSaint_antolin圣安托林 feud.Barony = &圣安托林Saint_antolinBarony{}

func init() {
    f := BSaint_antolin圣安托林.(*圣安托林Saint_antolinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saint_antolin",
		TitleName: "圣安托林",
		TitleCode: "b_saint_antolin",
	}
}
