// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/stainless-sdks/straddle-go/internal/apijson"
	"github.com/stainless-sdks/straddle-go/internal/requestconfig"
	"github.com/stainless-sdks/straddle-go/packages/param"
	"github.com/stainless-sdks/straddle-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type PageNumberSchemaMeta struct {
	MaxPageSize int64 `json:"max_page_size"`
	PageNumber  int64 `json:"page_number"`
	PageSize    int64 `json:"page_size"`
	TotalItems  int64 `json:"total_items"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxPageSize respjson.Field
		PageNumber  respjson.Field
		PageSize    respjson.Field
		TotalItems  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PageNumberSchemaMeta) RawJSON() string { return r.JSON.raw }
func (r *PageNumberSchemaMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageNumberSchema[T any] struct {
	Data []T                  `json:"data"`
	Meta PageNumberSchemaMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r PageNumberSchema[T]) RawJSON() string { return r.JSON.raw }
func (r *PageNumberSchema[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *PageNumberSchema[T]) GetNextPage() (res *PageNumberSchema[T], err error) {
	if len(r.Data) == 0 {
		return nil, nil
	}
	u := r.cfg.Request.URL
	currentPage, err := strconv.ParseInt(u.Query().Get("page_number"), 10, 64)
	if err != nil {
		currentPage = 1
	}
	cfg := r.cfg.Clone(context.Background())
	query := cfg.Request.URL.Query()
	query.Set("page_number", fmt.Sprintf("%d", currentPage+1))
	cfg.Request.URL.RawQuery = query.Encode()
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *PageNumberSchema[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &PageNumberSchema[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type PageNumberSchemaAutoPager[T any] struct {
	page *PageNumberSchema[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewPageNumberSchemaAutoPager[T any](page *PageNumberSchema[T], err error) *PageNumberSchemaAutoPager[T] {
	return &PageNumberSchemaAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *PageNumberSchemaAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Data) == 0 {
		return false
	}
	if r.idx >= len(r.page.Data) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Data) == 0 {
			return false
		}
	}
	r.cur = r.page.Data[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *PageNumberSchemaAutoPager[T]) Current() T {
	return r.cur
}

func (r *PageNumberSchemaAutoPager[T]) Err() error {
	return r.err
}

func (r *PageNumberSchemaAutoPager[T]) Index() int {
	return r.run
}
