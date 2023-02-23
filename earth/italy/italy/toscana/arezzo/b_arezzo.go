package arezzo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿雷佐ArezzoBarony struct {
	feud.BaseBarony
}

var BArezzo阿雷佐 feud.Barony = &阿雷佐ArezzoBarony{}

func init() {
	f := BArezzo阿雷佐.(*阿雷佐ArezzoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arezzo",
		TitleName: "阿雷佐",
		TitleCode: "b_arezzo",
	}
}
