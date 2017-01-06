package foreman

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestHostsService_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/hosts/host01", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"name":"host01"}`)
	})

	host, _, err := client.Hosts.Get("host01")
	if err != nil {
		t.Errorf("Hosts.Get returned error: %v", err)
	}

	expected := &Host{Name: String("host01")}
	if !reflect.DeepEqual(host, expected) {
		t.Errorf("Hosts.Get: \n\n%#v\n\n%#v\n\n", host, expected)
	}
}

func TestHostsService_Get_All(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/hosts", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"Name":"host01"}]`)
	})

	opt := &HostGetAllOptions{1, 1, 1, 1, 1, 1, GetOptions{2, 3}}
	hosts, _, err := client.Hosts.GetAll(opt)
	if err != nil {
		t.Errorf("Hosts.GetAll returned error: %v", err)
	}

	expected := []*Host{{Name: String("host01")}}
	if !reflect.DeepEqual(hosts, expected) {
		t.Errorf("Hosts.GetAll: \n\n%#v\n\n%#v\n\n", hosts, expected)
	}
}

func TestHostsService_Create(t *testing.T) {
	setup()
	defer teardown()

	input := &Host{Name: String("host01")}

	mux.HandleFunc("/api/hosts", func(w http.ResponseWriter, r *http.Request) {
		v := new(Host)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, expected %+v", v, input)
		}

		fmt.Fprint(w, `{"name":"host01"}`)
	})

	host, _, err := client.Hosts.Create(input)
	if err != nil {
		t.Errorf("Hosts.Create returned error: %v", err)
	}

	expected := &Host{Name: String("host01")}
	if !reflect.DeepEqual(host, expected) {
		t.Errorf("Hosts.Create returned %+v, expected %+v", host, expected)
	}
}

func TestHostsService_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/hosts/host01", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Hosts.Delete("host01")
	if err != nil {
		t.Errorf("Hosts.Delete returned %v", err)
	}
}
