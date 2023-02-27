package mahoba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿阇耶姞利呬AjaigarhBarony struct {
	feud.BaseBarony
}

var BAjaigarh阿阇耶姞利呬 feud.Barony = &阿阇耶姞利呬AjaigarhBarony{}

func init() {
    f := BAjaigarh阿阇耶姞利呬.(*阿阇耶姞利呬AjaigarhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ajaigarh",
		TitleName: "阿阇耶姞利呬",
		TitleCode: "b_ajaigarh",
	}
}
