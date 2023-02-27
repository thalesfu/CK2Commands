package monkh_khairkhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 富蕴FuyunBarony struct {
	feud.BaseBarony
}

var BFuyun富蕴 feud.Barony = &富蕴FuyunBarony{}

func init() {
    f := BFuyun富蕴.(*富蕴FuyunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fuyun",
		TitleName: "富蕴",
		TitleCode: "b_fuyun",
	}
}
