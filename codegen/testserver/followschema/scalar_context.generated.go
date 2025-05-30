// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package followschema

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) unmarshalNStringFromContextFunction2string(ctx context.Context, v any) (string, error) {
	res, err := UnmarshalStringFromContextFunction(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNStringFromContextFunction2string(ctx context.Context, sel ast.SelectionSet, v string) graphql.Marshaler {
	res := MarshalStringFromContextFunction(v)
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
	}
	return graphql.WrapContextMarshaler(ctx, res)
}

func (ec *executionContext) unmarshalNStringFromContextInterface2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐStringFromContextInterface(ctx context.Context, v any) (StringFromContextInterface, error) {
	var res StringFromContextInterface
	err := res.UnmarshalGQLContext(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNStringFromContextInterface2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐStringFromContextInterface(ctx context.Context, sel ast.SelectionSet, v StringFromContextInterface) graphql.Marshaler {
	return graphql.WrapContextMarshaler(ctx, v)
}

func (ec *executionContext) unmarshalNStringFromContextInterface2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐStringFromContextInterface(ctx context.Context, v any) (*StringFromContextInterface, error) {
	var res = new(StringFromContextInterface)
	err := res.UnmarshalGQLContext(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNStringFromContextInterface2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐStringFromContextInterface(ctx context.Context, sel ast.SelectionSet, v *StringFromContextInterface) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return graphql.WrapContextMarshaler(ctx, v)
}

// endregion ***************************** type.gotpl *****************************
