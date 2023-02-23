package usturt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿克苏AksuBarony struct {
	feud.BaseBarony
}

var BAksu阿克苏 feud.Barony = &阿克苏AksuBarony{}

func init() {
	f := BAksu阿克苏.(*阿克苏AksuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aksu",
		TitleName: "阿克苏",
		TitleCode: "b_aksu",
	}
}
