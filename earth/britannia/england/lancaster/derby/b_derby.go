package derby

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德比DerbyBarony struct {
	feud.BaseBarony
}

var BDerby德比 feud.Barony = &德比DerbyBarony{}

func init() {
    f := BDerby德比.(*德比DerbyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "derby",
		TitleName: "德比",
		TitleCode: "b_derby",
	}
}
