package schwyz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 屈斯纳赫特KussnachtBarony struct {
	feud.BaseBarony
}

var BKussnacht屈斯纳赫特 feud.Barony = &屈斯纳赫特KussnachtBarony{}

func init() {
    f := BKussnacht屈斯纳赫特.(*屈斯纳赫特KussnachtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kussnacht",
		TitleName: "屈斯纳赫特",
		TitleCode: "b_kussnacht",
	}
}
