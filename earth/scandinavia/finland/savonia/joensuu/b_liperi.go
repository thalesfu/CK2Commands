package joensuu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利佩里LiperiBarony struct {
	feud.BaseBarony
}

var BLiperi利佩里 feud.Barony = &利佩里LiperiBarony{}

func init() {
    f := BLiperi利佩里.(*利佩里LiperiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "liperi",
		TitleName: "利佩里",
		TitleCode: "b_liperi",
	}
}
