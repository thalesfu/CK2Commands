package darfur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌姆马利Um_marrihBarony struct {
	feud.BaseBarony
}

var BUm_marrih乌姆马利 feud.Barony = &乌姆马利Um_marrihBarony{}

func init() {
    f := BUm_marrih乌姆马利.(*乌姆马利Um_marrihBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "um_marrih",
		TitleName: "乌姆马利",
		TitleCode: "b_um_marrih",
	}
}
