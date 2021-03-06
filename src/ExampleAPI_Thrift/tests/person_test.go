package test

import (
	_ "ExampleAPI_Bigset_Thrift/src/ExampleAPI_Thrift/routers"
	"ExampleAPI_Bigset_Thrift/thrift/gen-go/myGeneric"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/astaxie/beego"
	log "github.com/astaxie/beego/logs"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

// test get all persons
func TestGetPersons(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/person", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, req)

	log.Trace("testing", "TestPerson", "Code[%d]\n%s", w.Code, w.Body.String())

	var response myGeneric.TPeronSetResult_
	_ = json.Unmarshal(w.Body.Bytes(), &response)

	Convey("Subject: Test Person Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})

	})
}

//test get persons pagination
func TestPagination(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/person?offset=0&limit=2", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, req)

	log.Trace("testing", "TestPerson", "Code[%d]\n%s", w.Code, w.Body.String())

	var response myGeneric.TPeronSetResult_
	_ = json.Unmarshal(w.Body.Bytes(), &response)

	Convey("Subject: Test Person Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})

	})
}

// test get person by id
func TestGetPersonById(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/person/p-4", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, req)

	log.Trace("testing", "TestPerson", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Person Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})

	})
}

// test create new person
func TestCreatePerson(t *testing.T) {
	var jsonStr = []byte(`{"personId":"p-12","personName":"Truong2","birthDate":"15-04-1998","personAddress":"HN2", "teamId":"t-4"}`)

	req, err := http.NewRequest("POST", "/v1/person", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, req)

	log.Trace("testing", "TestPerson", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Person Endpoint\n", t, func() {
		Convey("Status Code Should Be 201", func() {
			So(w.Code, ShouldEqual, 201)
		})
	})
}

// test update a person
func TestPutPerson(t *testing.T) {
	var jsonStr = []byte(`{"personId":"p-4","personName":"Truong_Put","birthDate":"15-04-1998","personAddress":"HN2", "teamId":"t-4"}`)

	req, err := http.NewRequest("PUT", "/v1/person/p-4", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, req)

	log.Trace("testing", "TestPerson", "Code[%d]\n%s", w.Code)

	Convey("Subject: Test Person Endpoint\n", t, func() {
		Convey("Status Code Should Be 204", func() {
			So(w.Code, ShouldEqual, 204)
		})
	})
}

// test get a list of person of the team
func TestGetPersonOfTeam(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/person/team/t-4", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, req)

	log.Trace("testing", "TestPerson", "Code[%d]\n%s", w.Code, w.Body.String())

	var response myGeneric.TPeronSetResult_
	_ = json.Unmarshal(w.Body.Bytes(), &response)

	Convey("Subject: Test Person Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})

	})
}

// test delete a person
func TestDeletePerson(t *testing.T) {

	req, err := http.NewRequest("DELETE", "/v1/person/p-4", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, req)

	log.Trace("testing", "TestPerson", "Code[%d]\n", w.Code)

	Convey("Subject: Test Person Endpoint\n", t, func() {
		Convey("Status Code Should Be 204", func() {
			So(w.Code, ShouldEqual, 204)
		})

	})
}
