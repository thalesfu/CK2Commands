package devagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 医罗补罗ElloraBarony struct {
	feud.BaseBarony
}

var BEllora医罗补罗 feud.Barony = &医罗补罗ElloraBarony{}

func init() {
	f := BEllora医罗补罗.(*医罗补罗ElloraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ellora",
		TitleName: "医罗补罗",
		TitleCode: "b_ellora",
	}
}
