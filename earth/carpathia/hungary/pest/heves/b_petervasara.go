package heves

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 彼得瓦沙劳PetervasaraBarony struct {
	feud.BaseBarony
}

var BPetervasara彼得瓦沙劳 feud.Barony = &彼得瓦沙劳PetervasaraBarony{}

func init() {
	f := BPetervasara彼得瓦沙劳.(*彼得瓦沙劳PetervasaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "petervasara",
		TitleName: "彼得瓦沙劳",
		TitleCode: "b_petervasara",
	}
}
