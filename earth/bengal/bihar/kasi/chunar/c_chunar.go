package chunar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ChunarCounty interface {
	feud.County
	BBidadi毗荼提() feud.Barony
	BBidh毗陀() feud.Barony
	BBilikore毗梨拘梨() feud.Barony
	BChunar周那罗() feud.Barony
	BDaudpur达乌德布尔() feud.Barony
	BKantit建提底() feud.Barony
	BVindhyachal频陀遮罗() feud.Barony
}

type 枢那罗ChunarCounty struct {
	feud.BaseCounty
	毗荼提Bidadi       feud.Barony
	毗陀Bidh          feud.Barony
	毗梨拘梨Bilikore    feud.Barony
	周那罗Chunar       feud.Barony
	达乌德布尔Daudpur    feud.Barony
	建提底Kantit       feud.Barony
	频陀遮罗Vindhyachal feud.Barony
}

func (c *枢那罗ChunarCounty) BBidadi毗荼提() feud.Barony {
	return c.毗荼提Bidadi
}

func (c *枢那罗ChunarCounty) BBidh毗陀() feud.Barony {
	return c.毗陀Bidh
}

func (c *枢那罗ChunarCounty) BBilikore毗梨拘梨() feud.Barony {
	return c.毗梨拘梨Bilikore
}

func (c *枢那罗ChunarCounty) BChunar周那罗() feud.Barony {
	return c.周那罗Chunar
}

func (c *枢那罗ChunarCounty) BDaudpur达乌德布尔() feud.Barony {
	return c.达乌德布尔Daudpur
}

func (c *枢那罗ChunarCounty) BKantit建提底() feud.Barony {
	return c.建提底Kantit
}

func (c *枢那罗ChunarCounty) BVindhyachal频陀遮罗() feud.Barony {
	return c.频陀遮罗Vindhyachal
}

var CChunar枢那罗 ChunarCounty = &枢那罗ChunarCounty{}

func init() {
	f := CChunar枢那罗.(*枢那罗ChunarCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1281",
		Title:     "chunar",
		TitleName: "枢那罗",
		TitleCode: "c_chunar",
		Baronies:  map[string]feud.Barony{},
	}

	f.毗荼提Bidadi = BBidadi毗荼提
	f.毗荼提Bidadi.SetParent(f)

	f.毗陀Bidh = BBidh毗陀
	f.毗陀Bidh.SetParent(f)

	f.毗梨拘梨Bilikore = BBilikore毗梨拘梨
	f.毗梨拘梨Bilikore.SetParent(f)

	f.周那罗Chunar = BChunar周那罗
	f.周那罗Chunar.SetParent(f)

	f.达乌德布尔Daudpur = BDaudpur达乌德布尔
	f.达乌德布尔Daudpur.SetParent(f)

	f.建提底Kantit = BKantit建提底
	f.建提底Kantit.SetParent(f)

	f.频陀遮罗Vindhyachal = BVindhyachal频陀遮罗
	f.频陀遮罗Vindhyachal.SetParent(f)

}
