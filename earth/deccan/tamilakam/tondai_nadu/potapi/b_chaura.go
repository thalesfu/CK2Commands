package potapi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 招罗ChauraBarony struct {
	feud.BaseBarony
}

var BChaura招罗 feud.Barony = &招罗ChauraBarony{}

func init() {
    f := BChaura招罗.(*招罗ChauraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chaura",
		TitleName: "招罗",
		TitleCode: "b_chaura",
	}
}
