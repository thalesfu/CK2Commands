package vodamayutja

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 恶醯掣呾逻AhichatraBarony struct {
	feud.BaseBarony
}

var BAhichatra恶醯掣呾逻 feud.Barony = &恶醯掣呾逻AhichatraBarony{}

func init() {
    f := BAhichatra恶醯掣呾逻.(*恶醯掣呾逻AhichatraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ahichatra",
		TitleName: "恶醯掣呾逻",
		TitleCode: "b_ahichatra",
	}
}
