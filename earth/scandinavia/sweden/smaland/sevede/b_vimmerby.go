package sevede

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维默比VimmerbyBarony struct {
	feud.BaseBarony
}

var BVimmerby维默比 feud.Barony = &维默比VimmerbyBarony{}

func init() {
	f := BVimmerby维默比.(*维默比VimmerbyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vimmerby",
		TitleName: "维默比",
		TitleCode: "b_vimmerby",
	}
}
