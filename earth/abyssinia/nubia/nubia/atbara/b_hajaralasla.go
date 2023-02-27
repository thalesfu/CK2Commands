package atbara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈贾雷阿萨HajaralaslaBarony struct {
	feud.BaseBarony
}

var BHajaralasla哈贾雷阿萨 feud.Barony = &哈贾雷阿萨HajaralaslaBarony{}

func init() {
    f := BHajaralasla哈贾雷阿萨.(*哈贾雷阿萨HajaralaslaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hajaralasla",
		TitleName: "哈贾雷阿萨",
		TitleCode: "b_hajaralasla",
	}
}
