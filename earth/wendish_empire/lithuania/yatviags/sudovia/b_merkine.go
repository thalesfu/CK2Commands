package sudovia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅尔基内MerkineBarony struct {
	feud.BaseBarony
}

var BMerkine梅尔基内 feud.Barony = &梅尔基内MerkineBarony{}

func init() {
    f := BMerkine梅尔基内.(*梅尔基内MerkineBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "merkine",
		TitleName: "梅尔基内",
		TitleCode: "b_merkine",
	}
}
