package bavaria

import (
	"github.com/thalesfu/CK2Commands/earth/germany/bavaria/bavaria/niederbayern"
	"github.com/thalesfu/CK2Commands/earth/germany/bavaria/bavaria/oberbayern"
	"github.com/thalesfu/CK2Commands/earth/germany/bavaria/bavaria/passau"
	"github.com/thalesfu/CK2Commands/earth/germany/bavaria/bavaria/regensburg"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BavariaDuke interface {
	feud.Duke
	CNiederbayern卡姆() niederbayern.NiederbayernCounty
	COberbayern弗赖辛() oberbayern.OberbayernCounty
	CPassau帕绍() passau.PassauCounty
	CRegensburg雷根斯堡() regensburg.RegensburgCounty
}

type 巴伐利亚BavariaDuke struct {
	feud.BaseDuke
	卡姆Niederbayern niederbayern.NiederbayernCounty
	弗赖辛Oberbayern  oberbayern.OberbayernCounty
	帕绍Passau       passau.PassauCounty
	雷根斯堡Regensburg regensburg.RegensburgCounty
}

func (d *巴伐利亚BavariaDuke) CNiederbayern卡姆() niederbayern.NiederbayernCounty {
	return d.卡姆Niederbayern
}

func (d *巴伐利亚BavariaDuke) COberbayern弗赖辛() oberbayern.OberbayernCounty {
	return d.弗赖辛Oberbayern
}

func (d *巴伐利亚BavariaDuke) CPassau帕绍() passau.PassauCounty {
	return d.帕绍Passau
}

func (d *巴伐利亚BavariaDuke) CRegensburg雷根斯堡() regensburg.RegensburgCounty {
	return d.雷根斯堡Regensburg
}

var DBavaria巴伐利亚 BavariaDuke = &巴伐利亚BavariaDuke{}

func init() {
	f := DBavaria巴伐利亚.(*巴伐利亚BavariaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "bavaria",
		TitleName: "巴伐利亚",
		TitleCode: "d_bavaria",
		Counties:  map[string]feud.County{},
	}

	f.卡姆Niederbayern = niederbayern.CNiederbayern卡姆
	f.卡姆Niederbayern.SetParent(f)

	f.弗赖辛Oberbayern = oberbayern.COberbayern弗赖辛
	f.弗赖辛Oberbayern.SetParent(f)

	f.帕绍Passau = passau.CPassau帕绍
	f.帕绍Passau.SetParent(f)

	f.雷根斯堡Regensburg = regensburg.CRegensburg雷根斯堡
	f.雷根斯堡Regensburg.SetParent(f)

}
