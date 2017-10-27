package karigo

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/kkaribu/jsonapi"
	"github.com/kkaribu/tchek"
)

func TestServerGetRequest(t *testing.T) {
	mApp := NewMockApp()
	mServer := httptest.NewServer(mApp)
	defer mServer.Close()

	col := jsonapi.WrapCollection(jsonapi.Wrap(&jsonapi.MockType3{}))
	col.Add(jsonapi.Wrap(&jsonapi.MockType3{
		ID: "mt3-1",
	}))
	col.Add(jsonapi.Wrap(&jsonapi.MockType3{
		ID: "mt3-2",
	}))

	mApp.Store = &MockStore{
		SelectCollectionFunc: func(tx Tx, resType string, from jsonapi.FromFilter, params *jsonapi.Params, c jsonapi.Collection) error {
			for i := 0; i < col.Len(); i++ {
				c.Add(col.Elem(i))
			}
			return nil
		},
	}

	url, err := url.Parse(mServer.URL + "/mocktypes3")
	tchek.UnintendedError(err)

	jurl, err := jsonapi.ParseURL(jsonapi.NewMockRegistry(), url)
	tchek.UnintendedError(err)

	res, err := http.Get(url.String())
	tchek.UnintendedError(err)

	doc := jsonapi.NewDocument()
	doc.Data = col
	doc.Meta["total-pages"] = 1

	expectedBody, err := jsonapi.Marshal(doc, jurl)
	tchek.UnintendedError(err)

	tchek.AreEqual(t, 0, string(expectedBody), string(readBody(res.Body)))
}

func readBody(r io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	return buf.String()
}
