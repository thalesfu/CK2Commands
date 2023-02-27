package anjou

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贡捷堡ChateaugontierBarony struct {
	feud.BaseBarony
}

var BChateaugontier贡捷堡 feud.Barony = &贡捷堡ChateaugontierBarony{}

func init() {
    f := BChateaugontier贡捷堡.(*贡捷堡ChateaugontierBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chateaugontier",
		TitleName: "贡捷堡",
		TitleCode: "b_chateaugontier",
	}
}
