package tibet

import (
	"github.com/thalesfu/CK2Commands/earth/tibet/guge"
	"github.com/thalesfu/CK2Commands/earth/tibet/kashmir"
	"github.com/thalesfu/CK2Commands/earth/tibet/kham"
	"github.com/thalesfu/CK2Commands/earth/tibet/nepal"
	"github.com/thalesfu/CK2Commands/earth/tibet/xixia"
	"github.com/thalesfu/CK2Commands/earth/tibet/yarlung"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TibetEmpire interface {
	feud.Empire
	KGuge古格() guge.GugeKingdom
	KKashmir迦湿弥罗() kashmir.KashmirKingdom
	KKham康区() kham.KhamKingdom
	KNepal尼波罗() nepal.NepalKingdom
	KXixia西夏() xixia.XixiaKingdom
	KYarlung卫藏() yarlung.YarlungKingdom
}

type 吐蕃TibetEmpire struct {
	feud.BaseEmpire
	古格Guge      guge.GugeKingdom
	迦湿弥罗Kashmir kashmir.KashmirKingdom
	康区Kham      kham.KhamKingdom
	尼波罗Nepal    nepal.NepalKingdom
	西夏Xixia     xixia.XixiaKingdom
	卫藏Yarlung   yarlung.YarlungKingdom
}

func (e *吐蕃TibetEmpire) KGuge古格() guge.GugeKingdom {
	return e.古格Guge
}

func (e *吐蕃TibetEmpire) KKashmir迦湿弥罗() kashmir.KashmirKingdom {
	return e.迦湿弥罗Kashmir
}

func (e *吐蕃TibetEmpire) KKham康区() kham.KhamKingdom {
	return e.康区Kham
}

func (e *吐蕃TibetEmpire) KNepal尼波罗() nepal.NepalKingdom {
	return e.尼波罗Nepal
}

func (e *吐蕃TibetEmpire) KXixia西夏() xixia.XixiaKingdom {
	return e.西夏Xixia
}

func (e *吐蕃TibetEmpire) KYarlung卫藏() yarlung.YarlungKingdom {
	return e.卫藏Yarlung
}

var ETibet吐蕃 TibetEmpire = &吐蕃TibetEmpire{}

func init() {
	f := ETibet吐蕃.(*吐蕃TibetEmpire)
	f.BaseEmpire = feud.BaseEmpire{
		Title:     "tibet",
		TitleName: "吐蕃",
		TitleCode: "e_tibet",
		Kingdoms:  map[string]feud.Kingdom{},
	}
	f.古格Guge = guge.KGuge古格
	f.古格Guge.SetParent(f)
	f.迦湿弥罗Kashmir = kashmir.KKashmir迦湿弥罗
	f.迦湿弥罗Kashmir.SetParent(f)
	f.康区Kham = kham.KKham康区
	f.康区Kham.SetParent(f)
	f.尼波罗Nepal = nepal.KNepal尼波罗
	f.尼波罗Nepal.SetParent(f)
	f.西夏Xixia = xixia.KXixia西夏
	f.西夏Xixia.SetParent(f)
	f.卫藏Yarlung = yarlung.KYarlung卫藏
	f.卫藏Yarlung.SetParent(f)
}
