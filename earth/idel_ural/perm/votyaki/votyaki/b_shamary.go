package votyaki

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙马雷ShamaryBarony struct {
	feud.BaseBarony
}

var BShamary沙马雷 feud.Barony = &沙马雷ShamaryBarony{}

func init() {
    f := BShamary沙马雷.(*沙马雷ShamaryBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shamary",
		TitleName: "沙马雷",
		TitleCode: "b_shamary",
	}
}
