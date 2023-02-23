package thrace

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/thrace/thrace/byzantion"
	"github.com/thalesfu/CK2Commands/earth/byzantium/thrace/thrace/kaliopolis"
	"github.com/thalesfu/CK2Commands/earth/byzantium/thrace/thrace/thrake"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ThraceDuke interface {
	feud.Duke
	CByzantion君士坦丁堡() byzantion.ByzantionCounty
	CKaliopolis卡利奥波利斯() kaliopolis.KaliopolisCounty
	CThrake色雷斯() thrake.ThrakeCounty
}

type 色雷斯ThraceDuke struct {
	feud.BaseDuke
	君士坦丁堡Byzantion   byzantion.ByzantionCounty
	卡利奥波利斯Kaliopolis kaliopolis.KaliopolisCounty
	色雷斯Thrake        thrake.ThrakeCounty
}

func (d *色雷斯ThraceDuke) CByzantion君士坦丁堡() byzantion.ByzantionCounty {
	return d.君士坦丁堡Byzantion
}

func (d *色雷斯ThraceDuke) CKaliopolis卡利奥波利斯() kaliopolis.KaliopolisCounty {
	return d.卡利奥波利斯Kaliopolis
}

func (d *色雷斯ThraceDuke) CThrake色雷斯() thrake.ThrakeCounty {
	return d.色雷斯Thrake
}

var DThrace色雷斯 ThraceDuke = &色雷斯ThraceDuke{}

func init() {
	f := DThrace色雷斯.(*色雷斯ThraceDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "thrace",
		TitleName: "色雷斯",
		TitleCode: "d_thrace",
		Counties:  map[string]feud.County{},
	}

	f.君士坦丁堡Byzantion = byzantion.CByzantion君士坦丁堡
	f.君士坦丁堡Byzantion.SetParent(f)

	f.卡利奥波利斯Kaliopolis = kaliopolis.CKaliopolis卡利奥波利斯
	f.卡利奥波利斯Kaliopolis.SetParent(f)

	f.色雷斯Thrake = thrake.CThrake色雷斯
	f.色雷斯Thrake.SetParent(f)

}
