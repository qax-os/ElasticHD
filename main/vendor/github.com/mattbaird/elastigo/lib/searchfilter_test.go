// Copyright 2013 Matthew Baird
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package elastigo

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFilterDsl(t *testing.T) {
	Convey("And filter", t, func() {
		filter := Filter().And(Filter().Term("test", "asdf")).
		And(Filter().Range("rangefield", 1, 2, 3, 4, "+08:00"))
		actual, err := GetJson(filter)

		actualFilters := actual["and"].([]interface{})

		So(err, ShouldBeNil)
		So(1, ShouldEqual, len(actual))
		So(2, ShouldEqual, len(actualFilters))
		So(true, ShouldEqual, HasKey(actualFilters[0].(map[string]interface{}), "term"))
		So(true, ShouldEqual, HasKey(actualFilters[1].(map[string]interface{}), "range"))
	})

	Convey("Or filter", t, func() {
		filter := Filter().Or(Filter().Term("test", "asdf"), Filter().Range("rangefield", 1, 2, 3, 4, "+08:00"))
		actual, err := GetJson(filter)

		actualFilters := actual["or"].([]interface{})

		So(err, ShouldBeNil)
		So(1, ShouldEqual, len(actual))
		So(2, ShouldEqual, len(actualFilters))
		So(true, ShouldEqual, HasKey(actualFilters[0].(map[string]interface{}), "term"))
		So(true, ShouldEqual, HasKey(actualFilters[1].(map[string]interface{}), "range"))
	})

	Convey("Not filter", t, func() {
		filter := Filter().Not(Filter().Term("test", "asdf")).
		Not(Filter().Range("rangefield", 1, 2, 3, 4, "+08:00"))
		actual, err  := GetJson(filter)

		actualFilters := actual["not"].([]interface{})

		So(err, ShouldBeNil)
		So(1, ShouldEqual, len(actual))
		So(2, ShouldEqual, len(actualFilters))
		So(true, ShouldEqual, HasKey(actualFilters[0].(map[string]interface{}), "term"))
		So(true, ShouldEqual, HasKey(actualFilters[1].(map[string]interface{}), "range"))
	})

	Convey("Terms filter", t, func() {
		filter := Filter().Terms("Sample", TEMAnd, "asdf", 123, true)
		actual, err  := GetJson(filter)

		actualTerms := actual["terms"].(map[string]interface{})
		actualValues := actualTerms["Sample"].([]interface{})

		So(err, ShouldBeNil)
		So(1, ShouldEqual, len(actual))
		So(3, ShouldEqual, len(actualValues))
		So(actualValues[0], ShouldEqual, "asdf")
		So(actualValues[1], ShouldEqual, float64(123))
		So(actualValues[2], ShouldEqual, true)
		So("and", ShouldEqual, actualTerms["execution"])
	})

	Convey("Term filter", t, func() {
		filter := Filter().Term("Sample", "asdf").Term("field2", 341.4)
		actual, err  := GetJson(filter)

		actualTerm := actual["term"].(map[string]interface{})

		So(err, ShouldBeNil)
		So(1, ShouldEqual, len(actual))
		So("asdf", ShouldEqual, actualTerm["Sample"])
		So(float64(341.4), ShouldEqual, actualTerm["field2"])
	})

	Convey("Range filter", t, func() {
		filter := Filter().Range("rangefield", 1, 2, 3, 4, "+08:00")
		actual, err  := GetJson(filter)
		//A bit lazy, probably should assert keys exist
		actualRange := actual["range"].(map[string]interface{})["rangefield"].(map[string]interface{})

		So(err, ShouldBeNil)
		So(1, ShouldEqual, len(actual))
		So(float64(1), ShouldEqual, actualRange["gte"])
		So(float64(2), ShouldEqual, actualRange["gt"])
		So(float64(3), ShouldEqual, actualRange["lte"])
		So(float64(4), ShouldEqual, actualRange["lt"])
		So("+08:00", ShouldEqual, actualRange["time_zone"])
	})

	Convey("Exists filter", t, func() {
		filter := Filter().Exists("field1")
		actual, err  := GetJson(filter)

		actualValue := actual["exists"].(map[string]interface{})

		So(err, ShouldBeNil)
		So(1, ShouldEqual, len(actual))
		So("field1", ShouldEqual, actualValue["field"])
	})

	Convey("Missing filter", t, func() {
		filter := Filter().Missing("field1")
		actual, err  := GetJson(filter)

		actualValue := actual["missing"].(map[string]interface{})

		So(err, ShouldBeNil)
		So(1, ShouldEqual, len(actual))
		So("field1", ShouldEqual, actualValue["field"])
	})

	Convey("Limit filter", t, func() {
		filter := Filter().Limit(100)
		actual, err  := GetJson(filter)

		actualValue := actual["limit"].(map[string]interface{})
		So(err, ShouldBeNil)
		So(1, ShouldEqual, len(actual))
		So(float64(100), ShouldEqual, actualValue["value"])
	})

	Convey("Type filter", t, func() {
		filter := Filter().Type("my_type")
		actual, err  := GetJson(filter)

		actualValue := actual["type"].(map[string]interface{})
		So(err, ShouldBeNil)
		So(1, ShouldEqual, len(actual))
		So("my_type", ShouldEqual, actualValue["value"])
	})

	Convey("Ids filter", t, func() {
		filter := Filter().Ids("test", "asdf", "fdsa")
		actual, err  := GetJson(filter)

		actualValue := actual["ids"].(map[string]interface{})
		actualValues := actualValue["values"].([]interface{})
		So(err, ShouldBeNil)
		So(1, ShouldEqual, len(actual))
		So(nil, ShouldEqual, actualValue["type"])
		So(3, ShouldEqual, len(actualValues))
		So("test", ShouldEqual, actualValues[0])
		So("asdf", ShouldEqual, actualValues[1])
		So("fdsa", ShouldEqual, actualValues[2])
	})

	Convey("IdsByTypes filter", t, func() {
		filter := Filter().IdsByTypes([]string{"my_type"}, "test", "asdf", "fdsa")
		actual, err  := GetJson(filter)

		actualValue := actual["ids"].(map[string]interface{})
		actualTypes := actualValue["type"].([]interface{})
		actualValues := actualValue["values"].([]interface{})
		So(err, ShouldBeNil)
		So(1, ShouldEqual, len(actual))
		So(1, ShouldEqual, len(actualTypes))
		So("my_type", ShouldEqual, actualTypes[0])
		So(3, ShouldEqual, len(actualValues))
		So("test", ShouldEqual, actualValues[0])
		So("asdf", ShouldEqual, actualValues[1])
		So("fdsa", ShouldEqual, actualValues[2])
	})

	Convey("GeoDistance filter", t, func() {
		filter := Filter().GeoDistance("100km", NewGeoField("pin.location", 32.3, 23.4))
		actual, err  := GetJson(filter)

		actualValue := actual["geo_distance"].(map[string]interface{})
		actualLocation := actualValue["pin.location"].(map[string]interface{})
		So(err, ShouldBeNil)
		So("100km", ShouldEqual, actualValue["distance"])
		So(float64(32.3), ShouldEqual, actualLocation["lat"])
		So(float64(23.4), ShouldEqual, actualLocation["lon"])
	})

	Convey("GeoDistanceRange filter", t, func() {
		filter := Filter().GeoDistanceRange("100km", "200km", NewGeoField("pin.location", 32.3, 23.4))
		actual, err  := GetJson(filter)

		actualValue := actual["geo_distance_range"].(map[string]interface{})
		actualLocation := actualValue["pin.location"].(map[string]interface{})
		So(err, ShouldBeNil)
		So("100km", ShouldEqual, actualValue["from"])
		So("200km", ShouldEqual, actualValue["to"])
		So(float64(32.3), ShouldEqual, actualLocation["lat"])
		So(float64(23.4), ShouldEqual, actualLocation["lon"])
	})
}

func TestFilters(t *testing.T) {

	c := NewTestConn()
	PopulateTestDB(t, c)
	defer TearDownTestDB(c)

	Convey("Exists filter", t, func() {
		qry := Search("oilers").Filter(
			Filter().Exists("goals"),
		)
		out, err := qry.Result(c)
		So(err, ShouldBeNil)
		So(out, ShouldNotBeNil)
		So(out.Hits, ShouldNotBeNil)
		So(out.Hits.Len(), ShouldEqual, 10)
		So(out.Hits.Total, ShouldEqual, 12)
	})

	Convey("Missing filter", t, func() {
		qry := Search("oilers").Filter(
			Filter().Missing("goals"),
		)
		out, err := qry.Result(c)
		So(err, ShouldBeNil)
		So(out, ShouldNotBeNil)
		So(out.Hits, ShouldNotBeNil)
		So(out.Hits.Total, ShouldEqual, 2)
	})

	Convey("Terms filter", t, func() {
		qry := Search("oilers").Filter(
			Filter().Terms("pos", TEMDefault, "RW", "LW"),
		)
		out, err := qry.Result(c)
		So(err, ShouldBeNil)
		So(out, ShouldNotBeNil)
		So(out.Hits, ShouldNotBeNil)
		So(out.Hits.Total, ShouldEqual, 6)
	})

	Convey("Filter involving an AND", t, func() {
		qry := Search("oilers").Filter(
			Filter().And(
				Filter().Terms("pos", TEMDefault, "LW"),
				Filter().Exists("PIM"),
			),
		)

		out, err := qry.Result(c)
		So(err, ShouldBeNil)
		So(out, ShouldNotBeNil)
		So(out.Hits, ShouldNotBeNil)
		So(out.Hits.Total, ShouldEqual, 2)
	})


	Convey("Filterng filter results", t, func() {
		qry := Search("oilers").Filter(
			Filter().Terms("pos", TEMDefault, "LW"),
		)
		qry.Filter(
			Filter().Exists("PIM"),
		)
		out, err := qry.Result(c)
		So(err, ShouldBeNil)
		So(out, ShouldNotBeNil)
		So(out.Hits, ShouldNotBeNil)
		So(out.Hits.Total, ShouldEqual, 2)
	})

	Convey("Filter involving OR", t, func() {
		qry := Search("oilers").Filter(
			Filter().Or(
				Filter().Terms("pos", TEMDefault, "G"),
				Filter().Range("goals", nil, 80, nil, nil, ""),
			),
		)
		out, err := qry.Result(c)
		So(err, ShouldBeNil)
		So(out, ShouldNotBeNil)
		So(out.Hits, ShouldNotBeNil)
		So(out.Hits.Total, ShouldEqual, 3)
	})
}
