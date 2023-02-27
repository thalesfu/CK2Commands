package hamadan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 甘吉纳梅GanjnamehBarony struct {
	feud.BaseBarony
}

var BGanjnameh甘吉纳梅 feud.Barony = &甘吉纳梅GanjnamehBarony{}

func init() {
    f := BGanjnameh甘吉纳梅.(*甘吉纳梅GanjnamehBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ganjnameh",
		TitleName: "甘吉纳梅",
		TitleCode: "b_ganjnameh",
	}
}
