package benghazi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图勒迈塞TulmaytathBarony struct {
	feud.BaseBarony
}

var BTulmaytath图勒迈塞 feud.Barony = &图勒迈塞TulmaytathBarony{}

func init() {
    f := BTulmaytath图勒迈塞.(*图勒迈塞TulmaytathBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tulmaytath",
		TitleName: "图勒迈塞",
		TitleCode: "b_tulmaytath",
	}
}
