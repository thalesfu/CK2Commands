package al_nadjaf

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔达喀纳赫TaqtaqanahBarony struct {
	feud.BaseBarony
}

var BTaqtaqanah塔达喀纳赫 feud.Barony = &塔达喀纳赫TaqtaqanahBarony{}

func init() {
    f := BTaqtaqanah塔达喀纳赫.(*塔达喀纳赫TaqtaqanahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "taqtaqanah",
		TitleName: "塔达喀纳赫",
		TitleCode: "b_taqtaqanah",
	}
}
