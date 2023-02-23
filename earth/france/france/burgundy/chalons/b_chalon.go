package chalons

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙隆ChalonBarony struct {
	feud.BaseBarony
}

var BChalon沙隆 feud.Barony = &沙隆ChalonBarony{}

func init() {
	f := BChalon沙隆.(*沙隆ChalonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chalon",
		TitleName: "沙隆",
		TitleCode: "b_chalon",
	}
}
