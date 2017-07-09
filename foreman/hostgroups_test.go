package foreman

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestHostgroupsService_Get_Hostgroup(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/hostgroups/hostgroup01", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"name":"hostgroup01"}`)
	})

	hostgroup, _, err := client.Hostgroups.Get(context.Background(), "hostgroup01")
	if err != nil {
		t.Errorf("Hosts.Get returned error: %v", err)
	}

	expected := &Hostgroup{Name: String("hostgroup01")}
	if !reflect.DeepEqual(hostgroup, expected) {
		t.Errorf("Hosts.Get: \n\n%#v\n\n%#v\n\n", hostgroup, expected)
	}
}

func TestHostgroupsService_Get_All(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/hostgroups", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"Name":"hostgroup01"}]`)
	})

	opt := &HostgroupGetAllOptions{"", "", "", GetOptions{2, 3}}
	hostgroups, _, err := client.Hostgroups.GetAll(context.Background(), opt)
	if err != nil {
		t.Errorf("Hosts.GetAll returned error: %v", err)
	}

	expected := []*Hostgroup{{Name: String("hostgroup01")}}
	if !reflect.DeepEqual(hostgroups, expected) {
		t.Errorf("Hosts.GetAll: \n\n%#v\n\n%#v\n\n", hostgroups, expected)
	}
}

func TestHostgroupsService_Create_Hostgroup(t *testing.T) {
	setup()
	defer teardown()

	input := &Hostgroup{Name: String("hostgroup01")}

	mux.HandleFunc("/api/hostgroups", func(w http.ResponseWriter, r *http.Request) {
		v := new(Hostgroup)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, expected %+v", v, input)
		}

		fmt.Fprint(w, `{"name":"hostgroup01"}`)
	})

	hostgroup, _, err := client.Hostgroups.Create(context.Background(), input)
	if err != nil {
		t.Errorf("Hosts.Create returned error: %v", err)
	}

	expected := &Hostgroup{Name: String("hostgroup01")}
	if !reflect.DeepEqual(hostgroup, expected) {
		t.Errorf("Hosts.Create returned %+v, expected %+v", hostgroup, expected)
	}
}
