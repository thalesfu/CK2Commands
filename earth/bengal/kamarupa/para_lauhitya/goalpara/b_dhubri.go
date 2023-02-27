package goalpara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菟婆利DhubriBarony struct {
	feud.BaseBarony
}

var BDhubri菟婆利 feud.Barony = &菟婆利DhubriBarony{}

func init() {
    f := BDhubri菟婆利.(*菟婆利DhubriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dhubri",
		TitleName: "菟婆利",
		TitleCode: "b_dhubri",
	}
}
