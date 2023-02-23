package navasarika

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 呾计湿伐罗TadkeshwarBarony struct {
	feud.BaseBarony
}

var BTadkeshwar呾计湿伐罗 feud.Barony = &呾计湿伐罗TadkeshwarBarony{}

func init() {
	f := BTadkeshwar呾计湿伐罗.(*呾计湿伐罗TadkeshwarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tadkeshwar",
		TitleName: "呾计湿伐罗",
		TitleCode: "b_tadkeshwar",
	}
}
