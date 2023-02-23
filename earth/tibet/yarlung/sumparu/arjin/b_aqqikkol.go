package arjin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿其克库勒AqqikkolBarony struct {
	feud.BaseBarony
}

var BAqqikkol阿其克库勒 feud.Barony = &阿其克库勒AqqikkolBarony{}

func init() {
	f := BAqqikkol阿其克库勒.(*阿其克库勒AqqikkolBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aqqikkol",
		TitleName: "阿其克库勒",
		TitleCode: "b_aqqikkol",
	}
}
