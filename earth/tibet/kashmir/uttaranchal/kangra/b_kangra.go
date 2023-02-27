package kangra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 建伽罗KangraBarony struct {
	feud.BaseBarony
}

var BKangra建伽罗 feud.Barony = &建伽罗KangraBarony{}

func init() {
    f := BKangra建伽罗.(*建伽罗KangraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kangra",
		TitleName: "建伽罗",
		TitleCode: "b_kangra",
	}
}
