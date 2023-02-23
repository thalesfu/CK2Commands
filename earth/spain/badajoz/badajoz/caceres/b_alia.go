package caceres

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿利亚AliaBarony struct {
	feud.BaseBarony
}

var BAlia阿利亚 feud.Barony = &阿利亚AliaBarony{}

func init() {
	f := BAlia阿利亚.(*阿利亚AliaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alia",
		TitleName: "阿利亚",
		TitleCode: "b_alia",
	}
}
