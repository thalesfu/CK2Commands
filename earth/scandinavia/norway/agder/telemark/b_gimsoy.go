package telemark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊姆绥GimsoyBarony struct {
	feud.BaseBarony
}

var BGimsoy伊姆绥 feud.Barony = &伊姆绥GimsoyBarony{}

func init() {
	f := BGimsoy伊姆绥.(*伊姆绥GimsoyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gimsoy",
		TitleName: "伊姆绥",
		TitleCode: "b_gimsoy",
	}
}
