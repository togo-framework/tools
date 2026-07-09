// Package tools is the Tools content pack for togo: Tool/stack resource with favourites.
//
// A content pack bundles resource definitions (model + sqlc + Atlas + REST +
// GraphQL + a page) that a togo app pulls in. Generate the resource with
// `togo make:resource` in the host app; this plugin wires the provider hook.
package tools

import "github.com/togo-framework/togo"

func init() {
	togo.RegisterProviderFunc("tools", togo.PriorityService, func(k *togo.Kernel) error {
		return nil
	})
}
