package belz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格拉博维茨GrabowiecBarony struct {
	feud.BaseBarony
}

var BGrabowiec格拉博维茨 feud.Barony = &格拉博维茨GrabowiecBarony{}

func init() {
    f := BGrabowiec格拉博维茨.(*格拉博维茨GrabowiecBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "grabowiec",
		TitleName: "格拉博维茨",
		TitleCode: "b_grabowiec",
	}
}
