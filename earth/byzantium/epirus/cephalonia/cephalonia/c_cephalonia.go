package cephalonia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CephaloniaCounty interface {
	feud.County
	BAtros阿特罗斯() feud.Barony
	BCerigo切里戈() feud.Barony
	BIthaca伊萨卡() feud.Barony
	BKefalonia凯法利尼亚() feud.Barony
	BLefkas莱夫卡斯() feud.Barony
	BPalaiofrourio帕兰法里兰() feud.Barony
	BPaxos帕克索斯() feud.Barony
}

type 凯法洛尼亚CephaloniaCounty struct {
	feud.BaseCounty
	阿特罗斯Atros          feud.Barony
	切里戈Cerigo          feud.Barony
	伊萨卡Ithaca          feud.Barony
	凯法利尼亚Kefalonia     feud.Barony
	莱夫卡斯Lefkas         feud.Barony
	帕兰法里兰Palaiofrourio feud.Barony
	帕克索斯Paxos          feud.Barony
}

func (c *凯法洛尼亚CephaloniaCounty) BAtros阿特罗斯() feud.Barony {
	return c.阿特罗斯Atros
}

func (c *凯法洛尼亚CephaloniaCounty) BCerigo切里戈() feud.Barony {
	return c.切里戈Cerigo
}

func (c *凯法洛尼亚CephaloniaCounty) BIthaca伊萨卡() feud.Barony {
	return c.伊萨卡Ithaca
}

func (c *凯法洛尼亚CephaloniaCounty) BKefalonia凯法利尼亚() feud.Barony {
	return c.凯法利尼亚Kefalonia
}

func (c *凯法洛尼亚CephaloniaCounty) BLefkas莱夫卡斯() feud.Barony {
	return c.莱夫卡斯Lefkas
}

func (c *凯法洛尼亚CephaloniaCounty) BPalaiofrourio帕兰法里兰() feud.Barony {
	return c.帕兰法里兰Palaiofrourio
}

func (c *凯法洛尼亚CephaloniaCounty) BPaxos帕克索斯() feud.Barony {
	return c.帕克索斯Paxos
}

var CCephalonia凯法洛尼亚 CephaloniaCounty = &凯法洛尼亚CephaloniaCounty{}

func init() {
	f := CCephalonia凯法洛尼亚.(*凯法洛尼亚CephaloniaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "474",
		Title:     "cephalonia",
		TitleName: "凯法洛尼亚",
		TitleCode: "c_cephalonia",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿特罗斯Atros = BAtros阿特罗斯
	f.阿特罗斯Atros.SetParent(f)

	f.切里戈Cerigo = BCerigo切里戈
	f.切里戈Cerigo.SetParent(f)

	f.伊萨卡Ithaca = BIthaca伊萨卡
	f.伊萨卡Ithaca.SetParent(f)

	f.凯法利尼亚Kefalonia = BKefalonia凯法利尼亚
	f.凯法利尼亚Kefalonia.SetParent(f)

	f.莱夫卡斯Lefkas = BLefkas莱夫卡斯
	f.莱夫卡斯Lefkas.SetParent(f)

	f.帕兰法里兰Palaiofrourio = BPalaiofrourio帕兰法里兰
	f.帕兰法里兰Palaiofrourio.SetParent(f)

	f.帕克索斯Paxos = BPaxos帕克索斯
	f.帕克索斯Paxos.SetParent(f)

}
