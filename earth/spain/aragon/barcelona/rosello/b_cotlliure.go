package rosello

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科利乌尔CotlliureBarony struct {
	feud.BaseBarony
}

var BCotlliure科利乌尔 feud.Barony = &科利乌尔CotlliureBarony{}

func init() {
	f := BCotlliure科利乌尔.(*科利乌尔CotlliureBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cotlliure",
		TitleName: "科利乌尔",
		TitleCode: "b_cotlliure",
	}
}
