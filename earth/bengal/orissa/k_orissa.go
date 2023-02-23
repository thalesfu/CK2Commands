package orissa

import (
	"github.com/thalesfu/CK2Commands/earth/bengal/orissa/dandakaranya"
	"github.com/thalesfu/CK2Commands/earth/bengal/orissa/kalinga"
	"github.com/thalesfu/CK2Commands/earth/bengal/orissa/tosali"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OrissaKingdom interface {
	feud.Kingdom
	DDandakaranya弹宅迦林() dandakaranya.DandakaranyaDuke
	DKalinga羯陵伽() kalinga.KalingaDuke
	DTosali睹舍离() tosali.TosaliDuke
}

type 乌里舍OrissaKingdom struct {
	feud.BaseKingdom
	弹宅迦林Dandakaranya dandakaranya.DandakaranyaDuke
	羯陵伽Kalinga       kalinga.KalingaDuke
	睹舍离Tosali        tosali.TosaliDuke
}

func (k *乌里舍OrissaKingdom) DDandakaranya弹宅迦林() dandakaranya.DandakaranyaDuke {
	return k.弹宅迦林Dandakaranya
}

func (k *乌里舍OrissaKingdom) DKalinga羯陵伽() kalinga.KalingaDuke {
	return k.羯陵伽Kalinga
}

func (k *乌里舍OrissaKingdom) DTosali睹舍离() tosali.TosaliDuke {
	return k.睹舍离Tosali
}

var KOrissa乌里舍 OrissaKingdom = &乌里舍OrissaKingdom{}

func init() {
	f := KOrissa乌里舍.(*乌里舍OrissaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "orissa",
		TitleName: "乌里舍",
		TitleCode: "k_orissa",
		Dukes:     map[string]feud.Duke{},
	}

	f.弹宅迦林Dandakaranya = dandakaranya.DDandakaranya弹宅迦林
	f.弹宅迦林Dandakaranya.SetParent(f)

	f.羯陵伽Kalinga = kalinga.DKalinga羯陵伽
	f.羯陵伽Kalinga.SetParent(f)

	f.睹舍离Tosali = tosali.DTosali睹舍离
	f.睹舍离Tosali.SetParent(f)

}
