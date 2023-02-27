package darum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达鲁姆DarumBarony struct {
	feud.BaseBarony
}

var BDarum达鲁姆 feud.Barony = &达鲁姆DarumBarony{}

func init() {
    f := BDarum达鲁姆.(*达鲁姆DarumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "darum",
		TitleName: "达鲁姆",
		TitleCode: "b_darum",
	}
}
