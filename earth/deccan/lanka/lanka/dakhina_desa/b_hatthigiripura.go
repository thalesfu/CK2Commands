package dakhina_desa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诃悉帝耆厘补罗HatthigiripuraBarony struct {
	feud.BaseBarony
}

var BHatthigiripura诃悉帝耆厘补罗 feud.Barony = &诃悉帝耆厘补罗HatthigiripuraBarony{}

func init() {
    f := BHatthigiripura诃悉帝耆厘补罗.(*诃悉帝耆厘补罗HatthigiripuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hatthigiripura",
		TitleName: "诃悉帝耆厘补罗",
		TitleCode: "b_hatthigiripura",
	}
}
