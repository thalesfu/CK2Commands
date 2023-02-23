package astibus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普雷舍沃PresevoBarony struct {
	feud.BaseBarony
}

var BPresevo普雷舍沃 feud.Barony = &普雷舍沃PresevoBarony{}

func init() {
	f := BPresevo普雷舍沃.(*普雷舍沃PresevoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "presevo",
		TitleName: "普雷舍沃",
		TitleCode: "b_presevo",
	}
}
