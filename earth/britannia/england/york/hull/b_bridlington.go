package hull

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布里德灵顿BridlingtonBarony struct {
	feud.BaseBarony
}

var BBridlington布里德灵顿 feud.Barony = &布里德灵顿BridlingtonBarony{}

func init() {
	f := BBridlington布里德灵顿.(*布里德灵顿BridlingtonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bridlington",
		TitleName: "布里德灵顿",
		TitleCode: "b_bridlington",
	}
}
