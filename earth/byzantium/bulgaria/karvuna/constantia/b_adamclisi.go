package constantia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿达姆克利西AdamclisiBarony struct {
	feud.BaseBarony
}

var BAdamclisi阿达姆克利西 feud.Barony = &阿达姆克利西AdamclisiBarony{}

func init() {
    f := BAdamclisi阿达姆克利西.(*阿达姆克利西AdamclisiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "adamclisi",
		TitleName: "阿达姆克利西",
		TitleCode: "b_adamclisi",
	}
}
