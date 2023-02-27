package kongu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伽罗补KaravurBarony struct {
	feud.BaseBarony
}

var BKaravur伽罗补 feud.Barony = &伽罗补KaravurBarony{}

func init() {
    f := BKaravur伽罗补.(*伽罗补KaravurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karavur",
		TitleName: "伽罗补",
		TitleCode: "b_karavur",
	}
}
