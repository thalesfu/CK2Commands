package pereyaslavl_zalessky

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚历山德罗夫AleksandrovBarony struct {
	feud.BaseBarony
}

var BAleksandrov亚历山德罗夫 feud.Barony = &亚历山德罗夫AleksandrovBarony{}

func init() {
    f := BAleksandrov亚历山德罗夫.(*亚历山德罗夫AleksandrovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aleksandrov",
		TitleName: "亚历山德罗夫",
		TitleCode: "b_aleksandrov",
	}
}
