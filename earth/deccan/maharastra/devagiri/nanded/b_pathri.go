package nanded

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕特里PathriBarony struct {
	feud.BaseBarony
}

var BPathri帕特里 feud.Barony = &帕特里PathriBarony{}

func init() {
	f := BPathri帕特里.(*帕特里PathriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pathri",
		TitleName: "帕特里",
		TitleCode: "b_pathri",
	}
}
