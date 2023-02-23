package france

import (
	"github.com/thalesfu/CK2Commands/earth/france/france/anjou"
	"github.com/thalesfu/CK2Commands/earth/france/france/berry"
	"github.com/thalesfu/CK2Commands/earth/france/france/burgundy"
	"github.com/thalesfu/CK2Commands/earth/france/france/champagne"
	"github.com/thalesfu/CK2Commands/earth/france/france/flanders"
	"github.com/thalesfu/CK2Commands/earth/france/france/normandy"
	"github.com/thalesfu/CK2Commands/earth/france/france/orleans"
	"github.com/thalesfu/CK2Commands/earth/france/france/picardie"
	"github.com/thalesfu/CK2Commands/earth/france/france/valois"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FranceKingdom interface {
	feud.Kingdom
	DAnjou安茹() anjou.AnjouDuke
	DBerry贝里() berry.BerryDuke
	DBurgundy勃艮第() burgundy.BurgundyDuke
	DChampagne香槟() champagne.ChampagneDuke
	DFlanders佛兰德斯() flanders.FlandersDuke
	DNormandy诺曼底() normandy.NormandyDuke
	DOrleans布卢瓦() orleans.OrleansDuke
	DPicardie皮卡第() picardie.PicardieDuke
	DValois巴黎() valois.ValoisDuke
}

type 法兰西FranceKingdom struct {
	feud.BaseKingdom
	安茹Anjou      anjou.AnjouDuke
	贝里Berry      berry.BerryDuke
	勃艮第Burgundy  burgundy.BurgundyDuke
	香槟Champagne  champagne.ChampagneDuke
	佛兰德斯Flanders flanders.FlandersDuke
	诺曼底Normandy  normandy.NormandyDuke
	布卢瓦Orleans   orleans.OrleansDuke
	皮卡第Picardie  picardie.PicardieDuke
	巴黎Valois     valois.ValoisDuke
}

func (k *法兰西FranceKingdom) DAnjou安茹() anjou.AnjouDuke {
	return k.安茹Anjou
}

func (k *法兰西FranceKingdom) DBerry贝里() berry.BerryDuke {
	return k.贝里Berry
}

func (k *法兰西FranceKingdom) DBurgundy勃艮第() burgundy.BurgundyDuke {
	return k.勃艮第Burgundy
}

func (k *法兰西FranceKingdom) DChampagne香槟() champagne.ChampagneDuke {
	return k.香槟Champagne
}

func (k *法兰西FranceKingdom) DFlanders佛兰德斯() flanders.FlandersDuke {
	return k.佛兰德斯Flanders
}

func (k *法兰西FranceKingdom) DNormandy诺曼底() normandy.NormandyDuke {
	return k.诺曼底Normandy
}

func (k *法兰西FranceKingdom) DOrleans布卢瓦() orleans.OrleansDuke {
	return k.布卢瓦Orleans
}

func (k *法兰西FranceKingdom) DPicardie皮卡第() picardie.PicardieDuke {
	return k.皮卡第Picardie
}

func (k *法兰西FranceKingdom) DValois巴黎() valois.ValoisDuke {
	return k.巴黎Valois
}

var KFrance法兰西 FranceKingdom = &法兰西FranceKingdom{}

func init() {
	f := KFrance法兰西.(*法兰西FranceKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "france",
		TitleName: "法兰西",
		TitleCode: "k_france",
		Dukes:     map[string]feud.Duke{},
	}

	f.安茹Anjou = anjou.DAnjou安茹
	f.安茹Anjou.SetParent(f)

	f.贝里Berry = berry.DBerry贝里
	f.贝里Berry.SetParent(f)

	f.勃艮第Burgundy = burgundy.DBurgundy勃艮第
	f.勃艮第Burgundy.SetParent(f)

	f.香槟Champagne = champagne.DChampagne香槟
	f.香槟Champagne.SetParent(f)

	f.佛兰德斯Flanders = flanders.DFlanders佛兰德斯
	f.佛兰德斯Flanders.SetParent(f)

	f.诺曼底Normandy = normandy.DNormandy诺曼底
	f.诺曼底Normandy.SetParent(f)

	f.布卢瓦Orleans = orleans.DOrleans布卢瓦
	f.布卢瓦Orleans.SetParent(f)

	f.皮卡第Picardie = picardie.DPicardie皮卡第
	f.皮卡第Picardie.SetParent(f)

	f.巴黎Valois = valois.DValois巴黎
	f.巴黎Valois.SetParent(f)

}
