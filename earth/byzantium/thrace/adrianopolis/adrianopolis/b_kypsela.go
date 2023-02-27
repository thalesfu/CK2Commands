package adrianopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基普瑟拉KypselaBarony struct {
	feud.BaseBarony
}

var BKypsela基普瑟拉 feud.Barony = &基普瑟拉KypselaBarony{}

func init() {
    f := BKypsela基普瑟拉.(*基普瑟拉KypselaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kypsela",
		TitleName: "基普瑟拉",
		TitleCode: "b_kypsela",
	}
}
