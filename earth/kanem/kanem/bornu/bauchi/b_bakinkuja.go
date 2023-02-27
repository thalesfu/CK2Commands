package bauchi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴金库贾BakinkujaBarony struct {
	feud.BaseBarony
}

var BBakinkuja巴金库贾 feud.Barony = &巴金库贾BakinkujaBarony{}

func init() {
    f := BBakinkuja巴金库贾.(*巴金库贾BakinkujaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bakinkuja",
		TitleName: "巴金库贾",
		TitleCode: "b_bakinkuja",
	}
}
