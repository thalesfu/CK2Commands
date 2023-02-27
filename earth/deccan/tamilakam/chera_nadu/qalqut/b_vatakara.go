package qalqut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伐多伽罗VatakaraBarony struct {
	feud.BaseBarony
}

var BVatakara伐多伽罗 feud.Barony = &伐多伽罗VatakaraBarony{}

func init() {
    f := BVatakara伐多伽罗.(*伐多伽罗VatakaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vatakara",
		TitleName: "伐多伽罗",
		TitleCode: "b_vatakara",
	}
}
