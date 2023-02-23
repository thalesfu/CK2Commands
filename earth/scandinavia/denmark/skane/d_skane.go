package skane

import (
	"github.com/thalesfu/CK2Commands/earth/scandinavia/denmark/skane/blekinge"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/denmark/skane/bornholm"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/denmark/skane/halland"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/denmark/skane/skane"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SkaneDuke interface {
	feud.Duke
	CBlekinge布莱金厄() blekinge.BlekingeCounty
	CBornholm博恩霍尔姆() bornholm.BornholmCounty
	CHalland哈兰() halland.HallandCounty
	CSkane斯堪尼亚() skane.SkaneCounty
}

type 斯堪尼亚SkaneDuke struct {
	feud.BaseDuke
	布莱金厄Blekinge  blekinge.BlekingeCounty
	博恩霍尔姆Bornholm bornholm.BornholmCounty
	哈兰Halland     halland.HallandCounty
	斯堪尼亚Skane     skane.SkaneCounty
}

func (d *斯堪尼亚SkaneDuke) CBlekinge布莱金厄() blekinge.BlekingeCounty {
	return d.布莱金厄Blekinge
}

func (d *斯堪尼亚SkaneDuke) CBornholm博恩霍尔姆() bornholm.BornholmCounty {
	return d.博恩霍尔姆Bornholm
}

func (d *斯堪尼亚SkaneDuke) CHalland哈兰() halland.HallandCounty {
	return d.哈兰Halland
}

func (d *斯堪尼亚SkaneDuke) CSkane斯堪尼亚() skane.SkaneCounty {
	return d.斯堪尼亚Skane
}

var DSkane斯堪尼亚 SkaneDuke = &斯堪尼亚SkaneDuke{}

func init() {
	f := DSkane斯堪尼亚.(*斯堪尼亚SkaneDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "skane",
		TitleName: "斯堪尼亚",
		TitleCode: "d_skane",
		Counties:  map[string]feud.County{},
	}

	f.布莱金厄Blekinge = blekinge.CBlekinge布莱金厄
	f.布莱金厄Blekinge.SetParent(f)

	f.博恩霍尔姆Bornholm = bornholm.CBornholm博恩霍尔姆
	f.博恩霍尔姆Bornholm.SetParent(f)

	f.哈兰Halland = halland.CHalland哈兰
	f.哈兰Halland.SetParent(f)

	f.斯堪尼亚Skane = skane.CSkane斯堪尼亚
	f.斯堪尼亚Skane.SetParent(f)

}
