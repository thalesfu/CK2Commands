package essex

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type EssexCounty interface {
	feud.County
	BBarking巴金() feud.Barony
	BColchester科尔切斯特() feud.Barony
	BDunmow邓莫() feud.Barony
	BHavering黑弗灵() feud.Barony
	BHedingham赫丁厄姆() feud.Barony
	BMaldon莫尔登() feud.Barony
	BPleshey布莱希() feud.Barony
	BWaltham沃尔瑟姆() feud.Barony
}

type 埃塞克斯EssexCounty struct {
	feud.BaseCounty
	巴金Barking       feud.Barony
	科尔切斯特Colchester feud.Barony
	邓莫Dunmow        feud.Barony
	黑弗灵Havering     feud.Barony
	赫丁厄姆Hedingham   feud.Barony
	莫尔登Maldon       feud.Barony
	布莱希Pleshey      feud.Barony
	沃尔瑟姆Waltham     feud.Barony
}

func (c *埃塞克斯EssexCounty) BBarking巴金() feud.Barony {
	return c.巴金Barking
}

func (c *埃塞克斯EssexCounty) BColchester科尔切斯特() feud.Barony {
	return c.科尔切斯特Colchester
}

func (c *埃塞克斯EssexCounty) BDunmow邓莫() feud.Barony {
	return c.邓莫Dunmow
}

func (c *埃塞克斯EssexCounty) BHavering黑弗灵() feud.Barony {
	return c.黑弗灵Havering
}

func (c *埃塞克斯EssexCounty) BHedingham赫丁厄姆() feud.Barony {
	return c.赫丁厄姆Hedingham
}

func (c *埃塞克斯EssexCounty) BMaldon莫尔登() feud.Barony {
	return c.莫尔登Maldon
}

func (c *埃塞克斯EssexCounty) BPleshey布莱希() feud.Barony {
	return c.布莱希Pleshey
}

func (c *埃塞克斯EssexCounty) BWaltham沃尔瑟姆() feud.Barony {
	return c.沃尔瑟姆Waltham
}

var CEssex埃塞克斯 EssexCounty = &埃塞克斯EssexCounty{}

func init() {
	f := CEssex埃塞克斯.(*埃塞克斯EssexCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "72",
		Title:     "essex",
		TitleName: "埃塞克斯",
		TitleCode: "c_essex",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴金Barking = BBarking巴金
	f.巴金Barking.SetParent(f)

	f.科尔切斯特Colchester = BColchester科尔切斯特
	f.科尔切斯特Colchester.SetParent(f)

	f.邓莫Dunmow = BDunmow邓莫
	f.邓莫Dunmow.SetParent(f)

	f.黑弗灵Havering = BHavering黑弗灵
	f.黑弗灵Havering.SetParent(f)

	f.赫丁厄姆Hedingham = BHedingham赫丁厄姆
	f.赫丁厄姆Hedingham.SetParent(f)

	f.莫尔登Maldon = BMaldon莫尔登
	f.莫尔登Maldon.SetParent(f)

	f.布莱希Pleshey = BPleshey布莱希
	f.布莱希Pleshey.SetParent(f)

	f.沃尔瑟姆Waltham = BWaltham沃尔瑟姆
	f.沃尔瑟姆Waltham.SetParent(f)

}
