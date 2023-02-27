package antalo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 恩托托EntotoBarony struct {
	feud.BaseBarony
}

var BEntoto恩托托 feud.Barony = &恩托托EntotoBarony{}

func init() {
    f := BEntoto恩托托.(*恩托托EntotoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "entoto",
		TitleName: "恩托托",
		TitleCode: "b_entoto",
	}
}
