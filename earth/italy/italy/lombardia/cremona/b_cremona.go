package cremona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克雷莫纳CremonaBarony struct {
	feud.BaseBarony
}

var BCremona克雷莫纳 feud.Barony = &克雷莫纳CremonaBarony{}

func init() {
	f := BCremona克雷莫纳.(*克雷莫纳CremonaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cremona",
		TitleName: "克雷莫纳",
		TitleCode: "b_cremona",
	}
}
