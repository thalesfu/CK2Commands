package medjybij

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 列季奇夫LetychivBarony struct {
	feud.BaseBarony
}

var BLetychiv列季奇夫 feud.Barony = &列季奇夫LetychivBarony{}

func init() {
    f := BLetychiv列季奇夫.(*列季奇夫LetychivBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "letychiv",
		TitleName: "列季奇夫",
		TitleCode: "b_letychiv",
	}
}
