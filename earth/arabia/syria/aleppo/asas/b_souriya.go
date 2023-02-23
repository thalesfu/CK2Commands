package asas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 叙利亚SouriyaBarony struct {
	feud.BaseBarony
}

var BSouriya叙利亚 feud.Barony = &叙利亚SouriyaBarony{}

func init() {
	f := BSouriya叙利亚.(*叙利亚SouriyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "souriya",
		TitleName: "叙利亚",
		TitleCode: "b_souriya",
	}
}
