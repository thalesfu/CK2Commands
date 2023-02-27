package crimea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克兹列夫KezlevBarony struct {
	feud.BaseBarony
}

var BKezlev克兹列夫 feud.Barony = &克兹列夫KezlevBarony{}

func init() {
    f := BKezlev克兹列夫.(*克兹列夫KezlevBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kezlev",
		TitleName: "克兹列夫",
		TitleCode: "b_kezlev",
	}
}
