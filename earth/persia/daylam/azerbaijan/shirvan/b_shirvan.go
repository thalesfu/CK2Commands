package shirvan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希尔凡ShirvanBarony struct {
	feud.BaseBarony
}

var BShirvan希尔凡 feud.Barony = &希尔凡ShirvanBarony{}

func init() {
    f := BShirvan希尔凡.(*希尔凡ShirvanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shirvan",
		TitleName: "希尔凡",
		TitleCode: "b_shirvan",
	}
}
