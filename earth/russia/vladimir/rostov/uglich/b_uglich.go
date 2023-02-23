package uglich

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌格利奇UglichBarony struct {
	feud.BaseBarony
}

var BUglich乌格利奇 feud.Barony = &乌格利奇UglichBarony{}

func init() {
	f := BUglich乌格利奇.(*乌格利奇UglichBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "uglich",
		TitleName: "乌格利奇",
		TitleCode: "b_uglich",
	}
}
