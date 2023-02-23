package blois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣艾尼昂StaignanBarony struct {
	feud.BaseBarony
}

var BStaignan圣艾尼昂 feud.Barony = &圣艾尼昂StaignanBarony{}

func init() {
	f := BStaignan圣艾尼昂.(*圣艾尼昂StaignanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "staignan",
		TitleName: "圣艾尼昂",
		TitleCode: "b_staignan",
	}
}
