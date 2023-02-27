package thisageta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜撒该塔伊ThisagetaBarony struct {
	feud.BaseBarony
}

var BThisageta杜撒该塔伊 feud.Barony = &杜撒该塔伊ThisagetaBarony{}

func init() {
    f := BThisageta杜撒该塔伊.(*杜撒该塔伊ThisagetaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "thisageta",
		TitleName: "杜撒该塔伊",
		TitleCode: "b_thisageta",
	}
}
