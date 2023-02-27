package vengipura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳勒瑟布尔姆NarasapuramBarony struct {
	feud.BaseBarony
}

var BNarasapuram纳勒瑟布尔姆 feud.Barony = &纳勒瑟布尔姆NarasapuramBarony{}

func init() {
    f := BNarasapuram纳勒瑟布尔姆.(*纳勒瑟布尔姆NarasapuramBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "narasapuram",
		TitleName: "纳勒瑟布尔姆",
		TitleCode: "b_narasapuram",
	}
}
