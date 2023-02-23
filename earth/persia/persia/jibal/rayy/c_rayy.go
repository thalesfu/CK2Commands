package rayy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RayyCounty interface {
	feud.County
	BEslamshahr埃斯兰夏尔() feud.Barony
	BHashtgerd哈什特盖尔德() feud.Barony
	BKaraj卡拉季() feud.Barony
	BPakdasht帕克达什特() feud.Barony
	BRayy雷伊() feud.Barony
	BRoudehen鲁代亨() feud.Barony
	BShahriar沙赫里亚尔() feud.Barony
	BTehran德黑兰() feud.Barony
}

type 雷伊RayyCounty struct {
	feud.BaseCounty
	埃斯兰夏尔Eslamshahr feud.Barony
	哈什特盖尔德Hashtgerd feud.Barony
	卡拉季Karaj        feud.Barony
	帕克达什特Pakdasht   feud.Barony
	雷伊Rayy          feud.Barony
	鲁代亨Roudehen     feud.Barony
	沙赫里亚尔Shahriar   feud.Barony
	德黑兰Tehran       feud.Barony
}

func (c *雷伊RayyCounty) BEslamshahr埃斯兰夏尔() feud.Barony {
	return c.埃斯兰夏尔Eslamshahr
}

func (c *雷伊RayyCounty) BHashtgerd哈什特盖尔德() feud.Barony {
	return c.哈什特盖尔德Hashtgerd
}

func (c *雷伊RayyCounty) BKaraj卡拉季() feud.Barony {
	return c.卡拉季Karaj
}

func (c *雷伊RayyCounty) BPakdasht帕克达什特() feud.Barony {
	return c.帕克达什特Pakdasht
}

func (c *雷伊RayyCounty) BRayy雷伊() feud.Barony {
	return c.雷伊Rayy
}

func (c *雷伊RayyCounty) BRoudehen鲁代亨() feud.Barony {
	return c.鲁代亨Roudehen
}

func (c *雷伊RayyCounty) BShahriar沙赫里亚尔() feud.Barony {
	return c.沙赫里亚尔Shahriar
}

func (c *雷伊RayyCounty) BTehran德黑兰() feud.Barony {
	return c.德黑兰Tehran
}

var CRayy雷伊 RayyCounty = &雷伊RayyCounty{}

func init() {
	f := CRayy雷伊.(*雷伊RayyCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "663",
		Title:     "rayy",
		TitleName: "雷伊",
		TitleCode: "c_rayy",
		Baronies:  map[string]feud.Barony{},
	}

	f.埃斯兰夏尔Eslamshahr = BEslamshahr埃斯兰夏尔
	f.埃斯兰夏尔Eslamshahr.SetParent(f)

	f.哈什特盖尔德Hashtgerd = BHashtgerd哈什特盖尔德
	f.哈什特盖尔德Hashtgerd.SetParent(f)

	f.卡拉季Karaj = BKaraj卡拉季
	f.卡拉季Karaj.SetParent(f)

	f.帕克达什特Pakdasht = BPakdasht帕克达什特
	f.帕克达什特Pakdasht.SetParent(f)

	f.雷伊Rayy = BRayy雷伊
	f.雷伊Rayy.SetParent(f)

	f.鲁代亨Roudehen = BRoudehen鲁代亨
	f.鲁代亨Roudehen.SetParent(f)

	f.沙赫里亚尔Shahriar = BShahriar沙赫里亚尔
	f.沙赫里亚尔Shahriar.SetParent(f)

	f.德黑兰Tehran = BTehran德黑兰
	f.德黑兰Tehran.SetParent(f)

}
