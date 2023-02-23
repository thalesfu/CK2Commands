package vizagipatam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VizagipatamCounty interface {
	feud.County
	BAnakapalle阿那迦波罗勒() feud.Barony
	BBavikonda婆毗军荼() feud.Barony
	BBheemunipatnam吠夷牟尼跋南() feud.Barony
	BKhuig丘俱() feud.Barony
	BKotipalli拘胝波梨() feud.Barony
	BSimhachalam僧诃遮蓝() feud.Barony
	BVizagipatam毗舍佉波吒南() feud.Barony
}

type 毗舍佉波吒南VizagipatamCounty struct {
	feud.BaseCounty
	阿那迦波罗勒Anakapalle     feud.Barony
	婆毗军荼Bavikonda        feud.Barony
	吠夷牟尼跋南Bheemunipatnam feud.Barony
	丘俱Khuig              feud.Barony
	拘胝波梨Kotipalli        feud.Barony
	僧诃遮蓝Simhachalam      feud.Barony
	毗舍佉波吒南Vizagipatam    feud.Barony
}

func (c *毗舍佉波吒南VizagipatamCounty) BAnakapalle阿那迦波罗勒() feud.Barony {
	return c.阿那迦波罗勒Anakapalle
}

func (c *毗舍佉波吒南VizagipatamCounty) BBavikonda婆毗军荼() feud.Barony {
	return c.婆毗军荼Bavikonda
}

func (c *毗舍佉波吒南VizagipatamCounty) BBheemunipatnam吠夷牟尼跋南() feud.Barony {
	return c.吠夷牟尼跋南Bheemunipatnam
}

func (c *毗舍佉波吒南VizagipatamCounty) BKhuig丘俱() feud.Barony {
	return c.丘俱Khuig
}

func (c *毗舍佉波吒南VizagipatamCounty) BKotipalli拘胝波梨() feud.Barony {
	return c.拘胝波梨Kotipalli
}

func (c *毗舍佉波吒南VizagipatamCounty) BSimhachalam僧诃遮蓝() feud.Barony {
	return c.僧诃遮蓝Simhachalam
}

func (c *毗舍佉波吒南VizagipatamCounty) BVizagipatam毗舍佉波吒南() feud.Barony {
	return c.毗舍佉波吒南Vizagipatam
}

var CVizagipatam毗舍佉波吒南 VizagipatamCounty = &毗舍佉波吒南VizagipatamCounty{}

func init() {
	f := CVizagipatam毗舍佉波吒南.(*毗舍佉波吒南VizagipatamCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1128",
		Title:     "vizagipatam",
		TitleName: "毗舍佉波吒南",
		TitleCode: "c_vizagipatam",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿那迦波罗勒Anakapalle = BAnakapalle阿那迦波罗勒
	f.阿那迦波罗勒Anakapalle.SetParent(f)

	f.婆毗军荼Bavikonda = BBavikonda婆毗军荼
	f.婆毗军荼Bavikonda.SetParent(f)

	f.吠夷牟尼跋南Bheemunipatnam = BBheemunipatnam吠夷牟尼跋南
	f.吠夷牟尼跋南Bheemunipatnam.SetParent(f)

	f.丘俱Khuig = BKhuig丘俱
	f.丘俱Khuig.SetParent(f)

	f.拘胝波梨Kotipalli = BKotipalli拘胝波梨
	f.拘胝波梨Kotipalli.SetParent(f)

	f.僧诃遮蓝Simhachalam = BSimhachalam僧诃遮蓝
	f.僧诃遮蓝Simhachalam.SetParent(f)

	f.毗舍佉波吒南Vizagipatam = BVizagipatam毗舍佉波吒南
	f.毗舍佉波吒南Vizagipatam.SetParent(f)

}
