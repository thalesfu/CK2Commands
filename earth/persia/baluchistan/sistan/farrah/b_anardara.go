package farrah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿那尔达雷AnardaraBarony struct {
	feud.BaseBarony
}

var BAnardara阿那尔达雷 feud.Barony = &阿那尔达雷AnardaraBarony{}

func init() {
	f := BAnardara阿那尔达雷.(*阿那尔达雷AnardaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "anardara",
		TitleName: "阿那尔达雷",
		TitleCode: "b_anardara",
	}
}
