package euphrates

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜贾尔DujailBarony struct {
	feud.BaseBarony
}

var BDujail杜贾尔 feud.Barony = &杜贾尔DujailBarony{}

func init() {
	f := BDujail杜贾尔.(*杜贾尔DujailBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dujail",
		TitleName: "杜贾尔",
		TitleCode: "b_dujail",
	}
}
