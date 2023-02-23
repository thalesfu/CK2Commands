package forcalquier

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿普特AptBarony struct {
	feud.BaseBarony
}

var BApt阿普特 feud.Barony = &阿普特AptBarony{}

func init() {
	f := BApt阿普特.(*阿普特AptBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "apt",
		TitleName: "阿普特",
		TitleCode: "b_apt",
	}
}
