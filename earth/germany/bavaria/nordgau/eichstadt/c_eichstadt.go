package eichstadt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type EichstadtCounty interface {
	feud.County
	BAbenberg阿本贝格() feud.Barony
	BEichstadt艾希施泰特() feud.Barony
	BHirschberg希尔施贝格() feud.Barony
	BLechtgemund莱希特格明德() feud.Barony
	BPollenfeld波伦费尔德() feud.Barony
	BTrubendingen特鲁本丁根() feud.Barony
	BWeissenburg魏森堡() feud.Barony
}

type 艾希施泰特EichstadtCounty struct {
	feud.BaseCounty
	阿本贝格Abenberg      feud.Barony
	艾希施泰特Eichstadt    feud.Barony
	希尔施贝格Hirschberg   feud.Barony
	莱希特格明德Lechtgemund feud.Barony
	波伦费尔德Pollenfeld   feud.Barony
	特鲁本丁根Trubendingen feud.Barony
	魏森堡Weissenburg    feud.Barony
}

func (c *艾希施泰特EichstadtCounty) BAbenberg阿本贝格() feud.Barony {
	return c.阿本贝格Abenberg
}

func (c *艾希施泰特EichstadtCounty) BEichstadt艾希施泰特() feud.Barony {
	return c.艾希施泰特Eichstadt
}

func (c *艾希施泰特EichstadtCounty) BHirschberg希尔施贝格() feud.Barony {
	return c.希尔施贝格Hirschberg
}

func (c *艾希施泰特EichstadtCounty) BLechtgemund莱希特格明德() feud.Barony {
	return c.莱希特格明德Lechtgemund
}

func (c *艾希施泰特EichstadtCounty) BPollenfeld波伦费尔德() feud.Barony {
	return c.波伦费尔德Pollenfeld
}

func (c *艾希施泰特EichstadtCounty) BTrubendingen特鲁本丁根() feud.Barony {
	return c.特鲁本丁根Trubendingen
}

func (c *艾希施泰特EichstadtCounty) BWeissenburg魏森堡() feud.Barony {
	return c.魏森堡Weissenburg
}

var CEichstadt艾希施泰特 EichstadtCounty = &艾希施泰特EichstadtCounty{}

func init() {
	f := CEichstadt艾希施泰特.(*艾希施泰特EichstadtCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1985",
		Title:     "eichstadt",
		TitleName: "艾希施泰特",
		TitleCode: "c_eichstadt",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿本贝格Abenberg = BAbenberg阿本贝格
	f.阿本贝格Abenberg.SetParent(f)

	f.艾希施泰特Eichstadt = BEichstadt艾希施泰特
	f.艾希施泰特Eichstadt.SetParent(f)

	f.希尔施贝格Hirschberg = BHirschberg希尔施贝格
	f.希尔施贝格Hirschberg.SetParent(f)

	f.莱希特格明德Lechtgemund = BLechtgemund莱希特格明德
	f.莱希特格明德Lechtgemund.SetParent(f)

	f.波伦费尔德Pollenfeld = BPollenfeld波伦费尔德
	f.波伦费尔德Pollenfeld.SetParent(f)

	f.特鲁本丁根Trubendingen = BTrubendingen特鲁本丁根
	f.特鲁本丁根Trubendingen.SetParent(f)

	f.魏森堡Weissenburg = BWeissenburg魏森堡
	f.魏森堡Weissenburg.SetParent(f)

}
