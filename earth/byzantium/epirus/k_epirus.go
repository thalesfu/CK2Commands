package epirus

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/epirus/cephalonia"
	"github.com/thalesfu/CK2Commands/earth/byzantium/epirus/dyrrachion"
	"github.com/thalesfu/CK2Commands/earth/byzantium/epirus/epirus"
	"github.com/thalesfu/CK2Commands/earth/byzantium/epirus/ohrid"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type EpirusKingdom interface {
	feud.Kingdom
	DCephalonia凯法洛尼亚() cephalonia.CephaloniaDuke
	DDyrrachion都拉齐翁() dyrrachion.DyrrachionDuke
	DEpirus伊庇鲁斯() epirus.EpirusDuke
	DOhrid奥赫里德() ohrid.OhridDuke
}

type 伊庇鲁斯EpirusKingdom struct {
	feud.BaseKingdom
	凯法洛尼亚Cephalonia cephalonia.CephaloniaDuke
	都拉齐翁Dyrrachion  dyrrachion.DyrrachionDuke
	伊庇鲁斯Epirus      epirus.EpirusDuke
	奥赫里德Ohrid       ohrid.OhridDuke
}

func (k *伊庇鲁斯EpirusKingdom) DCephalonia凯法洛尼亚() cephalonia.CephaloniaDuke {
	return k.凯法洛尼亚Cephalonia
}

func (k *伊庇鲁斯EpirusKingdom) DDyrrachion都拉齐翁() dyrrachion.DyrrachionDuke {
	return k.都拉齐翁Dyrrachion
}

func (k *伊庇鲁斯EpirusKingdom) DEpirus伊庇鲁斯() epirus.EpirusDuke {
	return k.伊庇鲁斯Epirus
}

func (k *伊庇鲁斯EpirusKingdom) DOhrid奥赫里德() ohrid.OhridDuke {
	return k.奥赫里德Ohrid
}

var KEpirus伊庇鲁斯 EpirusKingdom = &伊庇鲁斯EpirusKingdom{}

func init() {
	f := KEpirus伊庇鲁斯.(*伊庇鲁斯EpirusKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "epirus",
		TitleName: "伊庇鲁斯",
		TitleCode: "k_epirus",
		Dukes:     map[string]feud.Duke{},
	}

	f.凯法洛尼亚Cephalonia = cephalonia.DCephalonia凯法洛尼亚
	f.凯法洛尼亚Cephalonia.SetParent(f)

	f.都拉齐翁Dyrrachion = dyrrachion.DDyrrachion都拉齐翁
	f.都拉齐翁Dyrrachion.SetParent(f)

	f.伊庇鲁斯Epirus = epirus.DEpirus伊庇鲁斯
	f.伊庇鲁斯Epirus.SetParent(f)

	f.奥赫里德Ohrid = ohrid.DOhrid奥赫里德
	f.奥赫里德Ohrid.SetParent(f)

}
