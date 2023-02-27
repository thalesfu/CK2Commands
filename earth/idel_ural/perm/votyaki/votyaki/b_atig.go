package votyaki

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿季格AtigBarony struct {
	feud.BaseBarony
}

var BAtig阿季格 feud.Barony = &阿季格AtigBarony{}

func init() {
    f := BAtig阿季格.(*阿季格AtigBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "atig",
		TitleName: "阿季格",
		TitleCode: "b_atig",
	}
}
