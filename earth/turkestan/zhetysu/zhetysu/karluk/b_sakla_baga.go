package karluk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索葛莫贺Sakla_bagaBarony struct {
	feud.BaseBarony
}

var BSakla_baga索葛莫贺 feud.Barony = &索葛莫贺Sakla_bagaBarony{}

func init() {
    f := BSakla_baga索葛莫贺.(*索葛莫贺Sakla_bagaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sakla_baga",
		TitleName: "索葛莫贺",
		TitleCode: "b_sakla_baga",
	}
}
