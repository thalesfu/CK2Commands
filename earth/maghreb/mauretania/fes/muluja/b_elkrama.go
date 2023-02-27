package muluja

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克拉马ElkramaBarony struct {
	feud.BaseBarony
}

var BElkrama克拉马 feud.Barony = &克拉马ElkramaBarony{}

func init() {
    f := BElkrama克拉马.(*克拉马ElkramaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "elkrama",
		TitleName: "克拉马",
		TitleCode: "b_elkrama",
	}
}
