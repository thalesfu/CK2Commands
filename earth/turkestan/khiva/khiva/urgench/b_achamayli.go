package urgench

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿恰迈利AchamayliBarony struct {
	feud.BaseBarony
}

var BAchamayli阿恰迈利 feud.Barony = &阿恰迈利AchamayliBarony{}

func init() {
    f := BAchamayli阿恰迈利.(*阿恰迈利AchamayliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "achamayli",
		TitleName: "阿恰迈利",
		TitleCode: "b_achamayli",
	}
}
