package diafunu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米西里布古MissiribougouBarony struct {
	feud.BaseBarony
}

var BMissiribougou米西里布古 feud.Barony = &米西里布古MissiribougouBarony{}

func init() {
    f := BMissiribougou米西里布古.(*米西里布古MissiribougouBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "missiribougou",
		TitleName: "米西里布古",
		TitleCode: "b_missiribougou",
	}
}
