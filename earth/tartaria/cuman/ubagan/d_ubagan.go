package ubagan

import (
	"github.com/thalesfu/CK2Commands/earth/tartaria/cuman/ubagan/kartaly"
	"github.com/thalesfu/CK2Commands/earth/tartaria/cuman/ubagan/ubagan"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type UbaganDuke interface {
	feud.Duke
	CKartaly卡尔塔雷_阿亚特() kartaly.KartalyCounty
	CUbagan乌巴甘() ubagan.UbaganCounty
}

type 乌巴甘UbaganDuke struct {
	feud.BaseDuke
	卡尔塔雷_阿亚特Kartaly kartaly.KartalyCounty
	乌巴甘Ubagan       ubagan.UbaganCounty
}

func (d *乌巴甘UbaganDuke) CKartaly卡尔塔雷_阿亚特() kartaly.KartalyCounty {
	return d.卡尔塔雷_阿亚特Kartaly
}

func (d *乌巴甘UbaganDuke) CUbagan乌巴甘() ubagan.UbaganCounty {
	return d.乌巴甘Ubagan
}

var DUbagan乌巴甘 UbaganDuke = &乌巴甘UbaganDuke{}

func init() {
	f := DUbagan乌巴甘.(*乌巴甘UbaganDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "ubagan",
		TitleName: "乌巴甘",
		TitleCode: "d_ubagan",
		Counties:  map[string]feud.County{},
	}

	f.卡尔塔雷_阿亚特Kartaly = kartaly.CKartaly卡尔塔雷_阿亚特
	f.卡尔塔雷_阿亚特Kartaly.SetParent(f)

	f.乌巴甘Ubagan = ubagan.CUbagan乌巴甘
	f.乌巴甘Ubagan.SetParent(f)

}
