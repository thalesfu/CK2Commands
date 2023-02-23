package iasi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 温盖尼UngheniBarony struct {
	feud.BaseBarony
}

var BUngheni温盖尼 feud.Barony = &温盖尼UngheniBarony{}

func init() {
	f := BUngheni温盖尼.(*温盖尼UngheniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ungheni",
		TitleName: "温盖尼",
		TitleCode: "b_ungheni",
	}
}
