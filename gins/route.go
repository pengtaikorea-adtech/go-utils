package gins

import (
	"errors"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/pengtaikorea-adtech/go-utils/slices"
)

// RouteEntity - route entity
type RouteEntity struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}

// RouteGroup - route groups
type RouteGroup struct {
	Path      string
	Middles   gin.HandlersChain
	Entities  []RouteEntity
	SubGroups []RouteGroup
}

// RegisterGroup register group on the router
func RegisterGroup(router gin.IRouter, group RouteGroup) {
	// tail-recursion traverse starts,
	stack := []hierarchicalRegisterEntry{
		{router, group},
	}

	for 0 < len(stack) {
		// pop first (BFS)
		rs, stack := stack[0], stack[1:]
		// this router group
		grp := rs.rt.Group(rs.grp.Path)
		// register middlewares
		grp.Use(rs.grp.Middles...)
		// register route entities
		slices.Each(func(e interface{}, i int, es interface{}) error {
			if rt, ok := e.(RouteEntity); ok {
				grp.Handle(rt.Method, rt.Path, HandlerWrap(rt.Handler))
			}
			return nil
		}, rs.grp.Entities)
		// appending subgroups into stack
		subs, _ := slices.Map(func(e interface{}, i int, es interface{}) (interface{}, error) {
			if sg, ok := e.(RouteGroup); ok {
				return hierarchicalRegisterEntry{
					grp,
					sg,
				}, nil
			}
			return nil, errors.New("no result")
		}, rs.grp.SubGroups, reflect.TypeOf(hierarchicalRegisterEntry{}))
		subgroups := subs.([]hierarchicalRegisterEntry)
		stack = append(stack, subgroups...)
	}

}

type hierarchicalRegisterEntry struct {
	rt  gin.IRouter
	grp RouteGroup
}
