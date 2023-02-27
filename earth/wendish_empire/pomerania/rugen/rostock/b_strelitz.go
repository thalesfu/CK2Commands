package rostock

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 施特雷利茨StrelitzBarony struct {
	feud.BaseBarony
}

var BStrelitz施特雷利茨 feud.Barony = &施特雷利茨StrelitzBarony{}

func init() {
    f := BStrelitz施特雷利茨.(*施特雷利茨StrelitzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "strelitz",
		TitleName: "施特雷利茨",
		TitleCode: "b_strelitz",
	}
}
