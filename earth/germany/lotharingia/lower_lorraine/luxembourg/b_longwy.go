package luxembourg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 隆维LongwyBarony struct {
	feud.BaseBarony
}

var BLongwy隆维 feud.Barony = &隆维LongwyBarony{}

func init() {
    f := BLongwy隆维.(*隆维LongwyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "longwy",
		TitleName: "隆维",
		TitleCode: "b_longwy",
	}
}
