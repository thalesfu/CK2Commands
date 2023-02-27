package zabid

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法萨利FasalBarony struct {
	feud.BaseBarony
}

var BFasal法萨利 feud.Barony = &法萨利FasalBarony{}

func init() {
    f := BFasal法萨利.(*法萨利FasalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fasal",
		TitleName: "法萨利",
		TitleCode: "b_fasal",
	}
}
