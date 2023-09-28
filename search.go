package search

import (
	"encoding/json"
	"sync"
)

// RangeRelation
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-range-query.html#querying-range-fields
type RangeRelation string

func (rr RangeRelation) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(rr))
}

func (rr *RangeRelation) UnmarshalJSON(b []byte) error {
	var tmp string
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*rr = RangeRelation(tmp)
	return nil
}

const (
	RangeRelationIntersects RangeRelation = "INTERSECTS"
	RangeRelationContains   RangeRelation = "CONTAINS"
	RangeRelationWithin     RangeRelation = "WITHIN"
)

type MatchOperator string

func (mo MatchOperator) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(mo))
}

func (mo *MatchOperator) UnmarshalJSON(b []byte) error {
	var tmp string
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*mo = MatchOperator(tmp)
	return nil
}

const (
	MatchOperatorOr  MatchOperator = "OR"
	MatchOperatorAnd MatchOperator = "AND"
)

// ZeroTermsQuery
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-query.html#query-dsl-match-query-zero
type ZeroTermsQuery string

func (ztq ZeroTermsQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(ztq))
}

func (ztq *ZeroTermsQuery) UnmarshalJSON(b []byte) error {
	var tmp string
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*ztq = ZeroTermsQuery(tmp)
	return nil
}

const (
	ZeroTermsQueryAll  ZeroTermsQuery = "all"
	ZeroTermsQueryNone ZeroTermsQuery = "None"
)

// SortOrder
// https://www.elastic.co/guide/en/elasticsearch/reference/current/sort-search-results.html#_sort_order
type SortOrder string

func (so SortOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(so))
}

func (so *SortOrder) UnmarshalJSON(b []byte) error {
	var tmp string
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*so = SortOrder(tmp)
	return nil
}

const (
	SortOrderAsc  SortOrder = "asc"
	SortOrderDesc SortOrder = "desc"
)

// SortMode
// https://www.elastic.co/guide/en/elasticsearch/reference/current/sort-search-results.html#_sort_mode_option
type SortMode string

func (sm SortMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(sm))
}

func (sm *SortMode) UnmarshalJSON(b []byte) error {
	var tmp string
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*sm = SortMode(tmp)
	return nil
}

const (
	SortModeMin    SortMode = "min"
	SortModeMax    SortMode = "max"
	SortModeSum    SortMode = "sum"
	SortModeAvg    SortMode = "avg"
	SortModeMedian SortMode = "median"
)

// Bool
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-bool-query.html
type Bool struct {
	Must    []*Query `json:"must,omitempty"`
	MustNot []*Query `json:"must_not,omitempty"`
	Should  []*Query `json:"should,omitempty"`
	Filter  *Query   `json:"filter,omitempty"`
}

func NewBool(opts ...func(*Bool)) *Bool {
	b := new(Bool)
	for _, fn := range opts {
		fn(b)
	}
	return b
}

func (b *Bool) AddMust(m *Query) *Bool {
	b.Must = append(b.Must, m)
	return b
}

func (b *Bool) AddMustNot(mn *Query) *Bool {
	b.MustNot = append(b.MustNot, mn)
	return b
}

func (b *Bool) AddShould(s *Query) *Bool {
	b.Should = append(b.Should, s)
	return b
}

func (b *Bool) SetFilter(q *Query) *Bool {
	b.Filter = q
	return b
}

// Term
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-term-query.html
type Term struct {
	Value         string  `json:"value"`
	Boost         float64 `json:"boost,omitempty"`
	CaseSensitive *bool   `json:"case_sensitive,omitempty"`
}

func NewTerm(opts ...func(*Term)) *Term {
	t := new(Term)
	for _, fn := range opts {
		fn(t)
	}
	return t
}

func (t *Term) SetValue(v string) *Term {
	t.Value = v
	return t
}

func (t *Term) SetBoost(b float64) *Term {
	t.Boost = b
	return t
}

func (t *Term) SetCaseSensitivity(b bool) *Term {
	t.CaseSensitive = new(bool)
	*t.CaseSensitive = b
	return t
}

// Range
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-range-query.html
type Range struct {
	GT       any           `json:"gt,omitempty"`
	GTE      any           `json:"gte,omitempty"`
	LT       any           `json:"lt,omitempty"`
	LTE      any           `json:"lte,omitempty"`
	Boost    float64       `json:"boost,omitempty"`
	Format   string        `json:"format,omitempty"`
	Relation RangeRelation `json:"relation,omitempty"`
	TimeZone string        `json:"time_zone,omitempty"`
}

func NewRange(opts ...func(*Range)) *Range {
	r := new(Range)
	for _, fn := range opts {
		fn(r)
	}
	return r
}

func (r *Range) SetGT(gt any) *Range {
	r.GT = gt
	return r
}

func (r *Range) SetGTE(gte any) *Range {
	r.GTE = gte
	return r
}

func (r *Range) SetLT(lt any) *Range {
	r.LT = lt
	return r
}

func (r *Range) SetLTE(lte any) *Range {
	r.LTE = lte
	return r
}

func (r *Range) SetBoost(b float64) *Range {
	r.Boost = b
	return r
}

func (r *Range) SetFormat(f string) *Range {
	r.Format = f
	return r
}

func (r *Range) SetRelation(rr RangeRelation) *Range {
	r.Relation = rr
	return r
}

func (r *Range) SetTimeZone(tz string) *Range {
	r.TimeZone = tz
	return r
}

// MatchAll
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-all-query.html
type MatchAll struct {
	Boost float64 `json:"boost,omitempty"`
}

func NewMatchAll(opts ...func(*MatchAll)) *MatchAll {
	ma := new(MatchAll)
	for _, fn := range opts {
		fn(ma)
	}
	return ma
}

func (ma *MatchAll) SetBoost(b float64) *MatchAll {
	ma.Boost = b
	return ma
}

// MatchNone
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-all-query.html#query-dsl-match-none-query
type MatchNone struct{}

func NewMatchNone(opts ...func(*MatchNone)) *MatchNone {
	mn := new(MatchNone)
	for _, fn := range opts {
		fn(mn)
	}
	return mn
}

// Match
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-query.html
type Match struct {
	Query               any           `json:"query,omitempty"`
	Operator            MatchOperator `json:"operator,omitempty"`
	Boost               float64       `json:"boost,omitempty"`
	FuzzyTranspositions *bool         `json:"fuzzy_transpositions,omitempty"`
	FuzzyRewrite        string        `json:"fuzzy_rewrite,omitempty"`
	MaxExpansions       int           `json:"max_expansions,omitempty"`
	Lenient             *bool         `json:"lenient,omitempty"`

	// ZeroTermsQuery
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-query.html#query-dsl-match-query-zero
	ZeroTermsQuery ZeroTermsQuery `json:"zero_terms_query,omitempty"`

	// Fuzziness
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/common-options.html#fuzziness
	Fuzziness    string `json:"fuzziness,omitempty"`
	PrefixLength int    `json:"prefix_length,omitempty"`

	// MinimumShouldMatch
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-minimum-should-match.html
	MinimumShouldMatch string `json:"minimum_should_match,omitempty"`
}

func NewMatch(opts ...func(*Match)) *Match {
	m := new(Match)
	for _, fn := range opts {
		fn(m)
	}
	return m
}

func (m *Match) SetQuery(q any) *Match {
	m.Query = q
	return m
}

func (m *Match) SetOperator(mo MatchOperator) *Match {
	m.Operator = mo
	return m
}

func (m *Match) SetBoost(b float64) *Match {
	m.Boost = b
	return m
}

func (m *Match) SetFuzzyTransposition(b bool) *Match {
	m.FuzzyTranspositions = new(bool)
	*m.FuzzyTranspositions = b
	return m
}

func (m *Match) SetFuzzyRewrite(fr string) *Match {
	m.FuzzyRewrite = fr
	return m
}

func (m *Match) SetMaxExpansions(me int) *Match {
	m.MaxExpansions = me
	return m
}

func (m *Match) SetLenient(l bool) *Match {
	m.Lenient = new(bool)
	*m.Lenient = l
	return m
}

func (m *Match) SetZeroTermsQuery(ztq ZeroTermsQuery) *Match {
	m.ZeroTermsQuery = ztq
	return m
}

func (m *Match) SetFuzziness(f string) *Match {
	m.Fuzziness = f
	return m
}

func (m *Match) SetPrefixLength(n int) *Match {
	m.PrefixLength = n
	return m
}

func (m *Match) SetMinimumShouldMatch(msm string) *Match {
	m.MinimumShouldMatch = msm
	return m
}

// Query
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl.html
type Query struct {
	mu sync.Mutex

	MatchAll  *MatchAll  `json:"match_all,omitempty"`
	MatchNone *MatchNone `json:"match_none,omitempty"`

	Bool  *Bool            `json:"bool,omitempty"`
	Term  map[string]Term  `json:"term,omitempty"`
	Match map[string]Match `json:"match,omitempty"`
}

func NewQuery(opts ...func(*Query)) *Query {
	q := new(Query)
	for _, fn := range opts {
		fn(q)
	}
	return q
}

func (q *Query) SetMatchAll(maq *MatchAll) *Query {
	q.MatchAll = maq
	return q
}

func (q *Query) SetMatchNone(mnq *MatchNone) *Query {
	q.MatchNone = mnq
	return q
}

func (q *Query) SetBool(b *Bool) *Query {
	q.Bool = b
	return q
}

func (q *Query) SetTerm(field string, t Term) *Query {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.Term == nil {
		q.Term = make(map[string]Term)
	}
	q.Term[field] = t
	return q
}

func (q *Query) SetMatch(field string, m Match) *Query {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.Match == nil {
		q.Match = make(map[string]Match)
	}
	q.Match[field] = m
	return q
}

type Sort struct {
	Order  SortOrder `json:"order,omitempty"`
	Type   string    `json:"type,omitempty"`
	Format string    `json:"format,omitempty"`
	Mode   SortMode  `json:"mode,omitempty"`
}

func NewsSort(opts ...func(*Sort)) *Sort {
	s := new(Sort)
	for _, fn := range opts {
		fn(s)
	}
	return s
}

func (s *Sort) SetOrder(o SortOrder) *Sort {
	s.Order = o
	return s
}

func (s *Sort) SetType(t string) *Sort {
	s.Type = t
	return s
}

func (s *Sort) SetFormat(f string) *Sort {
	s.Format = f
	return s
}

func (s *Sort) SetMode(m SortMode) *Sort {
	s.Mode = m
	return s
}

// Search
// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-search.html
type Search struct {
	Size int `json:"size,omitempty"`

	// Sort
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/sort-search-results.html
	Sort []map[string]any `json:"sort,omitempty"`

	Query *Query `json:"query,omitempty"`

	// Source
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-fields.html#source-filtering
	Source any `json:"_source,omitempty"`
}

func NewSearch(opts ...func(*Search)) *Search {
	s := new(Search)
	for _, fn := range opts {
		fn(s)
	}
	return s
}

func (s *Search) SetSize(sz int) *Search {
	s.Size = sz
	return s
}

func (s *Search) SetQuery(q *Query) *Search {
	s.Query = q
	return s
}

func (s *Search) SetSource(sc any) *Search {
	s.Source = sc
	return s
}

func (s *Search) AddSort(st map[string]any) *Search {
	s.Sort = append(s.Sort, st)
	return s
}
