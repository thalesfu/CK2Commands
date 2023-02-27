package koloneia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 盖尔贾尼斯GerjanisBarony struct {
	feud.BaseBarony
}

var BGerjanis盖尔贾尼斯 feud.Barony = &盖尔贾尼斯GerjanisBarony{}

func init() {
    f := BGerjanis盖尔贾尼斯.(*盖尔贾尼斯GerjanisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gerjanis",
		TitleName: "盖尔贾尼斯",
		TitleCode: "b_gerjanis",
	}
}
