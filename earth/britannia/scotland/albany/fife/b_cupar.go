package fife

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库珀CuparBarony struct {
	feud.BaseBarony
}

var BCupar库珀 feud.Barony = &库珀CuparBarony{}

func init() {
	f := BCupar库珀.(*库珀CuparBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cupar",
		TitleName: "库珀",
		TitleCode: "b_cupar",
	}
}
