package pannagallu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 般那伽楼PannagalluBarony struct {
	feud.BaseBarony
}

var BPannagallu般那伽楼 feud.Barony = &般那伽楼PannagalluBarony{}

func init() {
    f := BPannagallu般那伽楼.(*般那伽楼PannagalluBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pannagallu",
		TitleName: "般那伽楼",
		TitleCode: "b_pannagallu",
	}
}
