package nalsia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尼克什奇艾AnyksciaiBarony struct {
	feud.BaseBarony
}

var BAnyksciai阿尼克什奇艾 feud.Barony = &阿尼克什奇艾AnyksciaiBarony{}

func init() {
    f := BAnyksciai阿尼克什奇艾.(*阿尼克什奇艾AnyksciaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "anyksciai",
		TitleName: "阿尼克什奇艾",
		TitleCode: "b_anyksciai",
	}
}
