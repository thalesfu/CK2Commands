package nagauda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 恶醯掣呾逻补罗AhichhatrapurBarony struct {
	feud.BaseBarony
}

var BAhichhatrapur恶醯掣呾逻补罗 feud.Barony = &恶醯掣呾逻补罗AhichhatrapurBarony{}

func init() {
	f := BAhichhatrapur恶醯掣呾逻补罗.(*恶醯掣呾逻补罗AhichhatrapurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ahichhatrapur",
		TitleName: "恶醯掣呾逻补罗",
		TitleCode: "b_ahichhatrapur",
	}
}
