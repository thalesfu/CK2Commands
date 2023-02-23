package kanem

import (
	"github.com/thalesfu/CK2Commands/earth/kanem/hausaland"
	"github.com/thalesfu/CK2Commands/earth/kanem/kanem"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KanemEmpire interface {
	feud.Empire
	KHausaland豪萨() hausaland.HausalandKingdom
	KKanem加涅姆() kanem.KanemKingdom
}

type 加涅姆_博尔努KanemEmpire struct {
	feud.BaseEmpire
	豪萨Hausaland hausaland.HausalandKingdom
	加涅姆Kanem    kanem.KanemKingdom
}

func (e *加涅姆_博尔努KanemEmpire) KHausaland豪萨() hausaland.HausalandKingdom {
	return e.豪萨Hausaland
}

func (e *加涅姆_博尔努KanemEmpire) KKanem加涅姆() kanem.KanemKingdom {
	return e.加涅姆Kanem
}

var EKanem加涅姆_博尔努 KanemEmpire = &加涅姆_博尔努KanemEmpire{}

func init() {
	f := EKanem加涅姆_博尔努.(*加涅姆_博尔努KanemEmpire)
	f.BaseEmpire = feud.BaseEmpire{
		Title:     "kanem",
		TitleName: "加涅姆_博尔努",
		TitleCode: "e_kanem",
		Kingdoms:  map[string]feud.Kingdom{},
	}
	f.豪萨Hausaland = hausaland.KHausaland豪萨
	f.豪萨Hausaland.SetParent(f)
	f.加涅姆Kanem = kanem.KKanem加涅姆
	f.加涅姆Kanem.SetParent(f)
}
