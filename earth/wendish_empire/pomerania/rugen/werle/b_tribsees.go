package werle

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特里布塞斯TribseesBarony struct {
	feud.BaseBarony
}

var BTribsees特里布塞斯 feud.Barony = &特里布塞斯TribseesBarony{}

func init() {
    f := BTribsees特里布塞斯.(*特里布塞斯TribseesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tribsees",
		TitleName: "特里布塞斯",
		TitleCode: "b_tribsees",
	}
}
