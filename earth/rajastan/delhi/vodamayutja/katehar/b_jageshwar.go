package katehar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阇祇湿伐罗JageshwarBarony struct {
	feud.BaseBarony
}

var BJageshwar阇祇湿伐罗 feud.Barony = &阇祇湿伐罗JageshwarBarony{}

func init() {
	f := BJageshwar阇祇湿伐罗.(*阇祇湿伐罗JageshwarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jageshwar",
		TitleName: "阇祇湿伐罗",
		TitleCode: "b_jageshwar",
	}
}
