package meissen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔滕堡BelgernBarony struct {
	feud.BaseBarony
}

var BBelgern阿尔滕堡 feud.Barony = &阿尔滕堡BelgernBarony{}

func init() {
    f := BBelgern阿尔滕堡.(*阿尔滕堡BelgernBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "belgern",
		TitleName: "阿尔滕堡",
		TitleCode: "b_belgern",
	}
}
