package dyfed

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈弗福德HaverfordBarony struct {
	feud.BaseBarony
}

var BHaverford哈弗福德 feud.Barony = &哈弗福德HaverfordBarony{}

func init() {
	f := BHaverford哈弗福德.(*哈弗福德HaverfordBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "haverford",
		TitleName: "哈弗福德",
		TitleCode: "b_haverford",
	}
}
