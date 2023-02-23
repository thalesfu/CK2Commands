package rottenburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RottenburgCounty interface {
	feud.County
	BHohlach霍拉赫() feud.Barony
	BOchsenfurt奥克森富特() feud.Barony
	BRottenburg罗滕堡() feud.Barony
	BUffenheim乌芬海姆() feud.Barony
	BWeikersheim魏克斯海姆() feud.Barony
	BWeilerburg魏勒堡() feud.Barony
	BWeinsberg魏恩斯贝格() feud.Barony
}

type 魏恩斯贝格RottenburgCounty struct {
	feud.BaseCounty
	霍拉赫Hohlach       feud.Barony
	奥克森富特Ochsenfurt  feud.Barony
	罗滕堡Rottenburg    feud.Barony
	乌芬海姆Uffenheim    feud.Barony
	魏克斯海姆Weikersheim feud.Barony
	魏勒堡Weilerburg    feud.Barony
	魏恩斯贝格Weinsberg   feud.Barony
}

func (c *魏恩斯贝格RottenburgCounty) BHohlach霍拉赫() feud.Barony {
	return c.霍拉赫Hohlach
}

func (c *魏恩斯贝格RottenburgCounty) BOchsenfurt奥克森富特() feud.Barony {
	return c.奥克森富特Ochsenfurt
}

func (c *魏恩斯贝格RottenburgCounty) BRottenburg罗滕堡() feud.Barony {
	return c.罗滕堡Rottenburg
}

func (c *魏恩斯贝格RottenburgCounty) BUffenheim乌芬海姆() feud.Barony {
	return c.乌芬海姆Uffenheim
}

func (c *魏恩斯贝格RottenburgCounty) BWeikersheim魏克斯海姆() feud.Barony {
	return c.魏克斯海姆Weikersheim
}

func (c *魏恩斯贝格RottenburgCounty) BWeilerburg魏勒堡() feud.Barony {
	return c.魏勒堡Weilerburg
}

func (c *魏恩斯贝格RottenburgCounty) BWeinsberg魏恩斯贝格() feud.Barony {
	return c.魏恩斯贝格Weinsberg
}

var CRottenburg魏恩斯贝格 RottenburgCounty = &魏恩斯贝格RottenburgCounty{}

func init() {
	f := CRottenburg魏恩斯贝格.(*魏恩斯贝格RottenburgCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1984",
		Title:     "rottenburg",
		TitleName: "魏恩斯贝格",
		TitleCode: "c_rottenburg",
		Baronies:  map[string]feud.Barony{},
	}

	f.霍拉赫Hohlach = BHohlach霍拉赫
	f.霍拉赫Hohlach.SetParent(f)

	f.奥克森富特Ochsenfurt = BOchsenfurt奥克森富特
	f.奥克森富特Ochsenfurt.SetParent(f)

	f.罗滕堡Rottenburg = BRottenburg罗滕堡
	f.罗滕堡Rottenburg.SetParent(f)

	f.乌芬海姆Uffenheim = BUffenheim乌芬海姆
	f.乌芬海姆Uffenheim.SetParent(f)

	f.魏克斯海姆Weikersheim = BWeikersheim魏克斯海姆
	f.魏克斯海姆Weikersheim.SetParent(f)

	f.魏勒堡Weilerburg = BWeilerburg魏勒堡
	f.魏勒堡Weilerburg.SetParent(f)

	f.魏恩斯贝格Weinsberg = BWeinsberg魏恩斯贝格
	f.魏恩斯贝格Weinsberg.SetParent(f)

}
