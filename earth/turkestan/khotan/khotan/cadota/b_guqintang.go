package cadota

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古琴塘GuqintangBarony struct {
	feud.BaseBarony
}

var BGuqintang古琴塘 feud.Barony = &古琴塘GuqintangBarony{}

func init() {
    f := BGuqintang古琴塘.(*古琴塘GuqintangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "guqintang",
		TitleName: "古琴塘",
		TitleCode: "b_guqintang",
	}
}
