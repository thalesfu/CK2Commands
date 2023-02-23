package asirgarh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AsirgarhCounty interface {
	feud.County
	BAsirgarh阿斯梨姞利呬() feud.Barony
	BHansla汉罗() feud.Barony
	BJakpur阇迦补罗() feud.Barony
	BKhandwa肯德瓦() feud.Barony
	BKhargone迦戈讷() feud.Barony
	BMandleshwar曼荼丽湿伐罗() feud.Barony
	BSanawad娑那跋陀() feud.Barony
}

type 阿斯梨姞利呬AsirgarhCounty struct {
	feud.BaseCounty
	阿斯梨姞利呬Asirgarh    feud.Barony
	汉罗Hansla          feud.Barony
	阇迦补罗Jakpur        feud.Barony
	肯德瓦Khandwa        feud.Barony
	迦戈讷Khargone       feud.Barony
	曼荼丽湿伐罗Mandleshwar feud.Barony
	娑那跋陀Sanawad       feud.Barony
}

func (c *阿斯梨姞利呬AsirgarhCounty) BAsirgarh阿斯梨姞利呬() feud.Barony {
	return c.阿斯梨姞利呬Asirgarh
}

func (c *阿斯梨姞利呬AsirgarhCounty) BHansla汉罗() feud.Barony {
	return c.汉罗Hansla
}

func (c *阿斯梨姞利呬AsirgarhCounty) BJakpur阇迦补罗() feud.Barony {
	return c.阇迦补罗Jakpur
}

func (c *阿斯梨姞利呬AsirgarhCounty) BKhandwa肯德瓦() feud.Barony {
	return c.肯德瓦Khandwa
}

func (c *阿斯梨姞利呬AsirgarhCounty) BKhargone迦戈讷() feud.Barony {
	return c.迦戈讷Khargone
}

func (c *阿斯梨姞利呬AsirgarhCounty) BMandleshwar曼荼丽湿伐罗() feud.Barony {
	return c.曼荼丽湿伐罗Mandleshwar
}

func (c *阿斯梨姞利呬AsirgarhCounty) BSanawad娑那跋陀() feud.Barony {
	return c.娑那跋陀Sanawad
}

var CAsirgarh阿斯梨姞利呬 AsirgarhCounty = &阿斯梨姞利呬AsirgarhCounty{}

func init() {
	f := CAsirgarh阿斯梨姞利呬.(*阿斯梨姞利呬AsirgarhCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1287",
		Title:     "asirgarh",
		TitleName: "阿斯梨姞利呬",
		TitleCode: "c_asirgarh",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿斯梨姞利呬Asirgarh = BAsirgarh阿斯梨姞利呬
	f.阿斯梨姞利呬Asirgarh.SetParent(f)

	f.汉罗Hansla = BHansla汉罗
	f.汉罗Hansla.SetParent(f)

	f.阇迦补罗Jakpur = BJakpur阇迦补罗
	f.阇迦补罗Jakpur.SetParent(f)

	f.肯德瓦Khandwa = BKhandwa肯德瓦
	f.肯德瓦Khandwa.SetParent(f)

	f.迦戈讷Khargone = BKhargone迦戈讷
	f.迦戈讷Khargone.SetParent(f)

	f.曼荼丽湿伐罗Mandleshwar = BMandleshwar曼荼丽湿伐罗
	f.曼荼丽湿伐罗Mandleshwar.SetParent(f)

	f.娑那跋陀Sanawad = BSanawad娑那跋陀
	f.娑那跋陀Sanawad.SetParent(f)

}
