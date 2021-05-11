package planner

import (
	"fmt"
	"os"
	"testing"

	"github.com/gridbase/sqldb/db"
	"github.com/gridbase/sqldb/parser"
	. "github.com/smartystreets/goconvey/convey"
)

func TestPlan_Insert(t *testing.T) {
	Convey("Plan Insert one value", t, func() {
		const testFile = "test_plan_insert.db"
		defer func() {
			_ = os.Remove(testFile)
		}()
		table, err := db.Open(testFile)
		So(err, ShouldBeNil)
		plan := NewPlan(table)

		p := &parser.Parser{}
		ast, err := p.ParseInsert("INSERT INTO table VALUES " +
			"(10, auxten, \"auxtenwpc@gmail.com\")",
		)
		So(err, ShouldBeNil)

		cnt, err := plan.Insert(ast)
		So(err, ShouldBeNil)
		So(cnt, ShouldEqual, 1)

		p2 := &parser.Parser{}
		ast2, err := p2.ParseSelect("SELECT id, username, email FROM table")
		So(err, ShouldBeNil)

		plan2 := NewPlan(table)
		resultPipe, err := plan2.SelectPrepare(ast2)
		So(err, ShouldBeNil)
		for row := range resultPipe {
			fmt.Println(row.Id, string(row.Username[:]), string(row.Email[:]))
		}
	})
}

func TestPlan_Insert_multiple(t *testing.T) {
	Convey("Plan Insert multiple values", t, func() {
		const testFile = "test_plan_insert_multiple.db"
		defer func() {
			_ = os.Remove(testFile)
		}()
		table, err := db.Open(testFile)
		So(err, ShouldBeNil)
		plan := NewPlan(table)

		p := &parser.Parser{}
		ast, err := p.ParseInsert("INSERT INTO table (id, username, email) VALUES " +
			"(0, auxten, \"auxtenwpc@gmail.com\")," +
			"(1, hahaha, \"hahaha@gmail.com\")," +
			"(2, aaaa, \"aaaa@gmail.com\")," +
			"(3, jijiji, \"jijiji@gmail.com\")",
		)
		So(err, ShouldBeNil)

		cnt, err := plan.Insert(ast)
		So(err, ShouldBeNil)
		So(cnt, ShouldEqual, 4)

		p2 := &parser.Parser{}
		ast2, err := p2.ParseSelect("SELECT id, username, email FROM table LIMIT 10")
		So(err, ShouldBeNil)

		plan2 := NewPlan(table)
		resultPipe, err := plan2.SelectPrepare(ast2)
		So(err, ShouldBeNil)
		for row := range resultPipe {
			fmt.Println(row.Id, string(row.Username[:]), string(row.Email[:]))
		}
	})
}
