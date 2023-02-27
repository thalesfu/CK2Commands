package galindia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍亨施泰因HohensteinBarony struct {
	feud.BaseBarony
}

var BHohenstein霍亨施泰因 feud.Barony = &霍亨施泰因HohensteinBarony{}

func init() {
    f := BHohenstein霍亨施泰因.(*霍亨施泰因HohensteinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hohenstein",
		TitleName: "霍亨施泰因",
		TitleCode: "b_hohenstein",
	}
}
