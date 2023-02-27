package vastmanland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙克托普MunktorpBarony struct {
	feud.BaseBarony
}

var BMunktorp蒙克托普 feud.Barony = &蒙克托普MunktorpBarony{}

func init() {
    f := BMunktorp蒙克托普.(*蒙克托普MunktorpBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "munktorp",
		TitleName: "蒙克托普",
		TitleCode: "b_munktorp",
	}
}
