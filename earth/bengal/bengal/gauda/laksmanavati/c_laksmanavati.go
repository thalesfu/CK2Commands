package laksmanavati

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LaksmanavatiCounty interface {
	feud.County
	BBanan班安() feud.Barony
	BGaur乔罗() feud.Barony
	BLaksmanavati罗叉摩那伐底() feud.Barony
	BMalda摩罗陀() feud.Barony
	BPankak盘卡克() feud.Barony
	BPuthia补提阿() feud.Barony
	BShukla叔罗() feud.Barony
}

type 罗叉摩那伐底LaksmanavatiCounty struct {
	feud.BaseCounty
	班安Banan            feud.Barony
	乔罗Gaur             feud.Barony
	罗叉摩那伐底Laksmanavati feud.Barony
	摩罗陀Malda           feud.Barony
	盘卡克Pankak          feud.Barony
	补提阿Puthia          feud.Barony
	叔罗Shukla           feud.Barony
}

func (c *罗叉摩那伐底LaksmanavatiCounty) BBanan班安() feud.Barony {
	return c.班安Banan
}

func (c *罗叉摩那伐底LaksmanavatiCounty) BGaur乔罗() feud.Barony {
	return c.乔罗Gaur
}

func (c *罗叉摩那伐底LaksmanavatiCounty) BLaksmanavati罗叉摩那伐底() feud.Barony {
	return c.罗叉摩那伐底Laksmanavati
}

func (c *罗叉摩那伐底LaksmanavatiCounty) BMalda摩罗陀() feud.Barony {
	return c.摩罗陀Malda
}

func (c *罗叉摩那伐底LaksmanavatiCounty) BPankak盘卡克() feud.Barony {
	return c.盘卡克Pankak
}

func (c *罗叉摩那伐底LaksmanavatiCounty) BPuthia补提阿() feud.Barony {
	return c.补提阿Puthia
}

func (c *罗叉摩那伐底LaksmanavatiCounty) BShukla叔罗() feud.Barony {
	return c.叔罗Shukla
}

var CLaksmanavati罗叉摩那伐底 LaksmanavatiCounty = &罗叉摩那伐底LaksmanavatiCounty{}

func init() {
	f := CLaksmanavati罗叉摩那伐底.(*罗叉摩那伐底LaksmanavatiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1151",
		Title:     "laksmanavati",
		TitleName: "罗叉摩那伐底",
		TitleCode: "c_laksmanavati",
		Baronies:  map[string]feud.Barony{},
	}

	f.班安Banan = BBanan班安
	f.班安Banan.SetParent(f)

	f.乔罗Gaur = BGaur乔罗
	f.乔罗Gaur.SetParent(f)

	f.罗叉摩那伐底Laksmanavati = BLaksmanavati罗叉摩那伐底
	f.罗叉摩那伐底Laksmanavati.SetParent(f)

	f.摩罗陀Malda = BMalda摩罗陀
	f.摩罗陀Malda.SetParent(f)

	f.盘卡克Pankak = BPankak盘卡克
	f.盘卡克Pankak.SetParent(f)

	f.补提阿Puthia = BPuthia补提阿
	f.补提阿Puthia.SetParent(f)

	f.叔罗Shukla = BShukla叔罗
	f.叔罗Shukla.SetParent(f)

}
