package aylik

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克拉玛依KaramayBarony struct {
	feud.BaseBarony
}

var BKaramay克拉玛依 feud.Barony = &克拉玛依KaramayBarony{}

func init() {
	f := BKaramay克拉玛依.(*克拉玛依KaramayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karamay",
		TitleName: "克拉玛依",
		TitleCode: "b_karamay",
	}
}
