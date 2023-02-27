package tagadur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 娑底也瞢伽罗SatyamangalamBarony struct {
	feud.BaseBarony
}

var BSatyamangalam娑底也瞢伽罗 feud.Barony = &娑底也瞢伽罗SatyamangalamBarony{}

func init() {
    f := BSatyamangalam娑底也瞢伽罗.(*娑底也瞢伽罗SatyamangalamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "satyamangalam",
		TitleName: "娑底也瞢伽罗",
		TitleCode: "b_satyamangalam",
	}
}
