package fuqi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 才尔琼CaierqiongBarony struct {
	feud.BaseBarony
}

var BCaierqiong才尔琼 feud.Barony = &才尔琼CaierqiongBarony{}

func init() {
    f := BCaierqiong才尔琼.(*才尔琼CaierqiongBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "caierqiong",
		TitleName: "才尔琼",
		TitleCode: "b_caierqiong",
	}
}
