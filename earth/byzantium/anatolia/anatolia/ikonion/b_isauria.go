package ikonion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊苏里亚IsauriaBarony struct {
	feud.BaseBarony
}

var BIsauria伊苏里亚 feud.Barony = &伊苏里亚IsauriaBarony{}

func init() {
    f := BIsauria伊苏里亚.(*伊苏里亚IsauriaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "isauria",
		TitleName: "伊苏里亚",
		TitleCode: "b_isauria",
	}
}
