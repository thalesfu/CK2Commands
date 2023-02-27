package dimapur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诃夷卢摩毗阿罗HerombialBarony struct {
	feud.BaseBarony
}

var BHerombial诃夷卢摩毗阿罗 feud.Barony = &诃夷卢摩毗阿罗HerombialBarony{}

func init() {
    f := BHerombial诃夷卢摩毗阿罗.(*诃夷卢摩毗阿罗HerombialBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "herombial",
		TitleName: "诃夷卢摩毗阿罗",
		TitleCode: "b_herombial",
	}
}
