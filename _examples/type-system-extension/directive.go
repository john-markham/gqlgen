package type_system_extension

import (
	"context"
	"log"

	"github.com/john-markham/gqlgen/graphql"
)

func EnumLogging(ctx context.Context, obj any, next graphql.Resolver) (res any, err error) {
	rc := graphql.GetFieldContext(ctx)
	log.Printf("enum logging: %v, %s, %T, %+v", rc.Path(), rc.Field.Name, obj, obj)
	return next(ctx)
}

func FieldLogging(ctx context.Context, obj any, next graphql.Resolver) (res any, err error) {
	rc := graphql.GetFieldContext(ctx)
	log.Printf("field logging: %v, %s, %T, %+v", rc.Path(), rc.Field.Name, obj, obj)
	return next(ctx)
}

func InputLogging(ctx context.Context, obj any, next graphql.Resolver) (res any, err error) {
	rc := graphql.GetFieldContext(ctx)
	log.Printf("input object logging: %v, %s, %T, %+v", rc.Path(), rc.Field.Name, obj, obj)
	return next(ctx)
}

func ObjectLogging(ctx context.Context, obj any, next graphql.Resolver) (res any, err error) {
	rc := graphql.GetFieldContext(ctx)
	log.Printf("object logging: %v, %s, %T, %+v", rc.Path(), rc.Field.Name, obj, obj)
	return next(ctx)
}

func ScalarLogging(ctx context.Context, obj any, next graphql.Resolver) (res any, err error) {
	rc := graphql.GetFieldContext(ctx)
	log.Printf("scalar logging: %v, %s, %T, %+v", rc.Path(), rc.Field.Name, obj, obj)
	return next(ctx)
}

func UnionLogging(ctx context.Context, obj any, next graphql.Resolver) (res any, err error) {
	rc := graphql.GetFieldContext(ctx)
	log.Printf("union logging: %v, %s, %T, %+v", rc.Path(), rc.Field.Name, obj, obj)
	return next(ctx)
}
