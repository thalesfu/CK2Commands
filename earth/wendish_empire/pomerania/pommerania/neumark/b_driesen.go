package neumark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德里森DriesenBarony struct {
	feud.BaseBarony
}

var BDriesen德里森 feud.Barony = &德里森DriesenBarony{}

func init() {
    f := BDriesen德里森.(*德里森DriesenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "driesen",
		TitleName: "德里森",
		TitleCode: "b_driesen",
	}
}
