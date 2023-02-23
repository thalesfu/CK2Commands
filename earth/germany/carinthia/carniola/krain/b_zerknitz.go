package krain

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 齐尔克尼茨ZerknitzBarony struct {
	feud.BaseBarony
}

var BZerknitz齐尔克尼茨 feud.Barony = &齐尔克尼茨ZerknitzBarony{}

func init() {
	f := BZerknitz齐尔克尼茨.(*齐尔克尼茨ZerknitzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zerknitz",
		TitleName: "齐尔克尼茨",
		TitleCode: "b_zerknitz",
	}
}
