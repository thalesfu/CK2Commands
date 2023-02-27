package derbent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔乌斯TayusBarony struct {
	feud.BaseBarony
}

var BTayus塔乌斯 feud.Barony = &塔乌斯TayusBarony{}

func init() {
    f := BTayus塔乌斯.(*塔乌斯TayusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tayus",
		TitleName: "塔乌斯",
		TitleCode: "b_tayus",
	}
}
