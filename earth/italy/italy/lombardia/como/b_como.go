package como

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科莫ComoBarony struct {
	feud.BaseBarony
}

var BComo科莫 feud.Barony = &科莫ComoBarony{}

func init() {
    f := BComo科莫.(*科莫ComoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "como",
		TitleName: "科莫",
		TitleCode: "b_como",
	}
}
