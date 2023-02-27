package ili

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特克斯TekesBarony struct {
	feud.BaseBarony
}

var BTekes特克斯 feud.Barony = &特克斯TekesBarony{}

func init() {
    f := BTekes特克斯.(*特克斯TekesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tekes",
		TitleName: "特克斯",
		TitleCode: "b_tekes",
	}
}
