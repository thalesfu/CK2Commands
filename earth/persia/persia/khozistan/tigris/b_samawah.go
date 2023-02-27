package tigris

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨马瓦SamawahBarony struct {
	feud.BaseBarony
}

var BSamawah萨马瓦 feud.Barony = &萨马瓦SamawahBarony{}

func init() {
    f := BSamawah萨马瓦.(*萨马瓦SamawahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "samawah",
		TitleName: "萨马瓦",
		TitleCode: "b_samawah",
	}
}
