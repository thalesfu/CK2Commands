package albret

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙德马桑MontdemarsanBarony struct {
	feud.BaseBarony
}

var BMontdemarsan蒙德马桑 feud.Barony = &蒙德马桑MontdemarsanBarony{}

func init() {
    f := BMontdemarsan蒙德马桑.(*蒙德马桑MontdemarsanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montdemarsan",
		TitleName: "蒙德马桑",
		TitleCode: "b_montdemarsan",
	}
}
