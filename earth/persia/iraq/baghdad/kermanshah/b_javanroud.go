package kermanshah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贾万鲁德JavanroudBarony struct {
	feud.BaseBarony
}

var BJavanroud贾万鲁德 feud.Barony = &贾万鲁德JavanroudBarony{}

func init() {
    f := BJavanroud贾万鲁德.(*贾万鲁德JavanroudBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "javanroud",
		TitleName: "贾万鲁德",
		TitleCode: "b_javanroud",
	}
}
