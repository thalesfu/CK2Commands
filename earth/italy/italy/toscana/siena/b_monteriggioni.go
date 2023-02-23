package siena

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙特里焦尼MonteriggioniBarony struct {
	feud.BaseBarony
}

var BMonteriggioni蒙特里焦尼 feud.Barony = &蒙特里焦尼MonteriggioniBarony{}

func init() {
	f := BMonteriggioni蒙特里焦尼.(*蒙特里焦尼MonteriggioniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "monteriggioni",
		TitleName: "蒙特里焦尼",
		TitleCode: "b_monteriggioni",
	}
}
