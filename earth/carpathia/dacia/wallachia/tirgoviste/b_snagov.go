package tirgoviste

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯纳戈夫SnagovBarony struct {
	feud.BaseBarony
}

var BSnagov斯纳戈夫 feud.Barony = &斯纳戈夫SnagovBarony{}

func init() {
	f := BSnagov斯纳戈夫.(*斯纳戈夫SnagovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "snagov",
		TitleName: "斯纳戈夫",
		TitleCode: "b_snagov",
	}
}
