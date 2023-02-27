package travunia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特拉武尼亚TravuniaBarony struct {
	feud.BaseBarony
}

var BTravunia特拉武尼亚 feud.Barony = &特拉武尼亚TravuniaBarony{}

func init() {
    f := BTravunia特拉武尼亚.(*特拉武尼亚TravuniaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "travunia",
		TitleName: "特拉武尼亚",
		TitleCode: "b_travunia",
	}
}
