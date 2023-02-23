package amdo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 聂荣NyainrongBarony struct {
	feud.BaseBarony
}

var BNyainrong聂荣 feud.Barony = &聂荣NyainrongBarony{}

func init() {
	f := BNyainrong聂荣.(*聂荣NyainrongBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nyainrong",
		TitleName: "聂荣",
		TitleCode: "b_nyainrong",
	}
}
