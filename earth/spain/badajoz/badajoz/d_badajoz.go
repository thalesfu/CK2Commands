package badajoz

import (
	"github.com/thalesfu/CK2Commands/earth/spain/badajoz/badajoz/alcantara"
	"github.com/thalesfu/CK2Commands/earth/spain/badajoz/badajoz/badajoz"
	"github.com/thalesfu/CK2Commands/earth/spain/badajoz/badajoz/caceres"
	"github.com/thalesfu/CK2Commands/earth/spain/badajoz/badajoz/plasencia"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BadajozDuke interface {
	feud.Duke
	CAlcantara阿尔坎塔拉() alcantara.AlcantaraCounty
	CBadajoz巴达霍斯() badajoz.BadajozCounty
	CCaceres卡塞雷斯() caceres.CaceresCounty
	CPlasencia普拉森西亚() plasencia.PlasenciaCounty
}

type 巴达霍斯BadajozDuke struct {
	feud.BaseDuke
	阿尔坎塔拉Alcantara alcantara.AlcantaraCounty
	巴达霍斯Badajoz    badajoz.BadajozCounty
	卡塞雷斯Caceres    caceres.CaceresCounty
	普拉森西亚Plasencia plasencia.PlasenciaCounty
}

func (d *巴达霍斯BadajozDuke) CAlcantara阿尔坎塔拉() alcantara.AlcantaraCounty {
	return d.阿尔坎塔拉Alcantara
}

func (d *巴达霍斯BadajozDuke) CBadajoz巴达霍斯() badajoz.BadajozCounty {
	return d.巴达霍斯Badajoz
}

func (d *巴达霍斯BadajozDuke) CCaceres卡塞雷斯() caceres.CaceresCounty {
	return d.卡塞雷斯Caceres
}

func (d *巴达霍斯BadajozDuke) CPlasencia普拉森西亚() plasencia.PlasenciaCounty {
	return d.普拉森西亚Plasencia
}

var DBadajoz巴达霍斯 BadajozDuke = &巴达霍斯BadajozDuke{}

func init() {
	f := DBadajoz巴达霍斯.(*巴达霍斯BadajozDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "badajoz",
		TitleName: "巴达霍斯",
		TitleCode: "d_badajoz",
		Counties:  map[string]feud.County{},
	}

	f.阿尔坎塔拉Alcantara = alcantara.CAlcantara阿尔坎塔拉
	f.阿尔坎塔拉Alcantara.SetParent(f)

	f.巴达霍斯Badajoz = badajoz.CBadajoz巴达霍斯
	f.巴达霍斯Badajoz.SetParent(f)

	f.卡塞雷斯Caceres = caceres.CCaceres卡塞雷斯
	f.卡塞雷斯Caceres.SetParent(f)

	f.普拉森西亚Plasencia = plasencia.CPlasencia普拉森西亚
	f.普拉森西亚Plasencia.SetParent(f)

}
