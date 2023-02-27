package nikaea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科泰严KotyaionBarony struct {
	feud.BaseBarony
}

var BKotyaion科泰严 feud.Barony = &科泰严KotyaionBarony{}

func init() {
    f := BKotyaion科泰严.(*科泰严KotyaionBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kotyaion",
		TitleName: "科泰严",
		TitleCode: "b_kotyaion",
	}
}
