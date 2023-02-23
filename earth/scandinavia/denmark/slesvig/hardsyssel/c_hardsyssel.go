package hardsyssel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HardsysselCounty interface {
	feud.County
	BHolstebro霍尔斯特布罗() feud.Barony
	BLemvig莱姆维() feud.Barony
	BRingkobing灵克宾() feud.Barony
	BSkive斯基沃() feud.Barony
	BSkjern斯凯恩() feud.Barony
	BStruer斯楚厄() feud.Barony
	BUlburg乌尔堡() feud.Barony
	BVarde瓦德() feud.Barony
}

type 瓦德HardsysselCounty struct {
	feud.BaseCounty
	霍尔斯特布罗Holstebro feud.Barony
	莱姆维Lemvig       feud.Barony
	灵克宾Ringkobing   feud.Barony
	斯基沃Skive        feud.Barony
	斯凯恩Skjern       feud.Barony
	斯楚厄Struer       feud.Barony
	乌尔堡Ulburg       feud.Barony
	瓦德Varde         feud.Barony
}

func (c *瓦德HardsysselCounty) BHolstebro霍尔斯特布罗() feud.Barony {
	return c.霍尔斯特布罗Holstebro
}

func (c *瓦德HardsysselCounty) BLemvig莱姆维() feud.Barony {
	return c.莱姆维Lemvig
}

func (c *瓦德HardsysselCounty) BRingkobing灵克宾() feud.Barony {
	return c.灵克宾Ringkobing
}

func (c *瓦德HardsysselCounty) BSkive斯基沃() feud.Barony {
	return c.斯基沃Skive
}

func (c *瓦德HardsysselCounty) BSkjern斯凯恩() feud.Barony {
	return c.斯凯恩Skjern
}

func (c *瓦德HardsysselCounty) BStruer斯楚厄() feud.Barony {
	return c.斯楚厄Struer
}

func (c *瓦德HardsysselCounty) BUlburg乌尔堡() feud.Barony {
	return c.乌尔堡Ulburg
}

func (c *瓦德HardsysselCounty) BVarde瓦德() feud.Barony {
	return c.瓦德Varde
}

var CHardsyssel瓦德 HardsysselCounty = &瓦德HardsysselCounty{}

func init() {
	f := CHardsyssel瓦德.(*瓦德HardsysselCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1683",
		Title:     "hardsyssel",
		TitleName: "瓦德",
		TitleCode: "c_hardsyssel",
		Baronies:  map[string]feud.Barony{},
	}

	f.霍尔斯特布罗Holstebro = BHolstebro霍尔斯特布罗
	f.霍尔斯特布罗Holstebro.SetParent(f)

	f.莱姆维Lemvig = BLemvig莱姆维
	f.莱姆维Lemvig.SetParent(f)

	f.灵克宾Ringkobing = BRingkobing灵克宾
	f.灵克宾Ringkobing.SetParent(f)

	f.斯基沃Skive = BSkive斯基沃
	f.斯基沃Skive.SetParent(f)

	f.斯凯恩Skjern = BSkjern斯凯恩
	f.斯凯恩Skjern.SetParent(f)

	f.斯楚厄Struer = BStruer斯楚厄
	f.斯楚厄Struer.SetParent(f)

	f.乌尔堡Ulburg = BUlburg乌尔堡
	f.乌尔堡Ulburg.SetParent(f)

	f.瓦德Varde = BVarde瓦德
	f.瓦德Varde.SetParent(f)

}
