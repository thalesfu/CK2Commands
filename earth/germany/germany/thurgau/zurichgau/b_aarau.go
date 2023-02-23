package zurichgau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿劳AarauBarony struct {
	feud.BaseBarony
}

var BAarau阿劳 feud.Barony = &阿劳AarauBarony{}

func init() {
	f := BAarau阿劳.(*阿劳AarauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aarau",
		TitleName: "阿劳",
		TitleCode: "b_aarau",
	}
}
