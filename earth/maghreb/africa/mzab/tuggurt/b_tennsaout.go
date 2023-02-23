package tuggurt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 滕绍特TennsaoutBarony struct {
	feud.BaseBarony
}

var BTennsaout滕绍特 feud.Barony = &滕绍特TennsaoutBarony{}

func init() {
	f := BTennsaout滕绍特.(*滕绍特TennsaoutBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tennsaout",
		TitleName: "滕绍特",
		TitleCode: "b_tennsaout",
	}
}
