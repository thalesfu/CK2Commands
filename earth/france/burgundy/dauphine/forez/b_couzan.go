package forez

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库藏CouzanBarony struct {
	feud.BaseBarony
}

var BCouzan库藏 feud.Barony = &库藏CouzanBarony{}

func init() {
	f := BCouzan库藏.(*库藏CouzanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "couzan",
		TitleName: "库藏",
		TitleCode: "b_couzan",
	}
}
