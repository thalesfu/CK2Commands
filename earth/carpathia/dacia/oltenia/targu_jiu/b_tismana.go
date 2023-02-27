package targu_jiu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒂斯马纳TismanaBarony struct {
	feud.BaseBarony
}

var BTismana蒂斯马纳 feud.Barony = &蒂斯马纳TismanaBarony{}

func init() {
    f := BTismana蒂斯马纳.(*蒂斯马纳TismanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tismana",
		TitleName: "蒂斯马纳",
		TitleCode: "b_tismana",
	}
}
