package vermandois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库西CoucyBarony struct {
	feud.BaseBarony
}

var BCoucy库西 feud.Barony = &库西CoucyBarony{}

func init() {
    f := BCoucy库西.(*库西CoucyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "coucy",
		TitleName: "库西",
		TitleCode: "b_coucy",
	}
}
