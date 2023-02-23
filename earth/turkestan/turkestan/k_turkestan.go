package turkestan

import (
	"github.com/thalesfu/CK2Commands/earth/turkestan/turkestan/emba"
	"github.com/thalesfu/CK2Commands/earth/turkestan/turkestan/turkestan"
	"github.com/thalesfu/CK2Commands/earth/turkestan/turkestan/usturt"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TurkestanKingdom interface {
	feud.Kingdom
	DEmba恩巴() emba.EmbaDuke
	DTurkestan阿拉尔() turkestan.TurkestanDuke
	DUsturt乌斯秋尔特() usturt.UsturtDuke
}

type 乌古斯TurkestanKingdom struct {
	feud.BaseKingdom
	恩巴Emba       emba.EmbaDuke
	阿拉尔Turkestan turkestan.TurkestanDuke
	乌斯秋尔特Usturt  usturt.UsturtDuke
}

func (k *乌古斯TurkestanKingdom) DEmba恩巴() emba.EmbaDuke {
	return k.恩巴Emba
}

func (k *乌古斯TurkestanKingdom) DTurkestan阿拉尔() turkestan.TurkestanDuke {
	return k.阿拉尔Turkestan
}

func (k *乌古斯TurkestanKingdom) DUsturt乌斯秋尔特() usturt.UsturtDuke {
	return k.乌斯秋尔特Usturt
}

var KTurkestan乌古斯 TurkestanKingdom = &乌古斯TurkestanKingdom{}

func init() {
	f := KTurkestan乌古斯.(*乌古斯TurkestanKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "turkestan",
		TitleName: "乌古斯",
		TitleCode: "k_turkestan",
		Dukes:     map[string]feud.Duke{},
	}

	f.恩巴Emba = emba.DEmba恩巴
	f.恩巴Emba.SetParent(f)

	f.阿拉尔Turkestan = turkestan.DTurkestan阿拉尔
	f.阿拉尔Turkestan.SetParent(f)

	f.乌斯秋尔特Usturt = usturt.DUsturt乌斯秋尔特
	f.乌斯秋尔特Usturt.SetParent(f)

}
