package sijilmasa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提瑟尔德明TisserdmineBarony struct {
	feud.BaseBarony
}

var BTisserdmine提瑟尔德明 feud.Barony = &提瑟尔德明TisserdmineBarony{}

func init() {
    f := BTisserdmine提瑟尔德明.(*提瑟尔德明TisserdmineBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tisserdmine",
		TitleName: "提瑟尔德明",
		TitleCode: "b_tisserdmine",
	}
}
