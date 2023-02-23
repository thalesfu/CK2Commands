package warwick

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯尼尔沃思KenilworthBarony struct {
	feud.BaseBarony
}

var BKenilworth凯尼尔沃思 feud.Barony = &凯尼尔沃思KenilworthBarony{}

func init() {
	f := BKenilworth凯尼尔沃思.(*凯尼尔沃思KenilworthBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kenilworth",
		TitleName: "凯尼尔沃思",
		TitleCode: "b_kenilworth",
	}
}
