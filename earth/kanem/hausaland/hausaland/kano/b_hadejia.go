package kano

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈代贾HadejiaBarony struct {
	feud.BaseBarony
}

var BHadejia哈代贾 feud.Barony = &哈代贾HadejiaBarony{}

func init() {
	f := BHadejia哈代贾.(*哈代贾HadejiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hadejia",
		TitleName: "哈代贾",
		TitleCode: "b_hadejia",
	}
}
