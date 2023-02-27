package varadzin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托普利采TopliceBarony struct {
	feud.BaseBarony
}

var BToplice托普利采 feud.Barony = &托普利采TopliceBarony{}

func init() {
    f := BToplice托普利采.(*托普利采TopliceBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "toplice",
		TitleName: "托普利采",
		TitleCode: "b_toplice",
	}
}
