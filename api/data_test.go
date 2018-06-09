package api

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"io"
	"bytes"
)

func TestMeta_DataHandler(t *testing.T) {
	var tests = []struct {
		Body       io.Reader
		StatusCode int
	}{
		{
			Body:       nil,
			StatusCode: http.StatusBadRequest,
		},
		{
			Body:       bytes.NewReader([]byte(`{}`)),
			StatusCode: http.StatusBadRequest,
		},
		{
			Body: bytes.NewReader([]byte(`{
    "source": {
        "id": "adsf",
        "type": "adsf"
    },
    "type": "adsf",
    "event": "asdf",
    "data": {}
}`)),
			StatusCode: http.StatusCreated,
		},
		{
			Body: bytes.NewReader([]byte(`{
    "source": {
        "id": "",
        "type": "adsf"
    },
    "type": "adsf",
    "event": "",
    "data": {}
}`)),
			StatusCode: http.StatusBadRequest,
		},
		{
			Body: bytes.NewReader([]byte(`{
    "source": {
        "id": "adsf",
        "type": "adsf"
    },
    "type": "adsf",
    "event": "asdf",
    "data": {}
}`)),
			StatusCode: http.StatusCreated,
		},
		{
			Body: bytes.NewReader([]byte(`{
    "source": {
        "id": "",
        "type": "adsf"
    },
    "type": "adsf",
    "event": "",
    "data": {}
}`)),
			StatusCode: http.StatusBadRequest,
		},
		{
			Body: bytes.NewReader([]byte(`{
    "source": {
        "id": "adsf",
        "type": "adsf"
    },
    "type": "adsf",
    "event": "asdf",
    "data": {}
}`)),
			StatusCode: http.StatusCreated,
		},
		{
			Body: bytes.NewReader([]byte(`{
    "source": {
        "id": "",
        "type": "adsf"
    },
    "type": "adsf",
    "event": "",
    "data": {}
}`)),
			StatusCode: http.StatusBadRequest,
		},
		{
			Body: bytes.NewReader([]byte(`{
    "source": {
        "id": "adsf",
        "type": "adsf"
    },
    "type": "adsf",
    "event": "asdf",
    "data": {}
}`)),
			StatusCode: http.StatusCreated,
		},
		{
			Body: bytes.NewReader([]byte(`{
    "source": {
        "id": "",
        "type": "adsf"
    },
    "type": "adsf",
    "event": "",
    "data": {}
}`)),
			StatusCode: http.StatusBadRequest,
		},
		{
			Body: bytes.NewReader([]byte(`{
    "source": {
        "id": "adsf",
        "type": "adsf"
    },
    "type": "adsf",
    "event": "asdf",
    "data": {}
}`)),
			StatusCode: http.StatusCreated,
		},
		{
			Body: bytes.NewReader([]byte(`{
    "source": {
        "id": "",
        "type": "adsf"
    },
    "type": "adsf",
    "event": "",
    "data": {}
}`)),
			StatusCode: http.StatusBadRequest,
		},
		{
			Body: bytes.NewReader([]byte(`{
    "source": {
        "id": "adsf",
        "type": "adsf"
    },
    "type": "adsf",
    "event": "asdf",
    "data": {}
}`)),
			StatusCode: http.StatusCreated,
		},
		{
			Body: bytes.NewReader([]byte(`{
    "source": {
        "id": "1",
        "type": "adsf"
    },
    "type": "adsf",
    "event": "1",
    "data": {}
}`)),
			StatusCode: http.StatusCreated,
		},
		{
			Body: bytes.NewReader([]byte(`{
    "source": {
        "id": "ad1sf",
        "type": "advsf"
    },
    "type": "adsaf",
    "event": "asdf",
    "data": {}
}`)),
			StatusCode: http.StatusCreated,
		},
		{
			Body: bytes.NewReader([]byte(`{
    "source": {
        "id": "adsf",
        "type": "adsf"
    },
    "type": "adsf",
    "event": "asdf",
    "data": {}
}`)),
			StatusCode: http.StatusCreated,
		},
		{
			Body: bytes.NewReader([]byte(`{
    "source": {
        "id": "1",
        "type": "adsf"
    },
    "type": "adsf",
    "event": "1",
    "data": {}
}`)),
			StatusCode: http.StatusCreated,
		},
		{
			Body: bytes.NewReader([]byte(`{
    "source": {
        "id": "ad1sf",
        "type": "advsf"
    },
    "type": "adsaf",
    "event": "asdf",
    "data": {}
}`)),
			StatusCode: http.StatusCreated,
		},
		{
			Body: bytes.NewReader([]byte(`{
    "source": {
        "id": "f3f",
        "type": "asf"
    },
    "type": "adsf",
    "event": "",
    "data": {}
}`)),
			StatusCode: http.StatusBadRequest,
		},
		{
			Body: bytes.NewReader([]byte(`{
    "source": {
        "id": "adsf",
        "type": "adsf"
    },
    "type": "adsf",
    "event": "asdf",
    "data": {}
}`)),
			StatusCode: http.StatusCreated,
		},
		{
			Body: bytes.NewReader([]byte(`{
    "source": {
        "id": "",
        "type": "adsf"
    },
    "type": "adsf",
    "event": "",
    "data": {}
}`)),
			StatusCode: http.StatusBadRequest,
		},
		{
			Body: bytes.NewReader([]byte(`{
    "source": {
        "id": "adsf",
        "type": "adsf"
    },
    "type": "adsf",
    "event": "asdf",
    "data": {}
}`)),
			StatusCode: http.StatusCreated,
		},
		{
			Body: bytes.NewReader([]byte(`{
    "source": {
        "id": "",
        "type": "adsf"
    },
    "type": "adsf",
    "event": "",
    "data": {}
}`)),
			StatusCode: http.StatusBadRequest,
		},
		{
			Body: bytes.NewReader([]byte(`{
    "source": {
        "id": "adsf",
        "type": "adsf"
    },
    "type": "adsf",
    "event": "asdf",
    "data": {}
}`)),
			StatusCode: http.StatusCreated,
		},
		{
			Body: bytes.NewReader([]byte(`{
    "source": {
        "id": "1",
        "type": "adsf"
    },
    "type": "adsf",
    "event": "1",
    "data": {}
}`)),
			StatusCode: http.StatusCreated,
		},
		{
			Body: bytes.NewReader([]byte(`{
    "source": {
        "id": "ad1sf",
        "type": "advsf"
    },
    "type": "adsaf",
    "event": "asdf",
    "data": {}
}`)),
			StatusCode: http.StatusCreated,
		},
		{
			Body: bytes.NewReader([]byte(`{
    "source": {
        "id": "f3f",
        "type": "asf"
    },
    "type": "adsf",
    "event": "",
    "data": {}
}`)),
			StatusCode: http.StatusBadRequest,
		},
	}
	meta := Meta{DB: *db}

	for _, test := range tests {
		req, err := http.NewRequest("POST", "/v1/data", test.Body)
		if err != nil {
			t.Error("exception not expected", err)
			t.Fail()
		}

		rr := httptest.NewRecorder()

		meta.DataHandler(rr, req)
		status := rr.Code
		if test.StatusCode != status {
			t.Errorf("handler returned wrong status code: got %v want %v", status, test.StatusCode)
		}
	}

}
