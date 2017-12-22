package database

import (
	//	"time"

	"../../../share/model"
	"github.com/russross/meddler"

	//	log "github.com/Sirupsen/logrus"
)

type TestSKstore struct {
	meddler.DB
}

func NewTestSKstore(db meddler.DB) *TestSKstore {
	return &TestSKstore{db}
}

// 通过testsk的ID获取记录
func (db *TestSKstore) GetTestSKID(id int64) (*model.TestSK, error) {
	var testsk = new(model.TestSK)
	var err = meddler.Load(db, testSKTable, testsk, id)
	return testsk, err
}

func (db *TestSKstore) GetTestSKName(name string) (*model.TestSK, error) {
	var testsk = new(model.TestSK)
	var err = meddler.QueryRow(db, testsk, rebind(testskNQuery), name)
	return testsk, err
}

// 测试表名
const testSKTable = "testsks"

// 由测试人姓名取得对应记录
const testskNQuery = `
SELECT *
FROM testsks
WHERE testsk_name = ?
LIMIT 1
`
