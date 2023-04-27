package simple

import (
	"gitee.com/pangxianfei/simple/sqlcmd"
	"github.com/kataras/iris/v12"

	"gitee.com/pangxianfei/simple/strcase"
)

type QueryParams struct {
	Ctx    iris.Context
	SqlCnd sqlcmd.Cnd
}

func NewQueryParams(ctx iris.Context) *QueryParams {
	return &QueryParams{
		Ctx: ctx,
	}
}

func (q *QueryParams) getValueByColumn(column string) string {
	if q.Ctx == nil {
		return ""
	}
	fieldName := strcase.ToLowerCamel(column)
	return q.Ctx.FormValue(fieldName)
}

func (q *QueryParams) EqByReq(column string) *QueryParams {
	value := q.getValueByColumn(column)
	if len(value) > 0 {
		q.SqlCnd.Eq(column, value)

	}
	return q
}

func (q *QueryParams) NotEqByReq(column string) *QueryParams {
	value := q.getValueByColumn(column)
	if len(value) > 0 {
		q.SqlCnd.NotEq(column, value)
	}
	return q
}

func (q *QueryParams) GtByReq(column string) *QueryParams {
	value := q.getValueByColumn(column)
	if len(value) > 0 {
		q.SqlCnd.Gt(column, value)
	}
	return q
}

func (q *QueryParams) GteByReq(column string) *QueryParams {
	value := q.getValueByColumn(column)
	if len(value) > 0 {
		q.SqlCnd.Gte(column, value)
	}
	return q
}

func (q *QueryParams) LtByReq(column string) *QueryParams {
	value := q.getValueByColumn(column)
	if len(value) > 0 {
		q.SqlCnd.Lt(column, value)
	}
	return q
}

func (q *QueryParams) LteByReq(column string) *QueryParams {
	value := q.getValueByColumn(column)
	if len(value) > 0 {
		q.SqlCnd.Lte(column, value)
	}
	return q
}

func (q *QueryParams) LikeByReq(column string) *QueryParams {
	value := q.getValueByColumn(column)
	if len(value) > 0 {
		q.SqlCnd.Like(column, value)
	}
	return q
}

func (q *QueryParams) PageByReq() *QueryParams {
	if q.Ctx == nil {
		return q
	}
	paging := GetPaging(q.Ctx)
	q.Page(paging.Page, paging.Limit)
	return q
}

func (q *QueryParams) Asc(column string) *QueryParams {
	q.SqlCnd.Orders = append(q.SqlCnd.Orders, sqlcmd.OrderByCol{Column: column, Asc: true})
	return q
}

func (q *QueryParams) Desc(column string) *QueryParams {
	q.SqlCnd.Orders = append(q.SqlCnd.Orders, sqlcmd.OrderByCol{Column: column, Asc: false})
	return q
}

func (q *QueryParams) Limit(limit int) *QueryParams {
	q.Page(1, limit)
	return q
}

func (q *QueryParams) Page(page, limit int) *QueryParams {
	if q.SqlCnd.Paging == nil {
		q.SqlCnd.Paging = &sqlcmd.Paging{Page: page, Limit: limit}
	} else {
		q.SqlCnd.Paging.Page = page
		q.SqlCnd.Paging.Limit = limit
	}
	return q
}
